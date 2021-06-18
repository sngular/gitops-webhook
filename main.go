package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fluxcd/pkg/runtime/events"
)

var requests []*events.Event

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	var event *events.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	requests = append(requests, event)

	fmt.Printf("%s - %s(%s) in %s | %s %s: %s\n", event.Timestamp, event.InvolvedObject.Name, event.InvolvedObject.Kind, event.InvolvedObject.Namespace, event.ReportingController, event.Severity, event.Message)

	fmt.Fprintln(w, "Notification received!")
}

func handleAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Total notifications: %d\n\n", len(requests))

	for i, e := range requests {
		fmt.Fprintf(w, "Notification: %d\n", i+1)
		fmt.Fprintln(w, "  Involved Object:")
		fmt.Fprintf(w, "    Resource type: %s\n", e.InvolvedObject.Kind)
		fmt.Fprintf(w, "    Name: %s\n", e.InvolvedObject.Name)
		fmt.Fprintf(w, "    Namespace: %s\n", e.InvolvedObject.Namespace)
		fmt.Fprintf(w, "    Api version: %s\n", e.InvolvedObject.APIVersion)
		fmt.Fprintf(w, "    UID: %s\n", e.InvolvedObject.UID)
		fmt.Fprintf(w, "    Resource version: %s\n", e.InvolvedObject.ResourceVersion)
		fmt.Fprintf(w, "  Severity: %s\n", e.Severity)
		fmt.Fprintf(w, "  Timestamp: %s\n", e.Timestamp)
		fmt.Fprintf(w, "  Message: %s\n", e.Message)
		fmt.Fprintf(w, "  Reason: %s\n", e.Reason)
		fmt.Fprintf(w, "  Reporting Controller: %s\n", e.ReportingController)
		fmt.Fprintf(w, "  Reporting Instance: %s\n", e.ReportingController)
		fmt.Fprintln(w, "---------------------------------------------------------------------------------")
	}
}

func handleClear(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Notifications cleared!")
	requests = requests[:0]
}

func main() {
	port := "8080"
	log.Printf("Server started in port %s\n", port)

	http.HandleFunc("/webhook", handleWebhook)
	http.HandleFunc("/all", handleAll)
	http.HandleFunc("/clear", handleClear)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
