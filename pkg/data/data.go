package data

import (
	"encoding/json"
	"fmt"
	"github.com/jannickfahlbusch/CarStatsViewer-Exporter/pkg/metrics"
	"github.com/jannickfahlbusch/CarStatsViewer-Exporter/pkg/types"
	"log"
	"net/http"
)

type LiveDataHandler struct{}

func (handler *LiveDataHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	decoder := json.NewDecoder(request.Body)
	var data *types.LiveData
	err := decoder.Decode(data)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	rw.WriteHeader(http.StatusOK)

	metrics.Speed.WithLabelValues().Set(data.Speed)
	metrics.StateOfCharge.WithLabelValues(fmt.Sprintf("%t", data.IsCharging), fmt.Sprintf("%t", data.IsParked)).Set(data.StateOfCharge)
	metrics.StateOfCharge.WithLabelValues(fmt.Sprintf("%t", data.IsCharging), fmt.Sprintf("%t", data.IsParked)).Set(data.Power)
}
