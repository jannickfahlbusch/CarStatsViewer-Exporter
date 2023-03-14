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

	metrics.Speed.Set(data.CurrentSpeed)
	metrics.StateOfCharge.Set(float64(data.StateOfCharge))
	metrics.Power.WithLabelValues(fmt.Sprintf("%t", data.IsCharging), fmt.Sprintf("%t", data.IsFastCharging), fmt.Sprintf("%t", data.IsParked)).Set(data.CurrentPower)
}
