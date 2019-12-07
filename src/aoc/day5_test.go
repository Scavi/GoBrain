package aoc

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
	"utils"
)

func TestDay5A(t *testing.T) {
	fileData, err := utils.PreparePuzzleInput("day5.txt")
	if err != nil || len(fileData) != 1 {
		log.Fatal(err)
	}
	tmpInput := strings.Split(fileData[0], ",")
	puzzleInput, _ := utils.ConvertStringToInt(tmpInput)
	result := solveDay5A(puzzleInput)
	assert.Equal(t, 13346482, result)
}

func TestDay5B(t *testing.T) {
	fileData, err := utils.PreparePuzzleInput("day5.txt")
	if err != nil || len(fileData) != 1 {
		log.Fatal(err)
	}
	tmpInput := strings.Split(fileData[0], ",")
	puzzleInput, _ := utils.ConvertStringToInt(tmpInput)
	result := solveDay5B(puzzleInput)
	assert.Equal(t, 12111395, result)
}
