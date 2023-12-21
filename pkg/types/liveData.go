package types

import (
	"encoding/json"
	"time"
)

type UnixTime struct {
	time.Time
}

type DrivingPoint struct {
	Timestamp       UnixTime `json:"driving_point_epoch_time"`
	EnergyDelta     float64  `json:"energy_delta"`
	DistaceDelta    float64  `json:"distance_delta"`
	PointMarkerType int      `json:"point_marker_type"`
	StateOfCharge   float64  `json:"state_of_charge"`
	Latitude        float64  `json:"lat,omitempty"`
	Longitude       float64  `json:"lon,omitempty"`
	Altitude        float64  `json:"alt,omitempty"`
}

type ChargingSession struct {
	StartDate          UnixTime `json:"start_epoch_time"`
	EndDate            UnixTime `json:"end_epoc_time"`
	ChargedEnergy      float64  `json:"charged_energy"`
	ChargedSOC         float64  `json:"charged_soc"`
	OutsideTemperature float64  `json:"outside_temp"`

	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}

type LiveData struct {
	APIVersion          json.Number `json:"apiVersion"`
	AppVersion          string      `json:"appVersion"`
	Timestamp           UnixTime    `json:"timestamp"`
	Speed               float64     `json:"speed"`
	Power               float64     `json:"power"`
	SelectedGear        string      `json:"selectedGear"`
	IgnitionState       string      `json:"ignitionState"`
	ChargePortConnected bool        `json:"chargePortConnected"`
	BatteryLevel        float64     `json:"batteryLevel"`
	StateOfCharge       float64     `json:"stateOfCharge"`
	AmbientTemperature  float64     `json:"ambientTemperature"`

	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
	Altitude  float64 `json:"alt,omitempty"`

	DrivingPoints    []*DrivingPoint    `json:"drivingPoints,omitempty"`
	ChargingSessions []*ChargingSession `json:"chargingSessions,omitempty"`
}

func (liveData *LiveData) HasCoordinates() bool {
	return liveData.Latitude != 0 && liveData.Longitude != 0
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var epochTime int64
	err := json.Unmarshal(b, &epochTime)
	if err != nil {
		return err
	}

	u.Time = time.Unix(epochTime/1000, 0)

	return nil
}
