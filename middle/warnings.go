package middle

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

func FindWarnings() []Warning {
	warnings := []Warning{}

	health := CheckHealth()

	if (ReturnData{}) == health {
		warnings = append(warnings, Warning{
			Text:      "Cumcord is not installed! (or Discord is not running)",
			Action:    URLAndCloseWarningID,
			Parameter: "https://cumcord.com",
		})
	}

	return warnings
}
