package structs

type handledObject struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	HandlerURL string `json:"handlerUrl"`
}
type HandledAppsList struct {
	Apps []handledObject `json:"handledApps"`
}
