package types

type LiveData struct {
	AmbientTemperature   float64 `json:"ambientTemperature"`
	AvgConsumption       float64 `json:"avgConsumption"`
	AvgSpeed             float64 `json:"avgSpeed"`
	BatteryLevel         float64 `json:"batteryLevel"`
	ChargePortConnected  bool    `json:"chargePortConnected"`
	ChargeStartDate      string  `json:"chargeStartDate"`
	ChargeTime           int     `json:"chargeTime"`
	ChargedEnergy        float64 `json:"chargedEnergy"`
	CurrentGear          int     `json:"currentGear"`
	CurrentIgnitionState int     `json:"currentIgnitionState"`
	CurrentPower         float64 `json:"currentPower"`
	CurrentSpeed         float64 `json:"currentSpeed"`
	DriveState           int     `json:"driveState"`
	InstConsumption      float64 `json:"instConsumption"`
	IsCharging           bool    `json:"isCharging"`
	IsFastCharging       bool    `json:"isFastCharging"`
	IsParked             bool    `json:"isParked"`
	MaxBatteryLevel      float64 `json:"maxBatteryLevel"`
	StateOfCharge        int     `json:"stateOfCharge"`
	TravelTime           int     `json:"travelTime"`
	TraveledDistance     float64 `json:"traveledDistance"`
	TripStartDate        string  `json:"tripStartDate"`
	UsedEnergy           float64 `json:"usedEnergy"`
}
