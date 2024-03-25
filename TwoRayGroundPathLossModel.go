package RadioChannelModel

import "math"

type TwoRayGroundParam struct {
	Distance     float64
	Frequency    float64
	TXHeight     float64
	RXHeight     float64
	TXPowerInDbm float64
}

func TwoRayGroundPathLoss(p TwoRayGroundParam) float64 {
	distance, frequency, txHeight, rxHeight, txPowerInDbm := p.Distance, p.Frequency, p.TXHeight, p.RXHeight, p.TXPowerInDbm
	wavelength := c / frequency
	pl := 0.0
	d := 4 * π * txHeight * rxHeight / wavelength
	//distance <= d for Two_ray path; using Friis
	if distance <= d {
		pl = 20.0 * math.Log10((4*π*distance)/wavelength)
	} else {
		pl = 10.0 * math.Log10(txHeight*txHeight*rxHeight*rxHeight/math.Pow(distance, 4))
	}
	return txPowerInDbm - pl
}
