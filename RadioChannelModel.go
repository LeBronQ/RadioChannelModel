package RadioChannelModel

import (
	"math"
)

const (
	c          = 2.998e+8
	π          = math.Pi
	d0         = 1.0
	PacketSize = 10000
	Redundancy = 1
)

type Position struct {
	X float64
	Y float64
	Z float64
}

type WirelessNode struct {
	Frequency  float64
	BitRate    float64
	Modulation string
	BandWidth  float64
	M          float64
	PowerInDbm float64
}

type ChannelModel struct {
	LargeScaleModel string
	SmallScaleModel string
}

var Scene map[string]ChannelModel

func ChannelParameterCalculation(LinkID int64, txNode WirelessNode, rxNode WirelessNode, txPos Position, rxPos Position) float64 {
	//txScene := "LargeCity"
	//LModel := Scene[txScene].LargeScaleModel
	//SModel := Scene[txScene].SmallScaleModel
	LModel := "FreeSpacePathLossModel"
	SModel := "NakagamiFadingModel"
	Distance := calculateDistance3D(txPos, rxPos)
	Elevation := calculateElevation(txPos, rxPos)
	Mod := rxNode.Modulation
	M := rxNode.M
	BR := rxNode.BitRate
	Frequency := rxNode.Frequency
	BW := rxNode.BandWidth
	/*
		PLR := ChannelCalculation(LinkID, Distance, LModel, SModel, Frequency, BR, Mod, Elevation, BW, M, txNode.PowerInDbm)
		fmt.Printf("%e\n", PLR)*/
	PLR := ChannelCalculation(LinkID, Distance, LModel, SModel, Frequency, BR, Mod, Elevation, BW, M, txNode.PowerInDbm)
	return PLR
}

func ChannelCalculation(LinkID int64, Distance float64, LargeScaleModel string, SmallScaleModel string, Frequency float64, BitRate float64, Mod string, Elevation float64, BW float64, M float64, PowerInDbm float64) float64 {
	PathLoss, Fading, BER := 0.0, 0.0, 0.0
	switch LargeScaleModel {
	case "FreeSpacePathLossModel":
		PLParam := FreeSpaceParam{
			Distance:     Distance,
			Frequency:    Frequency,
			TXPowerInDbm: PowerInDbm,
		}
		PathLoss = FreeSpacePathLoss(PLParam)
		break
	case "LogDistancePathLossModel":
	case "":
	}
	switch SmallScaleModel {
	case "NakagamiFadingModel":
		FParam := NakagamiParam{
			TXPowerInDbm: PathLoss,
			Scenario:     "open_filed",
			Elevation:    Elevation,
		}
		Fading = NakagamiFadingModel(FParam)
	}
	SNR := CalculateSNR(BW, Fading, 0)
	switch Mod {
	case "BPSK":
		BParam := BPSKParam{
			Bandwidth: BW,
			SNR:       SNR,
			BitRate:   BitRate,
		}
		BER = CalculateBPSKBER(BParam)
	case "QAM":
		QParam := QAMParam{
			Bandwidth: BW,
			SNR:       SNR,
			BitRate:   BitRate,
			M:         M,
		}
		BER = CalculateQAMBER(QParam)
	}
	TParam := TransportParam{
		BER:             BER,
		PacketSizeInBit: PacketSize,
		Redundancy:      Redundancy,
	}
	PLR := CalculatePLR(TParam)
	return PLR
}
