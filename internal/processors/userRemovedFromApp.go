package processors

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	oktaTypes "github.com/theCodeBoy12/oktaWebhook/internal/constants"
	"github.com/theCodeBoy12/oktaWebhook/internal/server/structs"
)

func UserRemovedFromApplication(event structs.Event) {
	// TODO: based on the app, do something
	//Find what app were dealing with.
	var userTarget structs.Target
	var appTarget structs.Target
	// find the targets
	for _, target := range event.Target {
		if target.Type == oktaTypes.AppType {
			appTarget = target
			break
		}
	}
	for _, target := range event.Target {
		if target.Type == oktaTypes.UserType {
			userTarget = target
			break
		}
	}
	f, err := os.ReadFile("/Users/gabriel.s/okta webhook/internal/processors/handledApps.json")

	if err != nil {
		slog.Error("Failed to read handledApps.json", "error", err)
		return
	}
	var handledApps structs.HandledAppsList

	err = json.Unmarshal(f, &handledApps)
	if err != nil {
		slog.Error("Failed to unmarshal handledApps.json", "error", err)
		return
	}
	for _, app := range handledApps.Apps {
		if appTarget.ID == app.Id {

			body, err := json.Marshal(&structs.ActionPayload{
				Type:    event.EventType,
				UserKey: userTarget.AlternateID,
			})
			if err != nil {
				slog.Error("Failed to marshal action", "error", err)
				return
			}
			// time.Sleep(3 * time.Second)
			// slog.Info("Got body", "body", string(body))
			resps, err := http.Post(app.HandlerURL, "application/json", bytes.NewBuffer(body))
			if err != nil {
				slog.Error("Failed to send action", "error", err)
				return
			}
			defer resps.Body.Close()

		}
	}
}
