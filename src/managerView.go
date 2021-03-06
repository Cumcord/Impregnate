package src

import (
	"encoding/base64"
	"os"
	"path"
	"time"

	"github.com/Cumcord/impregnate/middle"
	"github.com/lexisother/frenyard"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
	"github.com/lexisother/frenyard/integration"
)

func (app *UpApplication) ShowManagerView(installed bool, back framework.ButtonBehavior) {
	if !installed {
		showInstallScreen(app)
	} else {
		showUninstallScreen(app, back)
	}
}

func showInstallScreen(app *UpApplication) {
	if _, err := os.Stat(path.Join(app.Config.DiscordPath, "app/plugged.txt")); err == nil {
		app.MessageBox("Already Installed!", "Cumcord is already installed. Please restart your client.", func() {
			app.CachedPrimaryView = nil
			app.ShowPrimaryView()
		})
	} else {
		log := "-- Log started at " + time.Now().Format(time.RFC1123) + " --"
		app.ShowWaiter("Installing...", func(progress func(string)) {
			log += "\nChecking for app folder..."
			progress(log)
			resources, _ := os.Stat(path.Join(app.Config.DiscordPath, "app"))
			if resources != nil {
				log += "\nRenaming app folder..."
				progress(log)
				os.Rename(path.Join(app.Config.DiscordPath, "app"), path.Join(app.Config.DiscordPath, "plug"))
			}
			os.Mkdir(path.Join(app.Config.DiscordPath, "app"), 0755)
			index, _ := os.Create(path.Join(app.Config.DiscordPath, "app/index.js"))
			packageJson, _ := os.Create(path.Join(app.Config.DiscordPath, "app/package.json"))
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
			pluggedFile, _ := os.Create(path.Join(app.Config.DiscordPath, "app/plugged.txt"))
			pluggedFile.WriteString("this file was added to indicate this was a cumcord installation. balls.")
			app.MessageBox("Install Complete", log, func() {
				app.CachedPrimaryView = nil
				app.ShowPrimaryView()
			})
		})
	}
}

func showUninstallScreen(app *UpApplication, back framework.ButtonBehavior) {
	if _, err := os.Stat(path.Join(app.Config.DiscordPath, "app/plugged.txt")); err != nil {
		app.MessageBox("Not installed!", "Cumcord is not installed. Please install it before trying to remove it.", func() {
			app.CachedPrimaryView = nil
			app.ShowPrimaryView()
		})
	} else {
		listSlots := []framework.FlexboxSlot{
			{
				Grow: 1,
			},
			{
				Element: framework.NewUILabelPtr(integration.NewTextTypeChunk("Are you sure you want to uninstall Cumcord?", design.GlobalFont), 0xFFFFFFFF, 0, frenyard.Alignment2i{}),
			},
			{
				Basis:  frenyard.Scale(design.DesignScale, 32),
				Shrink: 1,
			},
			{
				Element: framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
					DirVertical: false,
					Slots: []framework.FlexboxSlot{
						{
							Grow: 1,
						},
						{
							Element: design.ButtonAction(design.ThemeRemoveActionButton, "Uninstall", func() {
								if _, err := os.Stat(path.Join(app.Config.DiscordPath, "app/plugged.txt")); err != nil {
								} else {
									log := "-- Log started at " + time.Now().Format(time.RFC1123) + " --"
									app.ShowWaiter("Uninstalling...", func(progress func(string)) {
										log += "\nDeleting the app directory..."
										progress(log)
										os.RemoveAll(path.Join(app.Config.DiscordPath, "app"))
										log += "\nDone! Checking if plug directory exists..."
										progress(log)
										if _, err := os.Stat(path.Join(app.Config.DiscordPath, "plug")); err == nil {
											log += "\nRestoring the plug directory..."
											progress(log)
											os.Rename(path.Join(app.Config.DiscordPath, "plug"), path.Join(app.Config.DiscordPath, "app"))
										}
										log += "\n-- Complete; Restart your Discord client! --"
										progress(log)
									}, func() {
										app.MessageBox("Uninstall Complete", log, func() {
											app.CachedPrimaryView = nil
											app.ShowPrimaryView()
										})
									})
								}
							}),
						},
						{
							Basis: frenyard.Scale(design.DesignScale, 32),
						},
						{
							Element: design.ButtonAction(design.ThemeOkActionButton, "Cancel", back),
						},
						{
							Grow: 1,
						},
					},
				}),
			},
			{
				Grow: 1,
			},
		}

		app.Teleport(design.LayoutDocument(design.Header{
			Title: "Cumcord",
			Back:  back,
		}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
			DirVertical: true,
			Slots:       listSlots,
		}), true))
	}
}
