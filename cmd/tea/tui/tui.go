package tui

import (
	"fmt"
	"log"
	"os"

	"github.com/TheLazron/gofi/cmd/tea/tui/constants"
	tea "github.com/charmbracelet/bubbletea"
)

func StartTea() error {
	if f, err := tea.LogToFile("debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging")
		os.Exit(1)
	} else {
		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	p, _ := InitWifi()

	constants.P = tea.NewProgram(p, tea.WithAltScreen())

	if _, err := constants.P.Run(); err != nil {
		fmt.Println("Error running terminal", err)
		os.Exit(1)
	}
	return nil
}
