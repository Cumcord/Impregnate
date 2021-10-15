package middle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Cumcord/impregnate/middle/api"
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
