package processors

import (
	"encoding/json"
	"log/slog"
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
			if err := p.userRemovedFromApplication(); err != nil {
				slog.Error("Failed to process event", "error", err)
				return
			}
		default:
			slog.Error("Invalid event type", "event_type", event.EventType)
			return
		}
	}
}
