package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jannickfahlbusch/CarStatsViewer-Exporter/pkg/metrics"
	"github.com/jannickfahlbusch/CarStatsViewer-Exporter/pkg/types"
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

	metrics.Speed.Set(data.Speed)
	metrics.StateOfCharge.Set(data.StateOfCharge)
	metrics.Power.WithLabelValues(fmt.Sprintf("%t", data.IsCharging), fmt.Sprintf("%t", data.IsParked)).Set(data.Power)
}
