package aoc

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
	"utils"
)

func TestDay2(t *testing.T) {
	var input = []string{"1", "9", "10", "3", "2", "3", "11", "0", "99", "30", "40", "50"}
	puzzleInput, _ := utils.ConvertStringToInt(input)
	result := solveDay2A(puzzleInput)
	assert.Equal(t, 3500, result)
}


func TestDay2A(t *testing.T) {
	fileData, err := utils.PreparePuzzleInput("day2.txt")
	if err != nil || len(fileData) != 1 {
		log.Fatal(err)
	}
	tmpInput := strings.Split(fileData[0], ",")
	puzzleInput, _ := utils.ConvertStringToInt(tmpInput)
	puzzleInput[1] = 12
	puzzleInput[2] = 2
	result := solveDay2A(puzzleInput)
	assert.Equal(t, 3224742, result)
}


func TestDay2B(t *testing.T) {
	fileData, err := utils.PreparePuzzleInput("day2.txt")
	if err != nil || len(fileData) != 1 {
		log.Fatal(err)
	}
	tmpInput := strings.Split(fileData[0], ",")
	puzzleInput, _ := utils.ConvertStringToInt(tmpInput)
	result := solveDay2B(puzzleInput)
	assert.Equal(t, 5190822, result)
}
