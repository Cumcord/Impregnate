package src

import (
	"github.com/Cumcord/impregnate/middle/api"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
)

func (app *UpApplication) ShowPluginView(back framework.ButtonBehavior, plugin struct {
	string
	api.Plugin
}) {
	slots := []framework.FlexboxSlot{
		{
			Element: design.ListItem(design.ListItemDetails{
				Text: "hi",
			}),
		},
	}

	app.Teleport(design.LayoutDocument(design.Header{
		Title:   plugin.Plugin.Name,
		Forward: back,
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       slots,
	}), true))
}
