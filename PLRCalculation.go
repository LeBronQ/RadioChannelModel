package RadioChannelModel

import "math"

type TransportParam struct {
	BER             float64
	PacketSizeInBit float64
	Redundancy      int64
}

func CalculatePLR(p TransportParam) float64 {
	//todo: coding scheme
	BER, PacketSizeInBit, redundancy := p.BER, p.PacketSizeInBit, p.Redundancy
	PER := 1 - math.Pow(1-BER, PacketSizeInBit)
	PLR := math.Pow(PER, float64(redundancy))
	return PLR
}
