package remo_api

import (
	"encoding/json"
	"errors"
	"github.com/mizuho1998/grafana_monitor/app/model"
	"log"
	"time"
)

func GetDevices() ([]model.Device, error) {
	url := "https://api.nature.global/1/devices"

	body, err := GetResponseBody(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var devices []model.Device
	err = json.Unmarshal(body, &devices)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return devices, nil
}

func GetNewestEvents() (*model.NewestEvents, error) {
	devices, err := GetDevices()
	if err != nil {
		return nil, err
	}

	zero_time := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, v := range devices {
		if v.NewestEvents.Hu.CreatedAt != zero_time {
			return &v.NewestEvents, nil
		}
	}

	return nil, errors.New("not found newest_events")
}
