package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "car_stats_viewer"
	SubSystem = "car"
)

var (
	Speed = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "speed",
			Help:      "Speed of the Car",
		})

	StateOfCharge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "state_of_charge",
			Help:      "Battery state of charge",
		})

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
