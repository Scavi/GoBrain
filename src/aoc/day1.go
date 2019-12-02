package aoc

import (
	"strconv"
)

func solve1(inputLines []string) int32 {
	requiredFuel := int32(0)
	for _, line := range inputLines {
		number, _ := strconv.ParseInt(line, 10, 32)
		tmp := int32((number / 3) - 2)
		requiredFuel = requiredFuel + tmp
	}
	return requiredFuel
}

func _solve(fuelValue int32) int32 {
	if fuelValue <= 5 {
		return 0
	}
	requiredFuel := (fuelValue / 3) - 2
	requiredFuel = requiredFuel + _solve(requiredFuel)
	return requiredFuel
}

func solve2(inputLines []string) int32 {
	requiredFuel := int32(0)
	for _, line := range inputLines {
		number, _ := strconv.ParseInt(line, 10, 32)
		requiredFuel = requiredFuel + _solve(int32(number))
	}
	return requiredFuel
}
