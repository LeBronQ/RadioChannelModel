package RadioChannelModel

import "math"

type HataOkumuraParam struct {
	Distance     float64
	Frequency    float64
	City         string
	Environment  string
	BaseHeight   float64
	MobileHeight float64
	TXPowerInDbm float64
}

func HataOkumuraModel(p HataOkumuraParam) float64 {
	frequency, city, environment, baseHeight, mobileHeight, distance, txPowerInDbm := p.Frequency, p.City, p.Environment, p.BaseHeight, p.MobileHeight, p.Distance, p.TXPowerInDbm
	distance /= 1000
	fMhz := frequency / 1e+6
	logfMhz := math.Log10(fMhz)
	Ch, pl := 0.0, 0.0
	if frequency <= 1.5e+9 {
		if city == "LargeCity" {
			if fMhz < 200 {
				Ch = 8.29*math.Pow(math.Log10(1.54*mobileHeight), 2) - 1.1
			} else {
				Ch = 3.2*math.Pow(math.Log10(11.75*mobileHeight), 2) - 4.97
			}
		} else {
			Ch = 0.8 + (1.1*logfMhz-0.7)*mobileHeight - 1.56*logfMhz
		}
		pl = 69.55 + 26.16*logfMhz - 13.82*math.Log10(baseHeight) - Ch + (44.9-6.55*math.Log10(baseHeight))*math.Log10(distance)
		if environment == "Suburban" {
			pl = pl - 2*math.Pow(math.Log10(fMhz/28), 2) - 5.4
		} else if environment == "OpenArea" {
			pl = pl - 4.78*math.Pow(logfMhz, 2) + 18.33*logfMhz - 40.94
		}
	} else {
		C, F := 0.0, 0.0
		if city == "LargeCity" {
			F = 3.2 * math.Pow(math.Log10(11.75*mobileHeight), 2)
			C = 3.0
		} else {
			F = (1.1*logfMhz-0.7)*mobileHeight - (1.56*logfMhz - 0.8)
		}
		pl = 46.3 + (33.9 * logfMhz) - 13.82*math.Log10(baseHeight) +
			((44.9 - (6.55 * math.Log10(baseHeight))) * math.Log10(distance)) - F + C
	}
	return txPowerInDbm - pl
}
