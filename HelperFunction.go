package RadioChannelModel

import "math"

// helper function
func calculateDistance3D(src Position, dst Position) float64 {
	dx := dst.X - src.X
	dy := dst.Y - src.Y
	dz := dst.Z - src.Z
	distance3D := math.Sqrt(dx*dx + dy*dy + dz*dz)
	return distance3D
}

func calculate_distance_2D(txPosition []float64, rxPosition []float64) float64 {
	dx := txPosition[0] - rxPosition[0]
	dy := txPosition[1] - rxPosition[1]
	distance2D := math.Sqrt(dx*dx + dy*dy)
	return distance2D
}

func calculateElevation(txPosition Position, rxPosition Position) float64 {
	dx := txPosition.X - rxPosition.X
	dy := txPosition.Y - rxPosition.Y
	dz := math.Abs(txPosition.Z - rxPosition.Z)
	dxy := math.Sqrt(dx*dx + dy*dy)
	elevation := math.Atan(dz / dxy)
	return elevation
}
