package middle

import (
	"context"
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

func FindPortAndRun(callback func(data ReturnData)) {
	go func() {
		data := CheckHealth()
		callback(data)
	}()
}

func CheckHealth() ReturnData {
	rangeStart := 6463
	rangeLength := 10

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	output := make(chan ReturnData, rangeLength)
	for current := 10; current < rangeStart+rangeLength; current++ {
		go checkHealthForPort(ctx, current, output)
	}

	return <-output
}

func checkHealthForPort(ctx context.Context, port int, output chan ReturnData) {
	dctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	url := &url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("127.0.0.1:%d", port),
		Path:   "/cumcord",
	}

	c, _, err := websocket.DefaultDialer.DialContext(dctx, url.String(), nil)
	if err != nil {
		return
	}
	defer c.Close()

	msg, _ := json.Marshal(&api.WebsocketData{
		Action: "GET_INFO",
		UUID:   "a",
	})
	if err = c.WriteMessage(websocket.TextMessage, msg); err != nil {
		return
	}

	_, raw, err := c.ReadMessage()
	if err != nil {
		return
	}

	var result ReturnData
	json.Unmarshal([]byte(raw), &result)

	if result.Name == "CUMCORD_WEBSOCKET" {
		output <- result
	}
}
