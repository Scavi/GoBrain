package aoc

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"utils"
)

func TestDay3TestA(t *testing.T) {
	fileData, err := utils.PreparePuzzleInput("day3.txt")
	if err != nil || len(fileData) != 2 {
		log.Fatal(err)
	}
	_, result := solveDay3(fileData[0], fileData[1])
	assert.Equal(t, 721, result)
}

func TestDay3B(t *testing.T) {
	fileData, err := utils.PreparePuzzleInput("day3.txt")
	if err != nil || len(fileData) != 2 {
		log.Fatal(err)
	}
	result, _ := solveDay3(fileData[0], fileData[1])
	assert.Equal(t, 7388, result)
}
