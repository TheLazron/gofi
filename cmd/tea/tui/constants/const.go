package constants

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	P          *tea.Program
	WindowSize tea.WindowSizeMsg
)

const (
	StrengthPoor      = "Poor"
	StrengthFair      = "Fair"
	StrengthGood      = "Good"
	StrengthExcellent = "Excellent"
)

// Global styling for connections
var DocStyle = lipgloss.NewStyle().Margin(0, 2)

// HelpStyle stylings for help cotext menu
var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

// Provide styling for error messages
var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render()

// Provides stylig for Alert messages
var AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62"))

type keymap struct {
	Select key.Binding
	Quit   key.Binding
}

var Keymap = keymap{
	Select: key.NewBinding(
		key.WithKeys(tea.KeyEnter.String()),
		key.WithHelp("Enter", "Select a network to connect to "),
	),
	Quit: key.NewBinding(
		key.WithKeys(tea.KeyEsc.String()),
		key.WithHelp("Esc", "Abort and Returnn"),
	),
}
