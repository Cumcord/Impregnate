package src

import (
	"github.com/Cumcord/impregnate/middle"
	"github.com/Cumcord/impregnate/middle/api"
)

var pluginList []api.Plugin

func (app *UpApplication) ShowPreface() {
	app.ShowWaiter("Loading...", func(progress func(string)) {
		progress("Fetching plugins...")
		pluginList = middle.FetchRemotePlugins()
	}, func() {
		app.CachedPrimaryView = nil
		app.ShowPrimaryView(pluginList)
	})

}
