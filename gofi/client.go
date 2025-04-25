package gofi

import (
	"encoding/json"

	"github.com/Wifx/gonetworkmanager"
)

type Client struct {
	gonetworkmanager.NetworkManager
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
		c := NewConnection(acsPoint)
		conns = append(conns, c)
	}
	return conns, nil
}

// func (c *Client) ListDevices() ([]Connection, error) {
// 	dcs, _ := c.GetAllDevices()
// 	for _, dc := range dcs {
// 		a, _ := dc.GetPropertyDeviceType()
// 		if a == gonetworkmanager.NmDeviceTypeWifi {
// 			// conns, _ := dc.GetPropertyActiveConnection()
// 			wireless, _ := gonetworkmanager.NewDeviceWireless(dc.GetPath())
// 			wireless.RequestScan()

// 			acspts, _ := wireless.GetPropertyAccessPoints()

// 			for _, acspt := range acspts {
// 				a, _ := acspt.MarshalJSON()
// 				fmt.Printf("%+v", string(a))
// 			}

// 			// g, _ := conns.GetPropertySpecificObject()
// 			// h, _ := g.MarshalJSON()
// 			// fmt.Printf("hr %v", string(h))

// 			// for _, conn := range conns {
// 			// a, _ := conn.
// 			// fmt.Printf("conns: %+v", a)
// 			// }
// 		}
// 	}
// 	devices, err := c.GetPropertyActiveConnections()
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(devices)
// 	var conns []Connection
// 	fmt.Println(len(devices))
// 	for _, device := range devices {
// 		device.GetPropertyUUID()
// 		a, _ := c.GetPropertyPrimaryConnection()
// 		v, err := a.GetPropertySpecificObject()
// 		f, _ := v.GetPropertySSID()
// 		fmt.Printf("Active %+v\n", f)
// 		// _, err := device.GetPropertyState()
// 		if err != nil {
// 			log.Println("Failed to load device", device.GetPath())
// 			continue
// 		}
// 		// p, _ := device.GetPropertyActiveConnection()
// 		// if p != nil {

// 		// fmt.Println(p.GetPropertyID())
// 		// }
// 		// conns = append(conns, Connection{
// 		// 	isActive: state == gonetworkmanager.NmDeviceStateActivated,
// 		// 	Device:   device,
// 		// })
// 	}
// 	return conns, nil
// }
