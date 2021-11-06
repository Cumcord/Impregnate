package src

import (
	"fmt"

	"github.com/Cumcord/impregnate/middle"
	"github.com/lexisother/frenyard"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
	"github.com/lexisother/frenyard/integration"
)

func (app *UpApplication) ResetWithDiscordInstance(save bool, location string) {
	app.DiscordInstance = nil
	app.Config.DiscordPath = location
	if save {
		middle.WriteConfig(app.Config)
	}
	app.ShowPreface()
}

func (app *UpApplication) ShowPreface() {
	var discordLocations []middle.DiscordInstance
	// var pluginList []api.Plugin
	app.ShowWaiter("Loading...", func(progress func(string)) {
		progress("Checking local installation...")
		di, err := middle.NewDiscordInstance(app.Config.DiscordPath)
		if err == nil {
			app.DiscordInstance = di
		} else {
			fmt.Printf("Failed check: %s\n", err.Error())
		}
		// progress("Fetching remote plugins...")
		// pluginList = middle.FetchRemotePlugins()
		progress("Not configured ; Autodetecting Discord locations...")
		discordLocations = middle.GetChannels()
	}, func() {
		if app.DiscordInstance == nil {
			app.ShowInstanceFinder(discordLocations)
		} else {
			app.CachedPrimaryView = nil
			app.ShowPrimaryView()
		}
	})
}

func (app *UpApplication) ShowInstanceFinder(locations []middle.DiscordInstance) {
	suggestSlots := []framework.FlexboxSlot{}
	for _, location := range locations {
		suggestSlots = append(suggestSlots, framework.FlexboxSlot{
			Element: design.ListItem(design.ListItemDetails{
				Icon:    design.DirectoryIconID,
				Text:    location.Channel,
				Subtext: location.Path,
				Click: func() {
					app.GSRightwards()
					app.ResetWithDiscordInstance(true, location.Path)
				},
			}),
			RespectMinimumSize: true,
		})
	}

	suggestSlots = append(suggestSlots, framework.FlexboxSlot{
		Grow:   1,
		Shrink: 0,
	})

	foundInstallsScroller := design.ScrollboxV(framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		WrapMode:    framework.FlexboxWrapModeNone,
		Slots:       suggestSlots,
	}))

	content := framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots: []framework.FlexboxSlot{
			{
				Element: framework.NewUILabelPtr(integration.NewTextTypeChunk("Welcome to the official Cumcord installer. Before we begin, we need to know what Discord instance to fuck up.", design.GlobalFont), design.ThemeText, 0, frenyard.Alignment2i{}),
			},
			{
				Basis: design.SizeMarginAroundEverything,
			},
			{
				Element:            foundInstallsScroller,
				Grow:               1,
				Shrink:             1,
				RespectMinimumSize: true,
			},
			{
				Basis: design.SizeMarginAroundEverything,
			},
		},
	})

	primary := design.LayoutDocument(design.Header{
		Title:    "Welcome",
		BackIcon: design.WarningIconID,
		Back: func() {
			app.GSLeftwards()
			app.ShowCredits(func() {
				app.GSRightwards()
				app.ShowInstanceFinder(locations)
			})
		},
	}, content, true)
	app.Teleport(primary)
}
