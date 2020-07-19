package model

import (
	"time"
)

type Appliance struct {
	ID         string          `json:"id"`
	Device     ApplianceDevice `json:"device,omitempty"`
	Model      Model           `json:"model,omitempty"`
	Type       string          `json:"type"`
	Nickname   string          `json:"nickname"`
	Image      string          `json:"image"`
	Settings   interface{}     `json:"settings"`
	Aircon     interface{}     `json:"aircon"`
	Signals    []interface{}   `json:"signals"`
	SmartMeter SmartMeter      `json:"smart_meter,omitempty"`
	Light      Light           `json:"light,omitempty"`
}

type ApplianceDevice struct {
	Name              string    `json:"name"`
	ID                string    `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	MacAddress        string    `json:"mac_address"`
	BtMacAddress      string    `json:"bt_mac_address"`
	SerialNumber      string    `json:"serial_number"`
	FirmwareVersion   string    `json:"firmware_version"`
	TemperatureOffset int       `json:"temperature_offset"`
	HumidityOffset    int       `json:"humidity_offset"`
}

type Model struct {
	ID           string `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
	Image        string `json:"image"`
}

type SmartMeter struct {
	EchonetliteProperties []struct {
		Name      string    `json:"name"`
		Epc       int       `json:"epc"`
		Val       string    `json:"val"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"echonetlite_properties"`
}

type Light struct {
	Buttons []struct {
		Name  string `json:"name"`
		Image string `json:"image"`
		Label string `json:"label"`
	} `json:"buttons"`
	State struct {
		Brightness string `json:"brightness"`
		Power      string `json:"power"`
		LastButton string `json:"last_button"`
	} `json:"state"`
}
