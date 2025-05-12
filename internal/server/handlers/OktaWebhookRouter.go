package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/theCodeBoy12/oktaWebhook/internal/processors"
	"github.com/theCodeBoy12/oktaWebhook/internal/server/structs"
)

func OktaWebhookRouter(ConfFilePath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			slog.Error("Invalid request body", "error", err)
			return
		}
		defer r.Body.Close()

		var eventHook structs.OktaEventHook
		err = json.Unmarshal(body, &eventHook)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			slog.Error("Invalid request body", "error", err)
			return
		}
		slog.Debug("Received event", "event", eventHook)
		err = eventHook.Validate()
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			slog.Error("Invalid request body", "error", err)
			return
		}
		// switch and check whether the event is one of the events that we can handle with this webhook handler.
		// for _, event := range eventHook.Data.Events {
		// 	switch event.EventType {
		// 	case oktaTypes.UserAddedToGroup:

		// 	case oktaTypes.UserRemovedFromGroup:

		// 	case oktaTypes.UserAddedtoApplication:

		// 	case oktaTypes.UserRemovedFromApplication:
		// 		go processors.UserRemovedFromApplication(event)
		// 	default:
		// 		http.Error(w, "Invalid event type", http.StatusBadRequest)
		// 		slog.Error("Invalid event type", "event_type", event.EventType)
		// 		return
		// 	}
		// }
		h := processors.Processor{
			EventHook:    eventHook,
			ConfFilePath: ConfFilePath,
		}
		// the event poster did its job and were now switching to handling the request solely on the server
		go h.Process()
		// so we reply back to it
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

	})
}
