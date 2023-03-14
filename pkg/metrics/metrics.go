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

	AverageConsumption = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "average_consumption",
			Help:      "Average Consumption",
		})

	AverageSpeed = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "average_speed",
			Help:      "Average Speed",
		})

	BatteryLevel = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "battery_level",
			Help:      "Battery level",
		})

	ChargedEnergy = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "charged_energy",
			Help:      "Charged energy",
		})

	Gear = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "gear",
			Help:      "Gear",
		})

	IgnitionState = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "ignition_state",
			Help:      "Ignition State",
		})

	Power = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "power",
			Help:      "Power usage",
		}, []string{"charging", "fastcharging", "parked"})

	Speed = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "speed",
			Help:      "Speed",
		})

	DriveState = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "drive_state",
			Help:      "Drive State",
		})

	InstantConsumption = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "instant_consumption",
			Help:      "Instant consumption",
		})

	MaxBatteryLevel = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "max_battery_level",
			Help:      "Max Battery Level",
		})

	StateOfCharge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "state_of_charge",
			Help:      "Battery state of charge",
		})

	TraveledDistance = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "distance_traveled",
			Help:      "Traveled distance",
		})

	UsedEnergy = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: SubSystem,
			Name:      "used_energy",
			Help:      "Used energy",
		})
)

func init() {
	prometheus.MustRegister(
		SubmittedDataPoints,
		AmbientTemperature,
		AverageConsumption,
		AverageSpeed,
		BatteryLevel,
		ChargedEnergy,
		Gear,
		IgnitionState,
		Power,
		Speed,
		StateOfCharge,
		DriveState,
		InstantConsumption,
		MaxBatteryLevel,
		TraveledDistance,
		UsedEnergy,
	)
}
