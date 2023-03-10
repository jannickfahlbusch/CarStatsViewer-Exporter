package types

type LiveData struct {
	StateOfCharge float64 `json:"soc"`
	UTC           int     `json:"utc"`
	Power         float64 `json:"power"`
	IsCharging    bool    `json:"is_charging"`
	IsParked      bool    `json:"is_parked"`
	Speed         float64 `json:"speed"`
}
