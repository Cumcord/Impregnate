package middle

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

type WarningID int

const (
	// NullActionWarningID cannot be automatically fixed.
	NullActionWarningID WarningID = iota
	// InstallOrUpdatePackageWarningID warnings can be solved by installing/updating the package Parameter.
	InstallOrUpdatePackageWarningID
	// URLAndCloseWarningID warnings can be solved manually by the user given navigation to a URL. The application closes.
	URLAndCloseWarningID
)

// Warning represents a warning to show the user on the primary view.
type Warning struct {
	Text      string
	Action    WarningID
	Parameter string
}

// Initialize some variables for storing version checking information
var remoteVersion string
var hasAlreadyCheckedUpdate = false

func checkUpdate() {
	if !hasAlreadyCheckedUpdate {
		hasAlreadyCheckedUpdate = true
	} else {
		return
	}

	res, _ := http.Get("https://raw.githubusercontent.com/Cumcord/Impregnate/master/middle/version.go")

	data, _ := ioutil.ReadAll(res.Body)
	remoteVersion = strings.Trim(string(data)[33:38], "\r\n")
}

func FindWarnings(config Config) []Warning {
	warnings := []Warning{}

	// health := CheckHealth()
	_, plugErr := os.Stat(path.Join(config.DiscordPath, "resources/app/plugged.txt"))

	if plugErr != nil {
		// if (ReturnData{}) == health {
		warnings = append(warnings, Warning{
			Text:      "Cumcord is not installed! (or Discord is not running)",
			Action:    InstallOrUpdatePackageWarningID,
			Parameter: "https://cumcord.com",
		})
	}

	if !hasAlreadyCheckedUpdate {
		checkUpdate()
	}
	if remoteVersion != version {
		warnings = append(warnings, Warning{
			Text:      "A new version of Impregnate is available! (v" + remoteVersion + ")",
			Action:    URLAndCloseWarningID,
			Parameter: "https://github.com/Cumcord/Impregnate/releases",
		})
	}

	return warnings
}
