package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Cumcord/impregnate/middle/api"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
)

var baseURL = "https://cumcordplugins.github.io/Condom"

func getPluginViewPluginList() []struct {
	string
	api.Plugin
} {
	plugins := make(map[string]api.Plugin)

	resp, err := http.Get(fmt.Sprintf("%s/plugins-large.json", baseURL))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	json.Unmarshal([]byte(body), &plugins)

	var pluginList []struct {
		string
		api.Plugin
	}
	for key := range plugins {
		pluginList = append(pluginList, struct {
			string
			api.Plugin
		}{key, plugins[key]})
	}

	return pluginList
}

func (app *UpApplication) ShowPluginListView() {

	slots := []framework.FlexboxSlot{}

	pluginList := getPluginViewPluginList()
	pluginListItems := []design.ListItemDetails{}
	for _, item := range pluginList {
		localPlugin := item

		pluginListItems = append(pluginListItems, design.ListItemDetails{
			Text:    item.Plugin.Name,
			Subtext: item.Plugin.Description,
			Click: func() {
				app.ShowPluginView(func() {
					app.GSRightwards()
					app.ShowPluginListView()
				}, localPlugin)
			},
		})
	}

	slots = append(slots, framework.FlexboxSlot{
		Element: design.NewUISearchBoxPtr("Search...", pluginListItems),
		Grow:    1,
	})

	app.Teleport(design.LayoutDocument(design.Header{
		Title: "Plugins",
		Forward: func() {
			app.GSRightwards()
			app.ShowPrimaryView()
		},
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       slots,
	}), true))
}