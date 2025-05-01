package tui

import (
	"fmt"

	"github.com/TheLazron/gofi/cmd/tea/tui/constants"
	"github.com/TheLazron/gofi/gofi"
	"github.com/Wifx/gonetworkmanager"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg struct{ error }
)

type Connections struct {
	list list.Model
	gonetworkmanager.NetworkManager
	gofiClient *gofi.Client
}

type Connection struct {
	AccessPoint gonetworkmanager.AccessPoint
	Name        string
	IsActive    bool
	IsProtected bool
	Strength    string
}

func (i Connection) Title() string { return i.Name }
func (i Connection) Description() string {

	securityIndicator := ""
	if i.IsProtected {
		securityIndicator = "P"
	}

	return fmt.Sprintf("Strength: %s | Is Active: %t | %s", i.Strength, i.IsActive, securityIndicator)
}

func (i Connection) FilterValue() string { return i.Name }

func (c Connections) Init() tea.Cmd {
	return nil
}

func InitWifi() (tea.Model, tea.Cmd) {

	gofiClient, err := gofi.New()

	if err != nil {
		return nil, nil
	}

	connections, err := gofiClient.ListDevices()
	if err != nil {
		return nil, nil
	}

	connItems := make([]list.Item, 0, 20)

	for _, conn := range connections {
		connItems = append(connItems, Connection{
			Name:        conn.Name,
			IsActive:    conn.IsActive,
			IsProtected: conn.Protected,
			Strength:    conn.Strength,
			AccessPoint: conn.AccessPoint,
		})
	}
	c := Connections{
		list:       list.New(connItems, list.NewDefaultDelegate(), 0, 0),
		gofiClient: gofiClient,
	}

	c.list.Title = "Available Connections"

	return &c, nil
}

func (c Connections) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Select):
			selectedItem := c.list.SelectedItem().(Connection)
			conn := InitConnector(selectedItem, c.gofiClient)
			return conn.Update(constants.WindowSize)
		case key.Matches(msg, constants.Keymap.Quit):
			fmt.Println("Quitting Program")
			return c, tea.Quit
		default:
			c.list, cmd = c.list.Update(msg)
		}
		cmds = append(cmds, cmd)
	}
	return c, tea.Batch(cmds...)
}

//View returns the text to be output to the terminal

func (c Connections) View() string {
	return constants.DocStyle.Render(c.list.View() + "\n")
}
