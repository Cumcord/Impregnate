package main

import (
  "github.com/20kdc/CCUpdaterUI/design"
  "github.com/20kdc/CCUpdaterUI/frenyard"
  "github.com/20kdc/CCUpdaterUI/frenyard/framework"
	"github.com/Cumcord/impregnate/src"
)

func main() {
	frenyard.TargetFrameTime = 0.016
	slideContainer := framework.NewUISlideTransitionContainerPtr(nil)
	slideContainer.FyEResize(design.SizeWindowInit)
	wnd, err := framework.CreateBoundWindow("Impregnate", true, design.ThemeBackground, slideContainer)
	if err != nil {
		panic(err)
	}
	design.Setup(frenyard.InferScale(wnd))
	wnd.SetSize(design.SizeWindow)
	// Ok, now get it ready.
	app := (&src.UpApplication{
		MainContainer: slideContainer,
		Window: wnd,
		UpQueued: make(chan func(), 16),
		TeleportSettings: framework.SlideTransition{},
	})
	// Started!
	frenyard.GlobalBackend.Run(func(frameTime float64) {
		select {
			case fn := <- app.UpQueued:

				fn()
			default:
		}
	})
}
