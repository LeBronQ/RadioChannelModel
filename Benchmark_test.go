package RadioChannelModel

import "testing"

func PLRCalculation(txPosition []float64, rxPosition []float64) float64 {
	frequency := 5.8e+9
	Bandwidth := 2.0e+7
	BitRate := 2.0e+6
	distance := calculate_distance_3D(txPosition, rxPosition)
	elevation := calculate_elevation(txPosition, rxPosition)
	pathLoss := free_space_path_loss(distance, frequency)
	fading := NakagamiFadingModel(pathLoss, "open_filed", elevation)
	SNR := calculate_SNR(Bandwidth, fading, 0)
	BER := calculate_BPSK_BER(Bandwidth, SNR, BitRate)
	PLR := calculate_PLR(BER, 10000, 1)
	return PLR
}

func Benchmark(b *testing.B) {
	txPosition := []float64{1000, 1000, 1000}
	rxPosition := []float64{0.0, 0.0, 0.0}
	for i := 0; i < b.N; i++ {
		PLRCalculation(txPosition, rxPosition)
	}
}
