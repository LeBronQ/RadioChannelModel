package RadioChannelModel

import "math"

// helper function
func calculate_distance_3D(txPosition []float64, rxPosition []float64) float64 {
	dx := txPosition[0] - rxPosition[0]
	dy := txPosition[1] - rxPosition[1]
	dz := txPosition[2] - rxPosition[2]
	distance3D := math.Sqrt(dx*dx + dy*dy + dz*dz)
	return distance3D
}

func calculate_distance_2D(txPosition []float64, rxPosition []float64) float64 {
	dx := txPosition[0] - rxPosition[0]
	dy := txPosition[1] - rxPosition[1]
	distance2D := math.Sqrt(dx*dx + dy*dy)
	return distance2D
}

func calculate_elevation(txPosition []float64, rxPosition []float64) float64 {
	dx := txPosition[0] - rxPosition[0]
	dy := txPosition[1] - rxPosition[1]
	dz := math.Abs(txPosition[2] - rxPosition[2])
	dxy := math.Sqrt(dx*dx + dy*dy)
	elevation := math.Atan(dz / dxy)
	return elevation
}
