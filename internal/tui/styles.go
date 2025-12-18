package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/kwame-Owusu/lista/internal/config"
)

var (
	// Background tones
	bg1  = lipgloss.Color("#3c3836") // dark1 :contentReference[oaicite:4]{index=4}
	gray = lipgloss.Color("#C5C7BC") // light gray

	// Foreground tones
	fg      = lipgloss.Color("#ebdbb2") // light foreground
	fgMuted = lipgloss.Color("#a89984") // muted foreground

	// Accent colors (neutral)
	red = lipgloss.Color("#cc241d")

	// Bright accents
	brightRed    = lipgloss.Color("#fb4934") // bright red :contentReference[oaicite:14]{index=14}
	brightGreen  = lipgloss.Color("#b8bb26") // bright green :contentReference[oaicite:15]{index=15}
	brightYellow = lipgloss.Color("#fabd2f") // bright yellow :contentReference[oaicite:16]{index=16}
	brightBlue   = lipgloss.Color("#83a598") // bright blue :contentReference[oaicite:17]{index=17}
	brightOrange = lipgloss.Color("#fe8019") // bright orange :contentReference[oaicite:20]{index=20}
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(brightYellow).
			Bold(true).
			MarginBottom(1)

	modalStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(fgMuted).
			Foreground(fg).
			Padding(1, 3)

	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(gray).
			Background(brightBlue).
			Padding(0, 1)

	// Completed + selected
	completedSelectedStyle = lipgloss.NewStyle().
				Foreground(fgMuted).
				Background(bg1).
				Strikethrough(true).
				Padding(0, 1)

	// Normal item
	itemStyle = lipgloss.NewStyle().
			Foreground(fg).
			Padding(0, 1)

	completedStyle = lipgloss.NewStyle().
			Foreground(fgMuted).
			Strikethrough(true).
			Padding(0, 1)

	// Priority styles
	highPriorityStyle = lipgloss.NewStyle().
				Foreground(brightRed).
				Bold(true)

	mediumPriorityStyle = lipgloss.NewStyle().
				Foreground(brightOrange).
				Bold(true)

	lowPriorityStyle = lipgloss.NewStyle().
				Foreground(brightGreen)

	// Help text
	helpStyle = lipgloss.NewStyle().
			Foreground(fgMuted).
			MarginTop(1)

	// Cursor indicator
	cursorStyle = lipgloss.NewStyle().
			Foreground(brightYellow).
			Bold(true)

	// Error messages
	errorStyle = lipgloss.NewStyle().
			Foreground(red).
			Bold(true)
)

func getPriorityStyle(priority string) lipgloss.Style {
	switch priority {
	case "High":
		return highPriorityStyle
	case "Medium":
		return mediumPriorityStyle
	default:
		return lowPriorityStyle
	}
}
