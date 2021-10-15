package middle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Cumcord/impregnate/middle/api"
	"golang.org/x/net/websocket"
)

var baseURL = "https://cumcordplugins.github.io/Condom"

func FetchRemotePlugins() []api.Plugin {
	plugins := make(map[string]api.Plugin)

	resp, err := http.Get(fmt.Sprintf("%s/plugins-large.json", baseURL))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	json.Unmarshal([]byte(body), &plugins)

	var pluginList []api.Plugin
	for key := range plugins {
		plugin := plugins[key]
		plugin.URL = key
		pluginList = append(pluginList, plugin)
	}

	return pluginList
}

type ReturnData struct {
	Name string `json:"name"`
}

func CheckHealth() ReturnData {
	rangeStart := 6463
	rangeLength := 10
	current := rangeStart
	var data api.WebsocketData
	var returnData ReturnData
	var finalData ReturnData

	data.Action = "GET_INFO"
	data.UUID = "a"

	for current <= rangeStart+rangeLength {
		fmt.Println(current)
		conn, err := websocket.Dial(fmt.Sprintf("ws://127.0.0.1:%d/cumcord", current), "", "http://localhost/")
		current += 1
		if err != nil {
			continue
		}
		if err = websocket.JSON.Send(conn, &data); err != nil {
			continue
		}
		if err = websocket.JSON.Receive(conn, &returnData); err != nil {
			continue
		}
		defer conn.Close()

		if returnData.Name == "CUMCORD_WEBSOCKET" {
			finalData.Name = "CUMCORD_WEBSOCKET"
		}
	}

	fmt.Println(finalData)
	return finalData
}
