package api

type WebsocketData struct {
	Action string `json:"action"`
	UUID   string `json:"uuid"`
	URL    string `json:"url"`
	Code   string `json:"code"`
}
