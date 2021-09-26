package src

import (
	"runtime"

	"github.com/20kdc/CCUpdaterUI/design"
	"github.com/20kdc/CCUpdaterUI/frenyard/framework"
)

func (app *UpApplication) ShowOptionsMenu(back framework.ButtonBehavior) {
	listSlots := []framework.FlexboxSlot{
		{
			Element: design.ListItem(design.ListItemDetails{
				Text:    "Credits",
				Subtext: "See who is behind Impregnate and related",
				Click: func() {
					app.GSRightwards()
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

	app.Teleport(design.LayoutDocument(design.Header{
		Title: "Options",
		Back:  back,
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       listSlots,
	}), true))
}
