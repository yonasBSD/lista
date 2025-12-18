package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// DataFilePath returns the OS-appropriate config file path for lista
func DataFilePath() (string, error) {
	var baseDir string
	var err error

	switch runtime.GOOS {
	case "windows":
		baseDir, err = os.UserConfigDir()
	case "darwin":
		baseDir, err = os.UserConfigDir()
	default:
		// Linux / BSD / Unix
		baseDir = os.Getenv("XDG_CONFIG_HOME")
		if baseDir == "" {
			baseDir, err = os.UserConfigDir()
		}
	}

	if err != nil {
		return "", err
	}

	configDir := filepath.Join(baseDir, "lista")
	return filepath.Join(configDir, "lista.json"), nil
}

// ConfigFilePath returns the OS-appropriate config file path for lista
func ConfigFilePath() (string, error) {
	var baseDir string
	var err error
	switch runtime.GOOS {
	case "windows":
		baseDir, err = os.UserConfigDir()
	case "darwin":
		baseDir, err = os.UserConfigDir()
	default:
		// Linux / BSD / Unix
		baseDir = os.Getenv("XDG_CONFIG_HOME")
		if baseDir == "" {
			baseDir, err = os.UserConfigDir()
		}
	}
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(baseDir, "lista")
	return filepath.Join(configDir, "lista.config.json"), nil
}
