package RadioChannelModel

import (
	"github.com/gonum/stat/distuv"
	"math"
)

const (
	c  = 2.998e+8
	π  = math.Pi
	d0 = 1.0
)

// path loss model
func free_space_path_loss(distance float64, frequency float64) float64 {
	wavelength := c / frequency
	fspl := 20.0 * math.Log10((4*π*distance)/wavelength)
	return fspl
}

func two_ray_ground_path_loss(distance float64, frequency float64, txHeight float64, rxHeight float64, txPowerInDbm float64) float64 {
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

func log_distance_pl(distance float64, frequency float64, scenario string, txHeight float64, foliage int64, txPowerInDbm float64) float64 {
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

func HataOkumuraModel(frequency float64, city string, environment string, baseHeight float64, mobileHeight float64, distance float64, txPowerInDbm float64) float64 {
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

// NakagamiFadingModel -- small-scale fading model
// Nakagami-m fading model. m = 1 -- Rayleigh fading model
func NakagamiFadingModel(txPowerInDbm float64, scenario string, elevation float64) float64 {
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

func calculate_SNR(BandwidthInHz float64, rxPowerInDbm float64, EnvironmentNoiseInmW float64) float64 {
	ThermalNoiseInDbm := -174 + 10*math.Log10(BandwidthInHz)
	ThermalNoiseInmW := math.Pow(10, ThermalNoiseInDbm/10)
	rxPowerInmW := math.Pow(10, rxPowerInDbm/10)
	SNRInDb := math.Log10(rxPowerInmW / (ThermalNoiseInmW + EnvironmentNoiseInmW))
	return SNRInDb
}

func calculate_BPSK_BER(BandwidthInHz float64, SNR float64, BitRate float64) float64 {
	EbN0 := SNR * BandwidthInHz / BitRate
	BER := 0.5 * math.Erfc(math.Sqrt(EbN0))
	return BER
}

func calculate_QAM_BER(BandwidthInHz float64, SNR float64, BitRate float64, M float64) float64 {
	EbN0 := SNR * BandwidthInHz / BitRate
	m := math.Log2(M)
	x := math.Sqrt(1.5 * m * EbN0 / (M - 1))
	x1 := (1 - 1/math.Sqrt(M)) * math.Erfc(x)
	BER := 2 / m * x1
	return BER
}

func calculate_PLR(BER float64, PacketSizeInBit float64, redundancy float64) float64 {
	//todo: coding scheme
	PER := 1 - math.Pow(1-BER, PacketSizeInBit)
	PLR := math.Pow(PER, redundancy)
	return PLR
}
