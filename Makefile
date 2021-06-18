go_version ?= 1.16
image ?= ghcr.io/sngular/gitops-webhook
tag = $(shell git rev-parse --short HEAD)

pre-commit: tidy fmt vet clean

build:
	go get && go build main.go

run:
	go run main.go

fmt:
	docker run --rm --name go-fmt \
		-e $UID=$(shell id -u) \
		-e $GUID=$(shell id -g) \
		-v $(shell pwd):/app \
		-v $(shell pwd)/.cache:/go/pkg \
		-w /app golang:${go_version} go fmt ./...
	git add **\*.go

tidy:
	docker run --rm --name go-tidy \
		-e $UID=$(shell id -u) \
		-e $GUID=$(shell id -g) \
		-v $(shell pwd):/app \
		-v $(shell pwd)/.cache:/go/pkg \
		-w /app golang:${go_version} go mod tidy

vet:
	docker run --rm --name go-vet \
		-e $UID=$(shell id -u) \
		-e $GUID=$(shell id -g) \
		-v $(shell pwd):/app \
		-v $(shell pwd)/.cache:/go/pkg \
		-w /app golang:${go_version} go vet ./...

clean:
	@rm -rf main bin/ *.out

docker-build:
	docker build --tag ${image}:${tag} .

docker-run: docker-build
	docker container run --rm --name gitops-webhook ${image}:${tag}

send-test-notification:
	@curl -X POST http://localhost:8080/webhook \
  	--data '{"involvedObject": {"kind":"GitRepository", "namespace":"flux-system", "name":"flux-system", "uid":"cc4d0095-83f4-4f08-98f2-d2e9f3731fb9", "apiVersion":"source.toolkit.fluxcd.io/v1beta1", "resourceVersion":"56921"}, "severity":"info", "timestamp":"2006-01-02T15:04:05Z", "message":"Fetched revision: main/731f7eaddfb6af01cb2173e18f0f75b0ba780ef1", "reason":"info", "reportingController":"source-controller", "reportingInstance":"source-controller-7c7b47f5f-8bhrp"}'

list-notifications:
	@curl -X POST http://localhost:8080/all

clear-notifications:
	@curl -X POST http://localhost:8080/clear
