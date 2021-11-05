package src

import (
	"os"
	"path"
	"runtime"
	"time"

	"github.com/lexisother/frenyard"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
	"github.com/lexisother/frenyard/integration"
)

func (app *UpApplication) ShowOptionsMenu(back framework.ButtonBehavior) {
	backHere := func() {
		app.GSLeftwards()
		app.ShowOptionsMenu(back)
	}

	listSlots := []framework.FlexboxSlot{
		{
			Element: design.ListItem(design.ListItemDetails{
				Text:    "Credits",
				Subtext: "See who is behind Impregnate and related",
				Click: func() {
					app.GSRightwards()
					app.ShowCredits(backHere)
				},
			}),
		},
		{
			Element: design.ListItem(design.ListItemDetails{
				Text:    "Build Information",
				Subtext: runtime.GOOS + " " + runtime.GOARCH + " " + runtime.Compiler + " " + runtime.Version(),
			}),
		},
		{
			Grow: 1,
		},
	}

	if _, err := os.Stat(path.Join(app.Config.DiscordPath, "resources/app/plugged.txt")); err == nil {
		listSlots = append([]framework.FlexboxSlot{
			{
				Element: design.ListItem(design.ListItemDetails{
					Text:    "Uninstall Cumcord",
					Subtext: "Uninstall the client mod.",
					Click: func() {
						app.GSRightwards()
						app.ShowUninstallMenu(backHere)
					},
				}),
			},
		}, listSlots...)
	}

	app.Teleport(design.LayoutDocument(design.Header{
		Title: "Options",
		Back:  back,
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       listSlots,
	}), true))
}

func (app *UpApplication) ShowUninstallMenu(back framework.ButtonBehavior) {
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
							log := "-- Log started at " + time.Now().Format(time.RFC1123) + " --"
							app.ShowWaiter("Uninstalling...", func(progress func(string)) {
								log += "\nDeleting the app directory..."
								progress(log)
								os.RemoveAll(path.Join(app.Config.DiscordPath, "resources/app"))
								log += "\nDone! Checking if plug directory exists..."
								progress(log)
								if _, err := os.Stat(path.Join(app.Config.DiscordPath, "resources/plug")); err == nil {
									log += "\nRestoring the plug directory..."
									progress(log)
									os.Rename(path.Join(app.Config.DiscordPath, "resources/plug"), path.Join(app.Config.DiscordPath, "resources/app"))
								}
								log += "\n-- Complete; Restart your Discord client! --"
								progress(log)
							}, func() {
								app.MessageBox("Uninstall Complete", log, func() {
									app.CachedPrimaryView = nil
									back()
								})
							})
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
		Title: "Uninstall",
		Back:  back,
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       listSlots,
	}), true))
}
