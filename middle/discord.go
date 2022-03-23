package middle

// Authored by: Creatable

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type DiscordInstance struct {
	Path    string
	Channel string
	Valid   bool
}

func GetInstance(channel string) (DiscordInstance, error) {
	channelString := "Discord"

	// Generate channel strings (e.g discord-canary, DiscordCanary, Discord Canary)
	if channel != "Stable" {
		switch os := runtime.GOOS; os {
		case "darwin":
			channelString = channelString + " " + channel
		case "windows":
			channelString = channelString + channel
		default: // Linux and BSD are basically the same thing
			channelString = channelString + "-" + channel
		}
	}

	instance := DiscordInstance{
		Path:    "",
		Channel: channel,
		Valid:   false,
	}

	switch OS := runtime.GOOS; OS {
	case "darwin":
		instance.Path = filepath.Join("/Applications", channelString+".app", "Contents", "Resources")
	case "windows":
		starterPath := filepath.Join(os.Getenv("localappdata"), channelString, "/")
		filepath.Walk(starterPath, func(path string, _ fs.FileInfo, _ error) error {

			if strings.HasPrefix(filepath.Base(path), "app-") {
				instance.Path = filepath.Join(path, "resources")
			}

			return nil
		})
	default: // Linux and BSD are *still* basically the same thing
		channels := []string{channelString, strings.ToLower(channelString)}
		path := os.Getenv("PATH")

		for _, channel := range channels {
			for _, pathItem := range strings.Split(path, ":") {
				joinedPath := filepath.Join(pathItem, channel)
				if _, err := os.Stat(joinedPath); err == nil {
					possiblepath, _ := filepath.EvalSymlinks(joinedPath)
					if possiblepath != joinedPath {
						instance.Path = filepath.Join(possiblepath, "..", "resources")
					}
				}
			}
		}

		if instance.Path == "" && channel == "Stable" {
			instance.Path = "/var/lib/flatpak/app/com.discordapp.Discord/x86_64/stable/active/files/discord/resources/"
		}
	}

	if _, err := os.Stat(instance.Path); err == nil {
		return instance, nil
	} else {
		return instance, errors.New("Instance doesn't exist")
	}
}

func GetChannels() []DiscordInstance {
	possible := []string{"Stable", "PTB", "Canary", "Development"}
	var channels []DiscordInstance

	for _, channel := range possible {
		c, err := GetInstance(channel)
		if err == nil {
			channels = append(channels, c)
		}
	}

	return channels
}

func NewDiscordInstance(path string) (*DiscordInstance, error) {
	instance := DiscordInstance{
		Path:    "",
		Channel: "Unknown",
		Valid:   false,
	}

	instance.Path = path

	if _, err := os.Stat(filepath.Join(instance.Path, "app.asar")); err == nil {
		return &instance, nil
	} else {
		return &instance, errors.New("Instance doesn't exist")
	}
}

func CheckDiscordLocation(dir string) *DiscordInstance {
	if BrowserVFSLocationReal(dir) {
		discordInstance, err := NewDiscordInstance(dir)
		if err == nil {
			fmt.Print("Discord instance found at " + dir + "\n")
			if _, err := os.Stat(discordInstance.Path); err == nil {
				discordInstance.Valid = true
			}
		}
	}

	return &DiscordInstance{
		Path: dir,
	}
}

