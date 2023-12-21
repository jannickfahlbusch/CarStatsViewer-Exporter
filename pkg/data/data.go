package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"pkg.jf-projects.de/carstatsviewer-exporter/pkg/data/cache"

	"github.com/pkg/errors"
	owntracksTypes "pkg.jf-projects.de/owntracks/types"

	"pkg.jf-projects.de/carstatsviewer-exporter/pkg/metrics"
	"pkg.jf-projects.de/carstatsviewer-exporter/pkg/types"
	"pkg.jf-projects.de/owntracks/client"
)

type Owntracks struct {
	Client client.Client
	User   string
	Device string
}

type LiveDataHandler struct {
	Owntracks    *Owntracks
	LogTelegrams bool
}

func (handler *LiveDataHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	payload, err := io.ReadAll(request.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}
	var data *types.LiveData

	err = json.Unmarshal(payload, &data)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	if handler.LogTelegrams {
		if err != nil {
			log.Printf("Failed to encode payload: %v\n", err)
		}
		log.Printf("Received payload: %v\n", string(payload))
	}

	metrics.SubmittedDataPoints.Inc()

	cache.SetLastPayload(data)

	metrics.Speed.Set(data.Speed)
	metrics.Power.WithLabelValues(fmt.Sprintf("%t", data.ChargePortConnected)).Set(data.Power)
	metrics.BatteryLevel.Set(data.BatteryLevel)
	metrics.StateOfCharge.Set(data.StateOfCharge)
	metrics.AmbientTemperature.Set(data.AmbientTemperature)

	if data.HasCoordinates() && handler.OwntracksEnabled() {
		err = handler.Owntracks.Publish(request.Context(), data)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (handler *LiveDataHandler) OwntracksEnabled() bool {
	return handler.Owntracks != nil && handler.Owntracks.Enabled()
}

func (ot *Owntracks) Enabled() bool {
	return ot.Client != nil
}

func (ot *Owntracks) Publish(ctx context.Context, data *types.LiveData) error {
	soc := int(data.StateOfCharge * 100)

	payload := &owntracksTypes.Location{
		Type:      owntracksTypes.LocationType,
		EpochTime: data.Timestamp.Unix(),
		Timestamp: &data.Timestamp.Time,
		Altitude:  data.Altitude,
		Battery: owntracksTypes.Battery{
			BatteryLevel: soc,
		},
		WiFi:      owntracksTypes.WiFi{},
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Trigger:   owntracksTypes.TriggerPing,
		TrackerID: "ps",
		Velocity:  data.Speed * 3.6, // Speed is in m/s, we want km/h here
	}

	if data.ChargePortConnected {
		payload.Battery.BatteryStatus = owntracksTypes.BatteryStatusCharging
	}

	err := ot.Client.Publish(ctx, ot.User, ot.Device, payload)
	if err != nil {
		err = errors.Wrap(err, "publishing liveData to owntracks failed")
	}

	return err
}
