package RadioChannelModel

import (
	"github.com/gonum/stat/distuv"
	"math"
)

type NakagamiParam struct {
	TXPowerInDbm float64
	Scenario     string
	Elevation    float64
}

// NakagamiFadingModel -- small-scale fading model
// Nakagami-m fading model. m = 1 -- Rayleigh fading model
func NakagamiFadingModel(p NakagamiParam) float64 {
	txPowerInDbm, scenario, elevation := p.TXPowerInDbm, p.Scenario, p.Elevation
	a, b := 0.5, 6.0
	if scenario == "open_field" {
		b = 15.0
	}
	m := a * math.Pow(math.E, b*elevation)
	PowerInWatt := math.Pow(10, (txPowerInDbm-30)/10)
	UnitGamma := distuv.Gamma{Alpha: m, Beta: PowerInWatt / m}
	rxPowerInDbm := 10*math.Log10(UnitGamma.Rand()) + 30.0
	return rxPowerInDbm
}
