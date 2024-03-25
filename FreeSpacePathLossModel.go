package RadioChannelModel

import "math"

type FreeSpaceParam struct {
	Distance     float64
	Frequency    float64
	TXPowerInDbm float64
}

func FreeSpacePathLoss(p FreeSpaceParam) float64 {
	wavelength := c / p.Frequency
	fspl := 20.0 * math.Log10((4*π*p.Distance)/wavelength)
	return p.TXPowerInDbm - fspl
}
