package tui

import (
	"github.com/TheLazron/gofi/gofi"
	"github.com/Wifx/gonetworkmanager"
	"github.com/charmbracelet/bubbles/textinput"
)

type WiFiNetwork struct {
	Flags      int    `json:"Flags"`
	Frequency  int    `json:"Frequency"`
	HWAddress  string `json:"HWAddress"`
	LastSeen   int    `json:"LastSeen"`
	MaxBitrate int    `json:"MaxBitrate"`
	Mode       string `json:"Mode"`
	RSNFlags   int    `json:"RSNFlags"`
	SSID       string `json:"SSID"`
	Strength   uint8  `json:"Strength"`
	WPAFlags   int    `json:"WPAFlags"`
}

type Connector struct {
	gofiClient *gofi.Client
	gonetworkmanager.NetworkManager
	selectedItem  Connection
	passwordInput textinput.Model
	error         string
}
