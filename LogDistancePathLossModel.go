package RadioChannelModel

import (
	"github.com/gonum/stat/distuv"
	"math"
)

type LogDistanceParam struct {
	Distance     float64
	Frequency    float64
	Scenario     string
	TXHeight     float64
	Foliage      int64
	TXPowerInDbm float64
}

func LogDistancePathLoss(p LogDistanceParam) float64 {
	distance, frequency, scenario, txHeight, foliage, txPowerInDbm := p.Distance, p.Frequency, p.Scenario, p.TXHeight, p.Foliage, p.TXPowerInDbm
	wavelength := c / frequency
	n := 2.0
	sigma := 0.
	pl0 := 10.0 * 2.0 * math.Log10((4*π*d0)/wavelength)
	if scenario == "open_field" {
		n = 2.0
		sigma = 3.0
		if frequency >= 3.1e+9 && frequency <= 5.3e+9 {
			if foliage == 0 {
				if txHeight < 1.0 {
					n = 2.9442
					sigma = 2.799
				} else {
					n = 2.5418
					sigma = 3.06
				}
			} else {
				n = 2.6471
				sigma = 3.06
			}
		}
	} else if scenario == "urban" {
		n = 3.0
		sigma = 3.0
		if frequency >= 9.0e+8 && frequency <= 1.2e+9 {
			n = 1.7
			sigma = 2.6
		} else if frequency >= 5.03e+9 && frequency <= 5.091e+9 {
			n = 2.0
			sigma = 3.2
		}
	} else if scenario == "shadowed_urban" {
		n = 4.0
		sigma = 5.0
	}
	shadowing := distuv.Normal{Mu: 0, Sigma: sigma * sigma}
	pl := pl0 + 10*n*math.Log10(distance/d0) + shadowing.Rand()
	return txPowerInDbm - pl
}
