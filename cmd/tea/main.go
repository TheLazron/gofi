package main

import (
	"log"

	"github.com/TheLazron/gofi/cmd/tea/tui"
)

func main() {
	err := tui.StartTea()
	if err != nil {
		log.Fatal(err)
	}
}
