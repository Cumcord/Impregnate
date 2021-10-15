package api

type Plugin struct {
	URL         string
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	License     string `json:"license"`
}
