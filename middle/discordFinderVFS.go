package middle

func DiscordFinderVFSList(vfsPath string) []*DiscordInstance {
	vfsEntries := []*DiscordInstance{}
	for _, fi := range BrowserVFSList(vfsPath) {
		if fi.Dir {
			// cdl, err := NewDiscordInstance(fi.Location)
			cdl := CheckDiscordLocation(fi.Location)
			// if err == nil {
			cdl.Path = fi.Location
			vfsEntries = append(vfsEntries, cdl)
			// }
		}
	}
	return vfsEntries
}
