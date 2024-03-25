package RadioChannelModel

import "math"

func CalculateSNR(BandwidthInHz float64, rxPowerInDbm float64, EnvironmentNoiseInmW float64) float64 {
	ThermalNoiseInDbm := -174 + 10*math.Log10(BandwidthInHz)
	ThermalNoiseInmW := math.Pow(10, ThermalNoiseInDbm/10)
	rxPowerInmW := math.Pow(10, rxPowerInDbm/10)
	SNRInDb := math.Log10(rxPowerInmW / (ThermalNoiseInmW + EnvironmentNoiseInmW))
	return SNRInDb
}
