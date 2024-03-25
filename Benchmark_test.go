package RadioChannelModel

import "testing"

func PLRCalculation(txPosition Position, rxPosition Position) float64 {
	frequency := 5.8e+9
	Bandwidth := 2.0e+7
	BitRate := 2.0e+6
	distance := calculateDistance3D(txPosition, rxPosition)
	elevation := calculateElevation(txPosition, rxPosition)
	PLParam := FreeSpaceParam{
		Distance:  distance,
		Frequency: frequency,
	}
	pathLoss := FreeSpacePathLoss(PLParam)
	FParam := NakagamiParam{
		TXPowerInDbm: pathLoss,
		Scenario:     "open_filed",
		Elevation:    elevation,
	}
	fading := NakagamiFadingModel(FParam)
	SNR := CalculateSNR(Bandwidth, fading, 0)
	BERP := BPSKParam{
		Bandwidth, SNR, BitRate,
	}
	BER := CalculateBPSKBER(BERP)
	TP := TransportParam{
		BER, 10000, 1,
	}
	PLR := CalculatePLR(TP)
	return PLR
}

func Benchmark(b *testing.B) {
	txPosition := Position{1000, 1000, 1000}
	rxPosition := Position{0.0, 0.0, 0.0}
	for i := 0; i < 1000000; i++ {
		PLRCalculation(txPosition, rxPosition)
	}
}
