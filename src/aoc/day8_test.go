package aoc

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"utils"
)

func TestDay8I(t *testing.T) {
	var puzzleInput = "123456789012"
	result := solveDay8A(puzzleInput, 3, 2)
	assert.Equal(t, 2, result)
}


func TestDay8II(t *testing.T) {
	var puzzleInput = "0222112222120000"
	solveDay8B(puzzleInput, 2, 2)
}

func TestDay8A(t *testing.T) {
	puzzleInput, err := utils.PreparePuzzleInput("day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := solveDay8A(puzzleInput[0], 25, 6)
	assert.Equal(t, 2684, result)  // too high
}

func TestDay8B(t *testing.T) {
	puzzleInput, err := utils.PreparePuzzleInput("day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	solveDay8B(puzzleInput[0], 25, 6)
}

