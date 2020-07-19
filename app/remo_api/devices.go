package remo_api

import (
	"encoding/json"
	"errors"
	"github.com/mizuho1998/grafana_monitor/app/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func GetDevices() ([]model.Device, error) {
	TOKEN := os.Getenv("TOKEN")
	url := "https://api.nature.global/1/devices"

	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+TOKEN)
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var devices []model.Device
	err = json.Unmarshal(body, &devices)
	if err != nil {
		log.Println(err)
		log.Println(res.Status)
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
