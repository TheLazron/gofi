package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TheLazron/gofi/gofi"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Model struct {
	list list.Model
}

type item struct {
	name, strength, isActive string
}

func (i item) Title() string       { return i.name }
func (i item) Description() string { return fmt.Sprintf("Strength: %s | isActive: False", i.strength) }

// func (i item) IsActive() string    { return i.isActive }
func (i item) FilterValue() string { return i.name }

func (m Model) View() string {
	return docStyle.Render(m.list.View())
}

func (m Model) Init() tea.Cmd {
	return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
func main() {
	client, err := gofi.New()
	if err != nil {
		log.Println("Failed to initialize client")
	}

	connections, err := client.ListDevices()
	if err != nil {
		log.Printf("Failed to list devices: %v", err)
	}

	items := make([]list.Item, 0, 20)

	for _, conn := range connections {
		items = append(items, item{
			name:     conn.Name,
			strength: conn.Strength,
			isActive: "False",
		})
	}

	m := Model{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}

	m.list.Title = "Available Connections"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
