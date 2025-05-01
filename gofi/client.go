package gofi

import (
	"encoding/json"
	"fmt"

	"github.com/Wifx/gonetworkmanager"
)

type Client struct {
	gonetworkmanager.NetworkManager
	Device gonetworkmanager.DeviceWireless
}

func New() (*Client, error) {
	networkManager, err := gonetworkmanager.NewNetworkManager()

	if err != nil {
		return nil, err
	}

	return &Client{
		NetworkManager: networkManager,
	}, nil
}

func (c *Client) GetWifiClient() (gonetworkmanager.DeviceWireless, error) {
	var wirelessDevice gonetworkmanager.DeviceWireless
	//@TODO: Check for availability of the wifi device

	devices, err := c.GetAllDevices()
	if err != nil {
		return nil, err
	}

	for _, device := range devices {
		deviceProperty, err := device.GetPropertyDeviceType()
		if err != nil {
			return nil, err
		}

		if deviceProperty == gonetworkmanager.NmDeviceTypeWifi {
			wireless, err := gonetworkmanager.NewDeviceWireless(device.GetPath())
			if err != nil {
				return nil, err
			}
			wirelessDevice = wireless
		}
	}
	c.Device = wirelessDevice
	return wirelessDevice, nil
}

func (c *Client) ListDevices() ([]Connection, error) {
	conns := make([]Connection, 0)
	wirelessDevice, err := c.GetWifiClient()
	if err != nil {
		return nil, err
	}

	err = wirelessDevice.RequestScan()
	if err != nil {
		return nil, err
	}

	connections, err := wirelessDevice.GetPropertyAccessPoints()
	if err != nil {
		return nil, err
	}

	for _, connection := range connections {
		conn, err := connection.MarshalJSON()
		if err != nil {
			continue
		}
		var acsPoint WiFiNetwork
		err = json.Unmarshal(conn, &acsPoint)
		if err != nil {
			continue
		}
		c := NewConnection(acsPoint, connection)
		conns = append(conns, c)
	}
	return conns, nil
}

type ConnectionOptions struct {
	Password string
}

func (c *Client) Connect(connectionOptions ConnectionOptions, accessPoint gonetworkmanager.AccessPoint, device gonetworkmanager.DeviceWireless) (string, bool) {

	fmt.Printf("Is protected: %+v", connectionOptions)
	//Check if protected
	var wifiConn WiFiNetwork
	conn, err := accessPoint.MarshalJSON()
	if err != nil {
		return "Failed to establish connection", false
	}

	err = json.Unmarshal(conn, &wifiConn)
	if err != nil {
		return "Failed to establish conection", false
	}

	isProtected := IsConnProtected(wifiConn)
	connection := make(map[string]map[string]any)
	if isProtected {
		connection["802-11-wireless"] = make(map[string]any)
		connection["802-11-wireless"]["security"] = "802-11-wireless-security"
		connection["802-11-wireless-security"] = make(map[string]any)
		connection["802-11-wireless-security"]["key-mgmt"] = "sae wpa-psk"
		connection["802-11-wireless-security"]["psk"] = connectionOptions.Password
	}

	activeConnection, err := c.AddAndActivateWirelessConnection(connection, device, accessPoint)
	fmt.Println("Active connection", activeConnection)
	if err != nil {
		fmt.Println("Failed to establish connection", err)
		return "Failed to establish connection", false
	}
	cn, _ := activeConnection.GetPropertyConnection()
	j, _ := cn.MarshalJSON()
	return fmt.Sprintf("Connected to %+v", string(j)), true

}
