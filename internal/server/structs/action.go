package structs

import "fmt"

type ActionPayload struct {
	Type    string
	UserKey string
}

func (a *ActionPayload) Validate() error {
	if a.Type == "" {
		return fmt.Errorf("type cannot be empty")
	}
	if a.UserKey == "" {
		return fmt.Errorf("userKey cannot be empty")
	}
	switch a.Type {
	case "remove":
		return nil
	case "add":
		return nil
	default:
		return fmt.Errorf("invalid type")
	}
}
