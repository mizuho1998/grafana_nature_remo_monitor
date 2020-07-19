package repository

import (
	"github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"time"
)

type SmartMeter struct {
	client   *influxdb2.Client
	writeAPI *api.WriteAPI
}

func NewSmartMeter(client *influxdb2.Client, wapi *api.WriteAPI) *SmartMeter {
	return &SmartMeter{client: client, writeAPI: wapi}
}

func (sm *SmartMeter) PostSmartMeter(val int, updateAt time.Time) error {
	p := influxdb2.NewPoint("instantaneousW",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"val": val, "nature_update_at": updateAt},
		time.Now())

	// write asynchronously
	(*sm.writeAPI).WritePoint(p)
	// Force all unwritten data to be sent
	(*sm.writeAPI).Flush()

	return nil
}
