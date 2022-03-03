package utils

import "github.com/chars-mc/opengl-exercises/domain"

const pointsPerQuadrant = 100

func GetPointsFromCoordinates(coordinates domain.Coordinates) []float32 {
	var points []float32

	for _, coordinate := range coordinates {
		point := []float32{
			float32(coordinate.X) / pointsPerQuadrant, // X axis
			float32(coordinate.Y) / pointsPerQuadrant, // Y axis
			0.0, // Z axis
		}
		points = append(points, point...)
	}

	return points
}
