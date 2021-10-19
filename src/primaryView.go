package src

import (
	"os"

	"github.com/Cumcord/impregnate/middle"
	"github.com/Cumcord/impregnate/middle/api"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
)

func (app *UpApplication) ShowPrimaryView(pluginList []api.Plugin, warnings []framework.FlexboxSlot) {
	if app.CachedPrimaryView != nil {
		app.Teleport(app.CachedPrimaryView)
		return
	}

	slots := []framework.FlexboxSlot{}

	middle.FindWarnings(func(foundWarns []middle.Warning) {
		slots := []framework.FlexboxSlot{}
		for _, v := range foundWarns {
			fixAction := framework.ButtonBehavior(nil)
			if v.Action == middle.URLAndCloseWarningID {
				url := v.Parameter
				fixAction = func() {
					middle.OpenURL(url)
					os.Exit(0)
				}
			}
			warnings = append(slots, framework.FlexboxSlot{
				Element: design.InformationPanel(design.InformationPanelDetails{
					Text:       v.Text,
					ActionText: "FIX",
					Action:     fixAction,
				}),
			})
		}

		slots = append(slots, warnings...)
		if warnings != nil {
			app.ShowPrimaryView(pluginList, slots)
		}
	})

	pluginListItems := []design.ListItemDetails{}
	for _, item := range pluginList {
		localPlugin := item

		pluginListItems = append(pluginListItems, design.ListItemDetails{
			Text:    item.Name,
			Subtext: item.Description,
			Click: func() {
				app.GSRightwards()
				app.ShowPluginView(func() {
					app.GSLeftwards()
					app.ShowPrimaryView(pluginList, warnings)
				}, localPlugin, nil)
			},
		})
	}

	slots = append(slots, framework.FlexboxSlot{
		Element: design.NewUISearchBoxPtr("Search...", pluginListItems),
		Grow:    1,
	})

	app.CachedPrimaryView = design.LayoutDocument(design.Header{
		Title: "Plugins",
		Back: func() {
			app.CachedPrimaryView = nil
			app.GSLeftwards()
			app.ResetWithDiscordInstance(false, "computer://")
		},
		BackIcon:    design.BackIconID,
		ForwardIcon: design.MenuIconID,
		Forward: func() {
			app.GSRightwards()
			app.ShowOptionsMenu(func() {
				app.GSLeftwards()
				app.ShowPrimaryView(pluginList, warnings)
			})
		},
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       slots,
	}), true)

	app.Teleport(app.CachedPrimaryView)
}
