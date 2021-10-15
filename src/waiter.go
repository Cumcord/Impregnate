/*
 * CCUpdaterUI/waiter.go
 * Written starting in 2019 by 20kdc
 * This work is licensed under the terms of the MIT license.
 * For a copy, see <https://opensource.org/licenses/MIT>.
 */

package src

import (
	"github.com/lexisother/frenyard"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
	"github.com/lexisother/frenyard/integration"
)

func (app *UpApplication) ShowWaiter(text string, a func(func(string)), b func()) {
	label := framework.NewUILabelPtr(integration.NewTextTypeChunk("", design.GlobalFont), design.ThemeText, 0, frenyard.Alignment2i{})
	app.Teleport(design.LayoutDocument(design.Header{
		Title: text,
	}, label, false))
	go func() {
		a(func(text string) {
			app.UpQueued <- func() {
				label.SetText(integration.NewTextTypeChunk(text, design.GlobalFont))
			}
		})
		app.UpQueued <- b
	}()
}

func (app *UpApplication) MessageBox(title string, text string, b func()) {
	app.Teleport(design.LayoutDocument(design.Header{
		Title: title,
	}, design.LayoutMsgbox(text, b), true))
}
