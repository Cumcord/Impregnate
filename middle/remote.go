package middle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Cumcord/impregnate/middle/api"
	"github.com/gorilla/websocket"
)

var baseURL = "https://cumcordplugins.github.io/Condom"

func FetchRemotePlugins() []api.Plugin {
	plugins := make(map[string]api.Plugin)

	resp, err := http.Get(fmt.Sprintf("%s/plugins-large.json", baseURL))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

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

	var Dialer = websocket.DefaultDialer
	Dialer.HandshakeTimeout = time.Duration(time.Duration.Milliseconds(500))
	for current <= rangeStart+rangeLength {
		u := url.URL{
			Scheme: "ws",
			Host:   fmt.Sprintf("127.0.0.1:%d", current),
			Path:   "/cumcord",
		}
		d, _ := json.Marshal(&data)
		c, _, err := Dialer.Dial(u.String(), nil)
		current += 1
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if err = c.WriteMessage(websocket.TextMessage, d); err != nil {
			continue
		}
		_, message, err := c.ReadMessage()
		if err != nil {
			continue
		}
		defer c.Close()

		json.Unmarshal([]byte(message), &returnData)

		if returnData.Name == "CUMCORD_WEBSOCKET" {
			finalData.Name = "CUMCORD_WEBSOCKET"
			break
		}
		fmt.Println(returnData)
	}

	return finalData
}
