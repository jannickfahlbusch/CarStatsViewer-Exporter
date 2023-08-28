package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "car_stats_viewer"
	SubSystem = "car"
)

var (
	SubmittedDataPoints = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "submitted_data_points_total",
			Help:      "Total amount of submitted data points",
		})

	AmbientTemperature = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "ambient_temperature",
			Help:      "Ambient Temperature",
		})

	BatteryLevel = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "battery_level",
			Help:      "Battery level",
		})

	Power = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "power",
			Help:      "Power usage",
		}, []string{"charging"})

	Speed = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "speed",
			Help:      "Speed",
		})

	StateOfCharge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "state_of_charge",
			Help:      "Battery state of charge",
		})
)

func init() {
	prometheus.MustRegister(
		SubmittedDataPoints,
		AmbientTemperature,
		BatteryLevel,
		Power,
		Speed,
		StateOfCharge,
	)
}
