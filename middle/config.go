package middle

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DiscordPath string `json:"discordPath"`
}

func getConfigPath() string {
	cfg, err := os.UserConfigDir()
	if err != nil {
		cfg = ""
	}
	return filepath.Join(cfg, "impregnate.json")
}

func ReadConfig() Config {
	cfg := &Config{}
	cfgFile, err := os.Open(getConfigPath())
	if err != nil {
		fmt.Printf("Failed to open config: %s\n", err.Error())
	}
	defer cfgFile.Close()
	if json.NewDecoder(cfgFile).Decode(cfg) != nil {
		fmt.Printf("Failed to decode config\n")
		return *cfg
	}
	WriteConfig(*cfg)
	return *cfg
}

func WriteConfig(cfg Config) {
	cfgPath := getConfigPath()
	os.MkdirAll(filepath.Dir(cfgPath), os.ModePerm)
	cfgFile, err := os.OpenFile(cfgPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to write config: %s\n", err.Error())
		return
	}
	defer cfgFile.Close()
	if json.NewEncoder(cfgFile).Encode(cfg) != nil {
		fmt.Printf("Failed to encode config\n")
	}
}
