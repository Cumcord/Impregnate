package src

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/Cumcord/impregnate/middle/api"
	"github.com/gorilla/websocket"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
)

type ReturnData struct {
	Status string `json:"status"`
}

func (app *UpApplication) ShowPluginView(back framework.ButtonBehavior, plugin api.Plugin, warnings []framework.FlexboxSlot) {
	slots := []framework.FlexboxSlot{}

	if warnings != nil {
		for _, warning := range warnings {
			slots = append(slots, warning)
		}
	}

	slots = append(slots,
		framework.FlexboxSlot{
			Element: design.ListItem(design.ListItemDetails{
				Text: "hi",
			}),
		},
		framework.FlexboxSlot{
			Grow: 1,
		},
	)

	slots = append(slots)

	buttons := []framework.UILayoutElement{
		design.ButtonAction(design.ThemeOkActionButton, "Install", func() {

			rangeStart := 6463
			rangeLength := 10
			current := rangeStart
			var data api.WebsocketData
			var returnData ReturnData

			data.Action = "INSTALL_PLUGIN"
			data.UUID = "a"
			data.URL = fmt.Sprintf("https://cumcordplugins.github.io/Condom/%s", plugin.URL)

			for current <= rangeStart+rangeLength {
				fmt.Println(current)
				u := url.URL{
					Scheme: "ws",
					Host:   fmt.Sprintf("127.0.0.1:%d", current),
					Path:   "/cumcord",
				}
				d, _ := json.Marshal(&data)
				c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
				current += 1
				if err != nil {
					continue
				}
				if err = c.WriteMessage(websocket.TextMessage, d); err != nil {
					continue
				}
				_, message, err := c.ReadMessage()
				if err != nil {
					continue
				}
				defer c.Close()
				json.Unmarshal([]byte(message), &returnData)
				fmt.Println(returnData.Status)
				if returnData.Status == "OK" {
					app.ShowPluginView(back, plugin, nil)
					break
				} else {
					warnings := []framework.FlexboxSlot{
						{
							Element: design.InformationPanel(design.InformationPanelDetails{
								Text: "Something went wrong!",
							}),
						},
					}
					app.ShowPluginView(back, plugin, warnings)
					break
				}
			}
		}),
	}

	slots = append(slots, framework.FlexboxSlot{
		Element: design.ButtonBar(buttons),
	})

	app.Teleport(design.LayoutDocument(design.Header{
		Title: plugin.Name,
		Back:  back,
	}, framework.NewUIFlexboxContainerPtr(framework.FlexboxContainer{
		DirVertical: true,
		Slots:       slots,
	}), true))
}
