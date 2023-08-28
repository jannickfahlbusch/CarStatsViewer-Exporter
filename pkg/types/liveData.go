package types

type DrivingPoints struct {
	Timestamp       int64   `json:"driving_point_epoch_time"`
	EnergyDelta     float64 `json:"energy_delta"`
	DistaceDelta    float64 `json:"distance_delta"`
	PointMarkerType int     `json:"point_marker_type"`
	StateOfCharge   float64 `json:"state_of_charge"`
	Latitude        float64 `json:"lat,omitempty"`
	Longitude       float64 `json:"lon,omitempty"`
	Altitude        float64 `json:"alt,omitempty"`
}

type ChargingSessions struct {
	StartDate          int64   `json:"start_epoch_time"`
	EndDate            int64   `json:"end_epoc_time"`
	ChargedEnergy      float64 `json:"charged_energy"`
	ChargedSOC         float64 `json:"charged_soc"`
	OutsideTemperature float64 `json:"outside_temp"`

	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}

type LiveData struct {
	APIVersion          string  `json:"apiVersion"`
	AppVersion          string  `json:"appVersion"`
	Timestamp           int64   `json:"timestamp"`
	Speed               float64 `json:"speed"`
	Power               float64 `json:"power"`
	SelectedGear        string  `json:"selectedGear"`
	IgnitionState       string  `json:"ignitionState"`
	ChargePortConnected bool    `json:"chargePortConnected"`
	BatteryLevel        float64 `json:"batteryLevel"`
	StateOfCharge       float64 `json:"stateOfCharge"`
	AmbientTemperature  float64 `json:"ambientTemperature"`

	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
	Altitude  float64 `json:"alt,omitempty"`

	DrivingPoints    *DrivingPoints    `json:"drivingPoints,omitempty"`
	ChargingSessions *ChargingSessions `json:"chargingSessions,omitempty"`
}

func (liveData *LiveData) HasCoordinates() bool {
	return liveData.Latitude != 0 && liveData.Longitude != 0
}
