package src

import (
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
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
		{
			Text:    "Yellowsink",
			Subtext: "Better to cum in the sink, than to sink in the cum.",
		},
		{
			Text:    "20kdc",
			Subtext: "Thank you for building this great framework.",
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
