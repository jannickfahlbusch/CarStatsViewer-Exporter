package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/pkg/errors"
	types2 "pkg.jf-projects.de/owntracks/types"

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
	Owntracks *Owntracks
}

func (handler *LiveDataHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	decoder := json.NewDecoder(request.Body)
	var data *types.LiveData
	err := decoder.Decode(&data)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	metrics.SubmittedDataPoints.Inc()

	metrics.AmbientTemperature.Set(data.AmbientTemperature)
	metrics.AverageConsumption.Set(data.AvgConsumption)
	metrics.AverageSpeed.Set(data.AvgSpeed)
	metrics.BatteryLevel.Set(data.BatteryLevel)
	metrics.ChargedEnergy.Set(data.ChargedEnergy)
	metrics.Gear.Set(float64(data.CurrentGear))
	metrics.IgnitionState.Set(float64(data.CurrentIgnitionState))
	metrics.Power.WithLabelValues(fmt.Sprintf("%t", data.IsCharging), fmt.Sprintf("%t", data.IsFastCharging), fmt.Sprintf("%t", data.IsParked)).Set(data.CurrentPower)
	metrics.Speed.Set(data.CurrentSpeed)
	metrics.DriveState.Set(float64(data.DriveState))
	metrics.StateOfCharge.Set(float64(data.StateOfCharge))
	metrics.InstantConsumption.Set(data.InstConsumption)
	metrics.MaxBatteryLevel.Set(data.MaxBatteryLevel)
	metrics.TraveledDistance.Set(data.TraveledDistance)
	metrics.UsedEnergy.Set(data.UsedEnergy)

	if data.HasCoordinates() && handler.Owntracks.Enabled() {
		err = handler.Owntracks.Publish(request.Context(), data)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (ot *Owntracks) Enabled() bool {
	return ot.Client != nil
}

func (ot *Owntracks) Publish(ctx context.Context, data *types.LiveData) error {
	currentTime := time.Now()
	payload := &types2.Location{
		Type:      types2.LocationType,
		EpochTime: currentTime.Unix(),
		Timestamp: &currentTime,
		Altitude:  data.Altitude,
		Battery: types2.Battery{
			BatteryLevel: data.StateOfCharge,
		},
		WiFi:              types2.WiFi{},
		Latitude:          data.Latitude,
		Longitude:         data.Longitude,
		Trigger:           types2.TriggerPing,
		TrackerID:         "ps",
		Velocity:          data.CurrentSpeed,
		DistanceTravelled: int(math.Round(data.TraveledDistance)),
	}

	if data.IsCharging || data.IsFastCharging {
		payload.Battery.BatteryStatus = types2.BatteryStatusCharging
	}

	err := ot.Client.Publish(ctx, ot.User, ot.Device, payload)
	if err != nil {
		err = errors.Wrap(err, "publishing liveData to owntracks failed")
	}

	return err
}
