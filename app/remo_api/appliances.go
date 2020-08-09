package remo_api

import (
	"encoding/json"
	"errors"
	"github.com/mizuho1998/grafana_monitor/app/model"
	"log"
	"strconv"
	"time"
)

func GetAppliances() ([]model.Appliance, error) {
	url := "https://api.nature.global/1/appliances"

	body, err := GetResponseBody(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var appliances []model.Appliance
	err = json.Unmarshal(body, &appliances)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return appliances, nil
}

func GetSmartMeter(appliances []model.Appliance) (*model.SmartMeter, error) {
	smartMeter := appliances[0].SmartMeter

	return &smartMeter, nil
}

func GetInstantaneous() (int, time.Time, error) {
	appliances, err := GetAppliances()
	if err != nil {
		return 0, time.Unix(0, 0), nil
	}
	smartMeter, err := GetSmartMeter(appliances)

	for _, v := range smartMeter.EchonetliteProperties {
		if v.Name == "measured_instantaneous" {
			val, _ := strconv.Atoi(v.Val)
			return val, v.UpdatedAt, nil
		}
	}

	return 0, time.Unix(0, 0), errors.New("not found measured_instantaneous in smartMeter.EchonetliteProperties.")
}
