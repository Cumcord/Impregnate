package middle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
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
	var finalData ReturnData

	results := make(chan ReturnData, rangeLength)
	var wg sync.WaitGroup
	for current := 10; current < rangeStart+rangeLength; current++ {
		wg.Add(1)
		go func(current int, results chan ReturnData) {
			defer wg.Done()
			checkHealthForPort(current, results)
		}(current, results)
	}
	wg.Wait()

	for result := range results {
		if result.Name == "CUMCORD_WEBSOCKET" {
			finalData = result
		} else {
			continue
		}
	}

	return finalData
}

func checkHealthForPort(port int, results chan<- ReturnData) {
	var returnData ReturnData
	var finalData ReturnData

	var data api.WebsocketData
	data.Action = "GET_INFO"
	data.UUID = "a"
	d, _ := json.Marshal(&data)

	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("127.0.0.1:%d", port),
		Path:   "/cumcord",
	}
	var Dialer = websocket.DefaultDialer
	Dialer.HandshakeTimeout = time.Duration(time.Duration.Milliseconds(500))
	c, _, err := Dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = c.WriteMessage(websocket.TextMessage, d); err != nil {
	}
	_, message, err := c.ReadMessage()
	if err != nil {
	}
	defer c.Close()

	json.Unmarshal([]byte(message), &returnData)

	if returnData.Name == "CUMCORD_WEBSOCKET" {
		finalData.Name = "CUMCORD_WEBSOCKET"
	}
	results <- finalData
}
