package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "car_stats_viewer"
	SubSystem = "car"
)

var (
	Speed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "speed",
			Help:      "Speed of the Car",
		}, []string{})

	StateOfCharge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "state_of_charge",
			Help:      "Battery state of charge",
		}, []string{"charging", "parked"})

	Power = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "power",
			Help:      "Power usage",
		}, []string{"charging", "parked"})
)

func init() {
	prometheus.MustRegister(
		Speed,
		StateOfCharge,
		Power,
	)
}
