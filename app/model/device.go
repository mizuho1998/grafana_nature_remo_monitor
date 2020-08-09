package model

import (
	"time"
)

type Device struct {
	Name              string       `json:"name"`
	ID                string       `json:"id"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	MacAddress        string       `json:"mac_address"`
	BtMacAddress      string       `json:"bt_mac_address,omitempty"`
	SerialNumber      string       `json:"serial_number"`
	FirmwareVersion   string       `json:"firmware_version"`
	TemperatureOffset int          `json:"temperature_offset"`
	HumidityOffset    int          `json:"humidity_offset"`
	Users             []User       `json:"users"`
	NewestEvents      NewestEvents `json:"newest_events,omitempty"`
}

type User struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Superuser bool   `json:"superuser"`
}

type NewestEvents struct {
	Hu struct {
		Val       float64   `json:"val"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"hu"`
	Il struct {
		Val       float64   `json:"val"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"il"`
	Mo struct {
		Val       float64   `json:"val"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"mo"`
	Te struct {
		Val       float64   `json:"val"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"te"`
}
