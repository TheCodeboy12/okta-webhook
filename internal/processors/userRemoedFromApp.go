package processors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	oktaTypes "github.com/theCodeBoy12/oktaWebhook/internal/constants"
	"github.com/theCodeBoy12/oktaWebhook/internal/server/structs"
)

func (p *Processor) userRemovedFromApplication() error {
	var userTarget *structs.Target
	var appTarget *structs.Target
	events := p.EventHook.Data.Events

	// usually there is only one events so this should run once.
	for _, event := range events {
		for _, target := range event.Target {
			if target.Type == oktaTypes.AppType {
				appTarget = &target
				break
			}
		}
		for _, target := range event.Target {
			if target.Type == oktaTypes.UserType {
				userTarget = &target
				break
			}
		}

		handledApps, err := p.readConfFile()
		if err != nil {
			// slog.Error("Failed to read conf file", "error", err)
			return fmt.Errorf("failed to read conf file: %w", err)
		}

		// find the app
		app := handledApps.Find(appTarget.ID)
		if app == nil {
			// slog.Error("App not found in conf file", "app_id", appTarget.ID)
			return fmt.Errorf("app not found in conf file: %s", appTarget.ID)
		}

		if appTarget.ID == app.Id {
			// prep the body for the request we are about to send
			body, err := json.Marshal(&structs.ActionPayload{
				Type:    event.EventType,
				UserKey: userTarget.AlternateID,
			})
			if err != nil {
				// slog.Error("Failed to marshal action", "error", err)
				return fmt.Errorf("failed to marshal action: %w", err)
			}

			resps, err := http.Post(app.HandlerURL, "application/json", bytes.NewBuffer(body))
			if err != nil {
				// slog.Error("Failed to send action", "error", err)
				return fmt.Errorf("failed to send action: %w", err)
			}
			defer resps.Body.Close()
		}
	}
	return nil
}
