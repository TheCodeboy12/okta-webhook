package structs

type handledObject struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	HandlerURL string `json:"handlerUrl"`
}
type HandledAppsList struct {
	Apps []handledObject `json:"handledApps"`
}

/*
*
Checks if an app id is in the list.
Returns a pointer to it if found, nil otherwie.
*/
func (h *HandledAppsList) Find(id string) *handledObject {
	for _, app := range h.Apps {
		if app.Id == id {
			return &app
		}
	}
	return nil
}
