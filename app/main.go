package main

import (
	"github.com/influxdata/influxdb-client-go"
	"github.com/joho/godotenv"
	"github.com/mizuho1998/grafana_monitor/app/remo_api"
	"github.com/mizuho1998/grafana_monitor/app/repository"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create client
	URL := os.Getenv("URL")
	client := influxdb2.NewClient("http://"+URL+":8086", "my-token")
	// Get non-blocking write client
	writeAPI := client.WriteAPI("my-org", "mydb")

	sm := repository.NewSmartMeter(&client, &writeAPI)
	ne_rep := repository.NewNewestEvents(&client, &writeAPI)

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			instantaneous, updateAt, err := remo_api.GetInstantaneous()
			if err != nil {
				log.Println(err)
			} else {
				sm.PostSmartMeter(instantaneous, updateAt)
			}

			ne, err := remo_api.GetNewestEvents()
			if err != nil {
				log.Println(err)
			} else {
				ne_rep.PostNewestEvents(ne)
			}
		}
	}

	defer client.Close()
}
