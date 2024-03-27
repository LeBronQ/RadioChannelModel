package RadioChannelModel

import "math"

type BPSKParam struct {
	Bandwidth float64
	SNR       float64
	BitRate   float64
}

type QAMParam struct {
	Bandwidth float64
	SNR       float64
	BitRate   float64
	M         float64
}

func CalculateBPSKBER(p BPSKParam) float64 {
	BandwidthInHz, SNR, BitRate := p.Bandwidth, p.SNR, p.BitRate
	SNR = math.Pow(10, SNR/10.0)
	EbN0 := SNR * BandwidthInHz / BitRate
	BER := 0.5 * math.Erfc(math.Sqrt(EbN0))
	return BER
}

func CalculateQAMBER(p QAMParam) float64 {
	BandwidthInHz, SNR, BitRate, M := p.Bandwidth, p.SNR, p.BitRate, p.M
	EbN0 := SNR * BandwidthInHz / BitRate
	m := math.Log2(M)
	x := math.Sqrt(1.5 * m * EbN0 / (M - 1))
	x1 := (1 - 1/math.Sqrt(M)) * math.Erfc(x)
	BER := 2 / m * x1
	return BER
}
