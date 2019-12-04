package utils

import "math"

type Point struct {
	X int
	Y int
}

type WirePart struct {
	P1 Point
	P2 Point
}

func WireDistance(wirePart WirePart) int {
	return PointDistance(wirePart.P1, wirePart.P2)
}

func PointDistance(p1 Point, p2 Point) int {
	tmpX := math.Pow(float64(p2.X-p1.X), 2)
	tmpY := math.Pow(float64(p2.Y-p1.Y), 2)
	return int(math.Sqrt(tmpX + tmpY))
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func ManhattanDistance(p1 Point, p2 Point) int {
	v1 := math.Abs(float64(p1.X) - float64(p2.X))
	v2 := math.Abs(float64(p1.Y) - float64(p2.Y))
	return int(v1) + int(v2)
}
