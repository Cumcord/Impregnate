package src

import (
	"os"
	"path"

	"github.com/Cumcord/impregnate/middle"
	"github.com/lexisother/frenyard"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
	"github.com/lexisother/frenyard/integration"
)

func (app *UpApplication) ShowPrimaryView() {
	if app.CachedPrimaryView != nil {
		app.Teleport(app.CachedPrimaryView)
		return
	}

	var installStatus string
	if _, installedOrNot := os.Stat(path.Join(app.Config.DiscordPath, "resources/app/plugged.txt")); installedOrNot == nil {
		installStatus = "installed!"
	} else {
		installStatus = "not installed."
	}

	slots := []framework.FlexboxSlot{
		{
			Grow: 1,
		},
		{
			Element: framework.NewUILabelPtr(integration.NewTextTypeChunk("Welcome to the Cumcord installer!", design.GlobalFont), 0xFFFFFFFF, 0, frenyard.Alignment2i{}),
		},
		{
			Element: framework.NewUILabelPtr(integration.NewTextTypeChunk("Cumcord is currently: "+installStatus, design.GlobalFont), 0xFFFFFFFF, 0, frenyard.Alignment2i{}),
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
						Element: design.ButtonAction(design.ThemeOkActionButton, "Install", func() {
							app.GSRightwards()
							app.ShowManagerView(false, func() {
								app.GSLeftwards()
								app.ShowPrimaryView()
							})
						}),
						Shrink: 1,
					},
					{
						Basis:  frenyard.Scale(design.DesignScale, 32),
						Shrink: 1,
					},
					{
						Element: design.ButtonAction(design.ThemeRemoveActionButton, "Uninstall", func() {
							app.GSRightwards()
							app.ShowManagerView(true, func() {
								app.GSLeftwards()
								app.ShowPrimaryView()
							})
						}),
						Shrink: 1,
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

	warnings := middle.FindWarnings(app.Config)
	for _, v := range warnings {
		fixAction := framework.ButtonBehavior(nil)
		if v.Action == middle.URLAndCloseWarningID {
			url := v.Parameter
			fixAction = func() {
				middle.OpenURL(url)
				os.Exit(0)
			}
		}
		slots = append([]framework.FlexboxSlot{
			{
				Element: design.InformationPanel(design.InformationPanelDetails{
					Text:       v.Text,
					ActionText: "FIX",
					Action:     fixAction,
				}),
			},
		}, slots...)
	}

	app.CachedPrimaryView = design.LayoutDocument(design.Header{
		Title: "Impregnate",
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
				app.ShowPrimaryView()
			})
		},
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       slots,
	}), true)

	app.Teleport(app.CachedPrimaryView)
}
