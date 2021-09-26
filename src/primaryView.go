package src

import (
	"github.com/20kdc/CCUpdaterUI/design"
	"github.com/20kdc/CCUpdaterUI/frenyard"
	"github.com/20kdc/CCUpdaterUI/frenyard/framework"
	"github.com/20kdc/CCUpdaterUI/frenyard/integration"
)

func (app *UpApplication) ShowPrimaryView() {
	slots := []framework.FlexboxSlot{}

	slots = append(slots, framework.FlexboxSlot{
		Element: framework.NewUILabelPtr(integration.NewTextTypeChunk("Welcome to the Cumcord installer!", design.GlobalFont), 0xFFFFFFFF, 0, frenyard.Alignment2i{}),
		Grow:    1,
	})

	app.Teleport(design.LayoutDocument(design.Header{
		Title: "Impregnate",
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       slots,
	}), true))
}
