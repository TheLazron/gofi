package gofi

import "github.com/Wifx/gonetworkmanager"

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
	IsActive    bool
	Protected   bool
	Strength    string
	Name        string
	AccessPoint gonetworkmanager.AccessPoint
}

func NewConnection(conn WiFiNetwork, accessPoint gonetworkmanager.AccessPoint) Connection {

	//WPA and RSN equals 0 implies the network isn't protected
	return Connection{
		IsActive:    false,
		Protected:   IsConnProtected(conn),
		Strength:    getStrength(conn.Strength),
		Name:        conn.SSID,
		AccessPoint: accessPoint,
	}
}

func IsConnProtected(conn WiFiNetwork) bool {
	return conn.WPAFlags != 0 || conn.RSNFlags != 0
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
