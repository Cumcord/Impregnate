package src

import (
	"github.com/yellowsink/frenyard"
	"github.com/yellowsink/frenyard/framework"
)

// Code in this file was taken from CCUpdaterUI/main.go

type UpApplication struct {
	MainContainer     *framework.UISlideTransitionContainer
	Window            frenyard.Window
	UpQueued          chan func()
	CachedPrimaryView framework.UILayoutElement
	TeleportSettings  framework.SlideTransition
}

const upTeleportLen float64 = 0.25

// GSLeftwards sets the teleportation affinity to LEFT.
func (app *UpApplication) GSLeftwards() {
	app.TeleportSettings.Reverse = true
	app.TeleportSettings.Vertical = false
	app.TeleportSettings.Length = upTeleportLen
}

// GSRightwards sets the teleportation affinity to RIGHT.
func (app *UpApplication) GSRightwards() {
	app.TeleportSettings.Reverse = false
	app.TeleportSettings.Vertical = false
	app.TeleportSettings.Length = upTeleportLen
}

// GSLeftwards sets the teleportation affinity to UP.
func (app *UpApplication) GSUpwards() {
	app.TeleportSettings.Reverse = true
	app.TeleportSettings.Vertical = true
	app.TeleportSettings.Length = upTeleportLen
}

// GSRightwards sets the teleportation affinity to DOWN.
func (app *UpApplication) GSDownwards() {
	app.TeleportSettings.Reverse = false
	app.TeleportSettings.Vertical = true
	app.TeleportSettings.Length = upTeleportLen
}

// GSInstant sets the teleportation affinity to INSTANT.
func (app *UpApplication) GSInstant() {
	// direction doesn't matter
	app.TeleportSettings.Length = 0
}

// Teleport starts a transition with the cached affinity settings.
func (app *UpApplication) Teleport(target framework.UILayoutElement) {
	forkTD := app.TeleportSettings
	forkTD.Element = target
	app.MainContainer.TransitionTo(forkTD)
}
