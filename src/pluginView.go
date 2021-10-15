package src

import (
	"fmt"

	"github.com/Cumcord/impregnate/middle/api"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
	"golang.org/x/net/websocket"
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
		{
			Grow: 1,
		},
	}

	buttons := []framework.UILayoutElement{
		design.ButtonAction(design.ThemeOkActionButton, "Install", func() {

			rangeStart := 6463
			rangeLength := 10
			current := rangeStart
			var data api.WebsocketData
			var returnData map[string]interface{}

			data.Action = "INSTALL_PLUGIN"
			data.UUID = "a"
			data.URL = fmt.Sprintf("https://cumcordplugins.github.io/Condom/%s", plugin.string)

			for current <= rangeStart+rangeLength {
				fmt.Println(current)
				conn, err := websocket.Dial(fmt.Sprintf("ws://127.0.0.1:%d/cumcord", current), "", "http://localhost/")
				current += 1
				if err != nil {
					continue
				}
				if err = websocket.JSON.Send(conn, &data); err != nil {
					continue
				}
				if err = websocket.JSON.Receive(conn, &returnData); err != nil {
					continue
				}
				defer conn.Close()
				fmt.Println(returnData)

			}
		}),
	}

	slots = append(slots, framework.FlexboxSlot{
		Element: design.ButtonBar(buttons),
	})

	app.Teleport(design.LayoutDocument(design.Header{
		Title: plugin.Plugin.Name,
		Back:  back,
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       slots,
	}), true))
}
