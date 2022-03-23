package src

import (
	"path/filepath"
	"sort"

	"github.com/Cumcord/impregnate/middle"
	"github.com/lexisother/frenyard/design"
	"github.com/lexisother/frenyard/framework"
)

func (app *UpApplication) ShowDiscordFinder(back framework.ButtonBehavior, vfsPath string) {
	var vfsList []*middle.DiscordInstance

	app.ShowWaiter("Reading", func(progress func(string)) {
		progress("Scanning to find all of the context in:\n" + vfsPath)
		vfsList = middle.DiscordFinderVFSList(vfsPath)
	}, func() {
		items := []design.ListItemDetails{}

		for _, v := range vfsList {
			thisLocation := v.Path
			ild := design.ListItemDetails{
				Icon: design.DirectoryIconID,
				Text: filepath.Base(thisLocation),
			}
			ild.Click = func() {
				app.GSRightwards()
				app.ShowDiscordFinder(func() {
					app.GSLeftwards()
					app.ShowDiscordFinder(back, vfsPath)
				}, thisLocation)
			}
			if v.Valid {
				ild.Click = func() {
					app.GSRightwards()
					app.ResetWithDiscordInstance(true, thisLocation)
				}
				ild.Text = "Discord " + v.Channel
				ild.Subtext = thisLocation
				ild.Icon = design.GameIconID
			} else if v.Path != "" {
				ild.Text = v.Path
				ild.Subtext = v.Path
				ild.Icon = design.DriveIconID
			}
			items = append(items, ild)
		}

		sort.Sort(design.SortListItemDetails(items))
		primary := design.LayoutDocument(design.Header{
			Back:  back,
			Title: "Enter Discord's location",
		}, design.NewUISearchBoxPtr("Directory name...", items), true)
		app.Teleport(primary)
	})
}
