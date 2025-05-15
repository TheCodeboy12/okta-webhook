package processors

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	oktaTypes "github.com/theCodeBoy12/oktaWebhook/internal/constants"
	"github.com/theCodeBoy12/oktaWebhook/internal/server/structs"
	"google.golang.org/api/idtoken"
)

func (p *Processor) userRemovedFromApplication() error {
	ctx := context.Background()
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
			// Create an HTTP client with ADC credentials built into it.
			// This is in order to post to a cloud run service located at the handler URL.
			client, err := idtoken.NewClient(ctx, app.HandlerURL)
			if err != nil {
				// slog.Error("Failed to create client", "error", err)
				return fmt.Errorf("failed to create client: %w", err)
			}
			resps, err := client.Post(app.HandlerURL, "application/json", bytes.NewBuffer(body))
			if err != nil {
				// slog.Error("Failed to send action", "error", err)
				return fmt.Errorf("failed to send action: %w", err)
			}
			defer resps.Body.Close()
		}
	}
	return nil
}
