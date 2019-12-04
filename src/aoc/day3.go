package aoc

import "C"
import (
	"math"
	"strconv"
	"strings"
	"utils"
)

func createWireLines(wire string) []utils.WirePart {
	var output []utils.WirePart
	wirePaths := strings.Split(wire, ",")
	currentX := 0
	currentY := 0
	for _, wirePath := range wirePaths {
		targetX := currentX
		targetY := currentY
		direction := wirePath[0:1]
		length, _ := strconv.Atoi(wirePath[1:])
		switch direction {
		case "U":
			targetY = currentY + length
		case "D":
			targetY = currentY - length
		case "R":
			targetX = currentX + length
		case "L":
			targetX = currentX - length
		default:
			panic("noohoo!")
		}
		wirePart := utils.WirePart{P1: utils.Point{X: currentX, Y: currentY}, P2: utils.Point{X: targetX, Y: targetY},}
		output = append(output, wirePart)
		currentX = wirePart.P2.X
		currentY = wirePart.P2.Y
	}
	return output
}

func ccw(point1 utils.Point, point2 utils.Point, point3 utils.Point) bool {
	return (point3.Y-point1.Y)*(point2.X-point1.X) > (point2.Y-point1.Y)*(point3.X-point1.X)
}

// https://bryceboe.com/2006/10/23/line-segment-intersection-algorithm/
func hasIntersection(wirePart1 utils.WirePart, wirePart2 utils.WirePart) bool {
	return ccw(wirePart1.P1, wirePart2.P1, wirePart2.P2) != ccw(wirePart1.P2, wirePart2.P1, wirePart2.P2) &&
		ccw(wirePart1.P1, wirePart1.P2, wirePart2.P1) != ccw(wirePart1.P1, wirePart1.P2, wirePart2.P2)
}

func findIntersection(wirePart1 utils.WirePart, wirePart2 utils.WirePart) utils.Point {
	a1 := wirePart1.P2.Y - wirePart1.P1.Y
	b1 := wirePart1.P1.X - wirePart1.P2.X
	c1 := a1*wirePart1.P1.X + b1*wirePart1.P1.Y

	a2 := wirePart2.P2.Y - wirePart2.P1.Y
	b2 := wirePart2.P1.X - wirePart2.P2.X
	c2 := a2*wirePart2.P1.X + b2*wirePart2.P1.Y

	delta := a1*b2 - a2*b1
	return utils.Point{X: int((b2*c1 - b1*c2) / delta), Y: int((a1*c2 - a2*c1) / delta)}
}

func solveDay3(wireInput1 string, wireInput2 string) (int, int) {
	start := utils.Point{X: 0, Y: 0}
	wire1 := createWireLines(wireInput1)
	wire2 := createWireLines(wireInput2)

	minDistance := math.MaxInt32
	minSteps := math.MaxInt32
	wireDistance1 := 0
	for _, wirePart1 := range wire1 {
		wireDistance2 := 0
		for _, wirePart2 := range wire2 {
			if hasIntersection(wirePart1, wirePart2) {
				intersectionPoint := findIntersection(wirePart1, wirePart2)
				minDistance = utils.Min(minDistance, utils.ManhattanDistance(start, intersectionPoint))
				missing := utils.PointDistance(wirePart1.P1, intersectionPoint) +
					utils.PointDistance(wirePart2.P1, intersectionPoint)
				minSteps = utils.Min(minSteps, wireDistance2+wireDistance1+missing)
			}
			wireDistance2 += utils.WireDistance(wirePart2)
		}
		wireDistance1 += utils.WireDistance(wirePart1)
	}
	return minSteps, minDistance
}
