package src

import (
	"github.com/lexisother/frenyard"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
	"github.com/lexisother/frenyard/integration"
)

func (app *UpApplication) ShowPrimaryView() {
	slots := []framework.FlexboxSlot{
		{
			Grow: 1,
		},
		{
			Element: framework.NewUILabelPtr(integration.NewTextTypeChunk("Welcome to the Cumcord installer!", design.GlobalFont), 0xFFFFFFFF, 0, frenyard.Alignment2i{}),
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
						Element: design.ButtonAction(design.ThemeOkActionButton, "Install", func() {}),
						Shrink:  1,
					},
					{
						Basis:  frenyard.Scale(design.DesignScale, 32),
						Shrink: 1,
					},
					{
						Element: design.ButtonAction(design.ThemeRemoveActionButton, "Uninstall", func() {}),
						Shrink:  1,
					},
					{
						Grow: 1,
					},
				},
			}),
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
						Element: design.ButtonAction(design.ThemeUpdateActionButton, "Plugins", func() {
							app.GSLeftwards()
							app.ShowPluginListView()
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

	app.Teleport(design.LayoutDocument(design.Header{
		Title:       "Impregnate",
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
	}), true))
}
