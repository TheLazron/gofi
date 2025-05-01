package tui

import (
	"fmt"
	"log"

	"github.com/TheLazron/gofi/cmd/tea/tui/constants"
	"github.com/TheLazron/gofi/gofi"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var cmd tea.Cmd

func (c Connector) Init() tea.Cmd {
	return textinput.Blink
}

func InitConnector(selectedItem Connection, gofiClient *gofi.Client) *Connector {
	c := Connector{selectedItem: selectedItem}

	c.gofiClient = gofiClient

	var ti textinput.Model
	if selectedItem.IsProtected {

		ti = textinput.New()
		ti.Width = 20
		ti.CharLimit = 100
		ti.Placeholder = "Password"
		ti.Prompt = fmt.Sprintf("Enter password to connect to %s ", selectedItem.Name)
		ti.Focus()
	}

	c.passwordInput = ti

	return &c
}

func (c Connector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, constants.Keymap.Select):
			// Send Request to connect to wifi

			log.Printf("Connecting to %s ", c.selectedItem.Name)
			device, err := c.gofiClient.GetWifiClient()
			if err != nil {
				fmt.Errorf("failed to conect: %v", err)
				return nil, tea.Quit
			}
			var connOptions gofi.ConnectionOptions

			if c.selectedItem.IsProtected {
				fmt.Println("Password entered", c.passwordInput.Value())
				connOptions.Password = c.passwordInput.Value()
			}

			res, connected := c.gofiClient.Connect(connOptions, c.selectedItem.AccessPoint, device)

			if connected {
				fmt.Println("Connected to network", res)
				return InitWifi()
			}

			fmt.Println("Failed to connect. Make sure the password is correct and try again.")
			return nil, tea.Quit

		case key.Matches(msg, constants.Keymap.Quit):
			// Go back to network selection screen
			return InitWifi()
		}

	case errMsg:
		c.error = msg.Error()
	}

	c.passwordInput, cmd = c.passwordInput.Update(msg)

	cmds = append(cmds, cmd)
	return c, tea.Batch(cmds...)
}

func (c Connector) View() string {

	if c.selectedItem.IsProtected {

		return fmt.Sprintf(
			"\n\n%s\n\n%s\n%s",
			"",
			c.passwordInput.View(),
			"(Esc to return)",
		)
	}

	return fmt.Sprintf(
		"\n%s\n\n%s\n%s",
		"",
		fmt.Sprintf("Press enter to confirm connection to %s", c.selectedItem.Name),
		"(Esc to return)",
	)
}
