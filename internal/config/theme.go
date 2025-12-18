package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Theme struct {
	// Backgrounds
	Background    string `json:"background"`
	BackgroundAlt string `json:"background_alt"`

	// Foregrounds
	TextPrimary   string `json:"text_primary"`
	TextSecondary string `json:"text_secondary"`
	TextMuted     string `json:"text_muted"`

	// Status/State
	Error   string `json:"error"`
	Warning string `json:"warning"`
	Success string `json:"success"`
	Info    string `json:"info"`

	// Priority indicators
	PriorityHigh   string `json:"priority_high"`
	PriorityMedium string `json:"priority_medium"`
	PriorityLow    string `json:"priority_low"`

	// UI Elements
	Accent          string `json:"accent"`           // For highlights, cursor
	AccentSecondary string `json:"accent_secondary"` // For selected items
	Border          string `json:"border"`
}

type Config struct {
	Theme Theme `json:"theme"`
}

// default theme is Gruvbox
func DefaultTheme() Theme {
	return Theme{
		Background:      "#3c3836",
		BackgroundAlt:   "#282828", // or use for modals
		TextPrimary:     "#ebdbb2",
		TextSecondary:   "#C5C7BC",
		TextMuted:       "#a89984",
		Error:           "#cc241d",
		Warning:         "#fe8019",
		Success:         "#b8bb26",
		Info:            "#83a598",
		PriorityHigh:    "#fb4934",
		PriorityMedium:  "#fe8019",
		PriorityLow:     "#b8bb26",
		Accent:          "#fabd2f", // cursor, focused elements
		AccentSecondary: "#83a598", // selected items
		Border:          "#a89984",
	}
}

func LoadConfig() (*Config, error) {
	configPath, err := ConfigFilePath()
	if err != nil {
		return &Config{Theme: DefaultTheme()}, nil
	}

	// Check if config file exists
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Config doesn't exist, create it with defaults
			return createDefaultConfig(configPath)
		}
		// Other read error
		return nil, err
	}

	// Parse existing config
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("invalid config file: %w", err)
	}

	// Fill in any missing values
	cfg.Theme = mergeWithDefaults(cfg.Theme)
	return &cfg, nil
}

func createDefaultConfig(configPath string) (*Config, error) {
	cfg := &Config{Theme: DefaultTheme()}
	permissions := 0755

	// Ensure directory exists
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, os.FileMode(permissions)); err != nil {
		return cfg, fmt.Errorf("failed to create config directory: %w", err)
	}

	// Write default config
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return cfg, err
	}

	writePermission := 0644 //read write file or dir, others can only read
	if err := os.WriteFile(configPath, data, os.FileMode(writePermission)); err != nil {
		return cfg, fmt.Errorf("failed to write config file: %w", err)
	}

	return cfg, nil
}

func mergeWithDefaults(theme Theme) Theme {
	defaults := DefaultTheme()
	if theme.Background == "" {
		theme.Background = defaults.Background
	}
	if theme.BackgroundAlt == "" {
		theme.BackgroundAlt = defaults.BackgroundAlt
	}
	if theme.TextPrimary == "" {
		theme.TextPrimary = defaults.TextPrimary
	}
	if theme.TextSecondary == "" {
		theme.TextSecondary = defaults.TextSecondary
	}
	if theme.TextMuted == "" {
		theme.TextMuted = defaults.TextMuted
	}
	if theme.Error == "" {
		theme.Error = defaults.Error
	}
	if theme.Warning == "" {
		theme.Warning = defaults.Warning
	}
	if theme.Success == "" {
		theme.Success = defaults.Success
	}
	if theme.Info == "" {
		theme.Info = defaults.Info
	}
	if theme.PriorityHigh == "" {
		theme.PriorityHigh = defaults.PriorityHigh
	}
	if theme.PriorityMedium == "" {
		theme.PriorityMedium = defaults.PriorityMedium
	}
	if theme.PriorityLow == "" {
		theme.PriorityLow = defaults.PriorityLow
	}
	if theme.Accent == "" {
		theme.Accent = defaults.Accent
	}
	if theme.AccentSecondary == "" {
		theme.AccentSecondary = defaults.AccentSecondary
	}
	if theme.Border == "" {
		theme.Border = defaults.Border
	}
	return theme
}
