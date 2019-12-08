package aoc

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"utils"
)

func TestDay6I(t *testing.T) {
	var puzzleInput = []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	result := solveDay6A(puzzleInput)
	assert.Equal(t, 42, result)
}

func TestDay6II(t *testing.T) {
	var puzzleInput = []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	result := solveDay6B(puzzleInput)
	assert.Equal(t, 4, result)
}

func TestDay6A(t *testing.T) {
	puzzleInput, err := utils.PreparePuzzleInput("day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := solveDay6A(puzzleInput)
	assert.Equal(t, 227612, result)
}


func TestDay6B(t *testing.T) {
	puzzleInput, err := utils.PreparePuzzleInput("day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := solveDay6B(puzzleInput)
	assert.Equal(t, 454, result)
}