package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"pkg.jf-projects.de/carstatsviewer-exporter/pkg/metrics"
	"pkg.jf-projects.de/carstatsviewer-exporter/pkg/types"
)

type LiveDataHandler struct{}

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
}
