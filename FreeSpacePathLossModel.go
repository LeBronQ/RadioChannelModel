package RadioChannelModel

import (
	"fmt"
	"math"
)

type FreeSpaceParam struct {
	Distance     float64
	Frequency    float64
	TXPowerInDbm float64
}

func FreeSpacePathLoss(p FreeSpaceParam) float64 {
	if p.Distance == 0 {
		fmt.Printf("Distance is 0, no path loss")
		return p.TXPowerInDbm
	} else {
		wavelength := c / p.Frequency
		fspl := 20.0 * math.Log10((4*π*p.Distance)/wavelength)
		return p.TXPowerInDbm - fspl
	}
}
