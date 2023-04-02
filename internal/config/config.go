package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// validKey regex to validate config key
var validKey = regexp.MustCompile("^[[:alnum:]]")

// Config represent the configs
type Config struct {
	ApiToken    string
	WorkspaceID string
}

// Write write config to local file
func Write(cfg Config) error {
	filename, err := filename()
	if err != nil {
		return err
	}

	var content string
	content += fmt.Sprintf("apitoken=%s\n", cfg.ApiToken)
	content += fmt.Sprintf("workspace_id=%s\n", cfg.WorkspaceID)

	return os.WriteFile(filename, []byte(content), 0400)
}

// Read read config from local file
func Read() (*Config, error) {
	filename, err := filename()
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if file == nil {
		return nil, fmt.Errorf("file not found")
	}

	return Parse(file), nil
}

// Parse parse content of file and create config
func Parse(content []byte) *Config {
	cfg := new(Config)
	for _, l := range strings.Split(string(content), "\n") {
		l = strings.TrimSpace(l)
		if validKey.MatchString(l) && strings.Contains(l, "=") {
			p := strings.Index(l, "=")
			key, value := l[:p], l[p+1:]
			switch key {
			case "apitoken":
				cfg.ApiToken = value
			case "workspace_id":
				cfg.WorkspaceID = value
			}
		}

	}
	return cfg
}

// filename return the filename of file
func filename() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/.timetrack", home), nil
}
