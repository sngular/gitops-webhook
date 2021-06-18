# gitops-webhook

Servicio Golang utilizado en los ejercicios de la serie [GitOps Flux](https://github.com/Sngular/gitops-flux-series).

Funcionalidades:
- Recibe notificaciones en el formato de eventos de [Flux](https://fluxcd.io/docs) en el endpoint `/webhook`.
- Lista todas las notificaciones cuando recibe una petición en el endpoint `/all`.
- Elimina todas las notificaciones recibidas si recibe una petición en el endpoint `/clear`.

El formato de las notificaciones se corresponde con el siguiente paquete: https://github.com/fluxcd/pkg/tree/main/runtime/events/

## Funcionamiento

Para ver su funcionamiento utilice el siguiente comando:

```bash
docker container run -d --rm ghcr.io/sngular/gitops-webhook:v0.1.0

2021/06/18 17:54:21 Server started in port 8080
```

Enviar notifiación de prueba:

```bash
# utilizando make
make send-test-notification

# utilizando curl
curl -X POST http://localhost:8080/webhook \
  --data '{"involvedObject": {"kind":"GitRepository", "namespace":"flux-system", "name":"flux-system", "uid":"cc4d0095-83f4-4f08-98f2-d2e9f3731fb9", "apiVersion":"source.toolkit.fluxcd.io/v1beta1", "resourceVersion":"56921"}, "severity":"info", "timestamp":"2006-01-02T15:04:05Z", "message":"Fetched revision: main/731f7eaddfb6af01cb2173e18f0f75b0ba780ef1", "reason":"info", "reportingController":"source-controller", "reportingInstance":"source-controller-7c7b47f5f-8bhrp"}'
```

<details>
  <summary>Resultado</summary>

  ```
  Notification received!
  ```
</details>

Listar notificaciones recibidas:

```bash
# utilizando make
make list-notifications

# utilizando curl
curl http://localhost:8080/all
```

<details>
  <summary>Resultado</summary>

  ```
  Total notifications: 1

  Notification: 1
    Involved Object:
      Resource type: GitRepository
      Name: flux-system
      Namespace: flux-system
      Api version: source.toolkit.fluxcd.io/v1beta1
      UID: cc4d0095-83f4-4f08-98f2-d2e9f3731fb9
      Resource version: 56921
    Severity: info
    Timestamp: 2006-01-02 16:04:05 +0100 CET
    Message: Fetched revision: main/731f7eaddfb6af01cb2173e18f0f75b0ba780ef1
    Reason: info
    Reporting Controller: source-controller
    Reporting Instance: source-controller
  ---------------------------------------------------------------------------------
 ```
</details>

Eliminar todas las notificaciones recibidas:

```bash
# utilizando make
make clear-notifications

# utilizando curl
curl http://localhost:8080/clear
```

<details>
  <summary>Resultado</summary>

  ```
  Notifications cleared!
  ```
</details>
