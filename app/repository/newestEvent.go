package repository

import (
	"github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"github.com/mizuho1998/grafana_monitor/app/model"
	"time"
)

type NewestEvents struct {
	client   *influxdb2.Client
	writeAPI *api.WriteAPI
}

func NewNewestEvents(client *influxdb2.Client, wapi *api.WriteAPI) *NewestEvents {
	return &NewestEvents{client: client, writeAPI: wapi}
}

func (sm *NewestEvents) PostNewestEvents(ne *model.NewestEvents) error {
	p := influxdb2.NewPoint("humidity",
		map[string]string{"unit": "humidity"},
		map[string]interface{}{"hu": ne.Hu.Val, "createdAt": ne.Hu.CreatedAt},
		time.Now())
	(*sm.writeAPI).WritePoint(p)
	(*sm.writeAPI).Flush()
	p = influxdb2.NewPoint("illuminance",
		map[string]string{"unit": "illuminance"},
		map[string]interface{}{"il": ne.Il.Val, "createdAt": ne.Il.CreatedAt},
		time.Now())
	(*sm.writeAPI).WritePoint(p)
	(*sm.writeAPI).Flush()

	p = influxdb2.NewPoint("motion",
		map[string]string{"unit": "motion"},
		map[string]interface{}{"mo": ne.Mo.Val, "createdAt": ne.Mo.CreatedAt},
		time.Now())
	(*sm.writeAPI).WritePoint(p)
	(*sm.writeAPI).Flush()

	p = influxdb2.NewPoint("temperature",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"te": ne.Te.Val, "createdAt": ne.Te.CreatedAt},
		time.Now())
	(*sm.writeAPI).WritePoint(p)
	(*sm.writeAPI).Flush()

	return nil
}
