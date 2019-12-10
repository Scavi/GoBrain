package aoc

import (
	"math"
	"utils"
)

func solveDay8A(inputPuzzle string, wide int, tall int) int {
	minZero := math.MaxInt64
	imageRange := wide * tall
	result := 0
	for i := 0; i < len(inputPuzzle); {
		layer := inputPuzzle[i : i+imageRange]
		zeroCount := utils.CountOccurrenceInString(layer, "0")
		if minZero > zeroCount {
			minZero = zeroCount
			result = utils.CountOccurrenceInString(layer, "1") * utils.CountOccurrenceInString(layer, "2")
		}
		i += imageRange
	}
	return result
}

func solveDay8B(inputPuzzle string, wide int, tall int) {
	image := make([][]string, tall)
	for i := range image {
		image[i] = make([]string, wide)
	}
	imageRange := wide * tall
	for i := 0; i < len(inputPuzzle); {
		layer := inputPuzzle[i : i+imageRange]
		for j := 0; j < len(layer); j++ {
			color := layer[j]
			y := j / wide
			x := j % wide
			if image[y][x] == "" && (color == '0' || color == '1') {
				image[y][x] = string(color)
			}
		}
		i += imageRange
	}
	_printImage(image)
}

func _printImage(image [][]string) {
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[0]); x++ {
			if image[y][x] == "1" {
				print(image[y][x])
			} else {
				print(" ")
			}
		}
		print("\n")
	}
}