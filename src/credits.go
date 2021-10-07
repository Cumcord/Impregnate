package src

import (
	"github.com/yellowsink/frenyard/design"
	"github.com/yellowsink/frenyard/framework"
)

func (app *UpApplication) ShowCredits(back framework.ButtonBehavior) {
	items := []design.ListItemDetails{
		{
			Text:    "Alyxia Sother",
			Subtext: "Developer of Impregnate, Cumcord semen demon.",
		},
		{
			Text:    "Creatable",
			Subtext: "Creator of Cumcord, God of Cum.",
		},
	}

	listSlots := []framework.FlexboxSlot{}
	for _, item := range items {
		listSlots = append(listSlots, framework.FlexboxSlot{
			Element: design.ListItem(item),
		})
	}

	app.Teleport(design.LayoutDocument(design.Header{
		Title: "Credits",
		Back:  back,
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       listSlots,
	}), true))
}
