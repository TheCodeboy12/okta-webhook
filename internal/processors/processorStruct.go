package processors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	oktaTypes "github.com/theCodeBoy12/oktaWebhook/internal/constants"
	"github.com/theCodeBoy12/oktaWebhook/internal/server/structs"
)

type Processor struct {
	EventHook    structs.OktaEventHook
	ConfFilePath string
}

func (p *Processor) readConfFile() (structs.HandledAppsList, error) {
	f, err := os.ReadFile(p.ConfFilePath)
	if err != nil {
		return structs.HandledAppsList{}, err
	}
	var handledApps structs.HandledAppsList
	if err := json.Unmarshal(f, &handledApps); err != nil {
		return structs.HandledAppsList{}, err
	}
	return handledApps, nil
}

func (p *Processor) Process() {
	for _, event := range p.EventHook.Data.Events {
		switch event.EventType {
		case oktaTypes.UserAddedToGroup:

		case oktaTypes.UserRemovedFromGroup:

		case oktaTypes.UserAddedtoApplication:

		case oktaTypes.UserRemovedFromApplication:
			p.userRemovedFromApplication()
		default:
			slog.Error("Invalid event type", "event_type", event.EventType)
			return
		}
	}
}

func (p *Processor) userRemovedFromApplication() error {
	// TODO: based on the app, do something
	//Find what app were dealing with.
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
		for _, app := range handledApps.Apps {
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
	}
	return nil
}
