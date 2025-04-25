package gofi

import (
	"github.com/Wifx/gonetworkmanager"
)

const (
	StrengthPoor      = "Poor"
	StrengthFair      = "Fair"
	StrengthGood      = "Good"
	StrengthExcellent = "Excellent"
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

type Connection struct {
	IsActive bool
	Strength string
	gonetworkmanager.Device
	Name string
}

func NewConnection(conn WiFiNetwork) Connection {
	return Connection{
		IsActive: false,
		Strength: getStrength(conn.Strength),
		Name:     conn.SSID,
	}
}

func getStrength(strength uint8) string {
	switch {
	case strength > 0 && strength <= 30:
		return StrengthPoor
	case strength > 30 && strength <= 50:
		return StrengthFair
	case strength > 50 && strength <= 70:
		return StrengthGood
	case strength > 70 && strength <= 100:
		return StrengthExcellent
	default:
		return "----"
	}
}
