package remo_api

import (
	"encoding/json"
	"errors"
	"github.com/mizuho1998/grafana_monitor/app/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetAppliances() ([]model.Appliance, error) {
	TOKEN := os.Getenv("TOKEN")
	url := "https://api.nature.global/1/appliances"

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

	var appliances []model.Appliance
	err = json.Unmarshal(body, &appliances)
	if err != nil {
		log.Println(res.Status)
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
