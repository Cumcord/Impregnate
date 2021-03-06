package main

import (
	"github.com/Cumcord/impregnate/middle"
	"github.com/Cumcord/impregnate/src"
	"github.com/lexisother/frenyard"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
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
		Config:           middle.ReadConfig(),
		MainContainer:    slideContainer,
		Window:           wnd,
		UpQueued:         make(chan func(), 16),
		TeleportSettings: framework.SlideTransition{},
	})
	app.ShowPreface()
	// Started!
	frenyard.GlobalBackend.Run(func(frameTime float64) {
		select {
		case fn := <-app.UpQueued:

			fn()
		default:
		}
	})
}
