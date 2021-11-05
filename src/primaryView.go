package src

import (
	"encoding/base64"
	"os"
	"path"
	"time"

	"github.com/Cumcord/impregnate/middle"
	"github.com/Cumcord/impregnate/middle/api"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
)

func (app *UpApplication) ShowPrimaryView(pluginList []api.Plugin) {
	if app.CachedPrimaryView != nil {
		app.Teleport(app.CachedPrimaryView)
		return
	}

	slots := []framework.FlexboxSlot{}

	warnings := middle.FindWarnings(app.Config)
	for _, v := range warnings {
		fixAction := framework.ButtonBehavior(nil)
		if v.Action == middle.URLAndCloseWarningID {
			url := v.Parameter
			fixAction = func() {
				middle.OpenURL(url)
				os.Exit(0)
			}
		} else if v.Action == middle.InstallOrUpdatePackageWarningID {
			fixAction = func() {
				app.GSRightwards()

				log := "-- Log started at " + time.Now().Format(time.RFC1123) + " --"
				app.ShowWaiter("Installing...", func(progress func(string)) {
					log += "\nChecking for app folder..."
					progress(log)
					resources, _ := os.Stat(path.Join(app.Config.DiscordPath, "resources/app"))
					if resources != nil {
						log += "\nRenaming app folder..."
						progress(log)
						os.Rename(path.Join(app.Config.DiscordPath, "resources/app"), path.Join(app.Config.DiscordPath, "resources/plug"))
					}
					os.Mkdir(path.Join(app.Config.DiscordPath, "resources/app"), 0755)
					index, _ := os.Create(path.Join(app.Config.DiscordPath, "resources/app/index.js"))
					packageJson, _ := os.Create(path.Join(app.Config.DiscordPath, "resources/app/package.json"))
					log += "\nWriting package.json..."
					progress(log)
					packageJson.WriteString(`{"name":"plug","main":"index.js"}`)
					log += "\nWriting index.js..."
					progress(log)
					decodedInjector, _ := base64.StdEncoding.DecodeString(middle.InjectorCode)
					index.WriteString(string(decodedInjector))
					log += "\n-- Complete; Restart your Discord client! --"
					progress(log)
				}, func() {
					if _, err := os.Stat(path.Join(app.Config.DiscordPath, "resources/app/plugged.txt")); err == nil {
						app.MessageBox("Already Installed!", "Cumcord is already installed. Please start your client.", func() {
							app.CachedPrimaryView = nil
							app.ShowPrimaryView(pluginList)
						})
					} else {
						pluggedFile, _ := os.Create(path.Join(app.Config.DiscordPath, "resources/app/plugged.txt"))
						pluggedFile.WriteString("this file was added to indicate this was a cumcord installation. balls.")
						app.MessageBox("Install Complete", log, func() {
							app.CachedPrimaryView = nil
							app.ShowPrimaryView(pluginList)
						})
					}
				})
			}
		}
		slots = append(slots, framework.FlexboxSlot{
			Element: design.InformationPanel(design.InformationPanelDetails{
				Text:       v.Text,
				ActionText: "FIX",
				Action:     fixAction,
			}),
		})
	}

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
					app.ShowPrimaryView(pluginList)
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
				app.ShowPrimaryView(pluginList)
			})
		},
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       slots,
	}), true)

	app.Teleport(app.CachedPrimaryView)
}
