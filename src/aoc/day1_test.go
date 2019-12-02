package aoc

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"utils"
)

func Test1(t *testing.T) {
	puzzleInput, err := utils.PreparePuzzleInput("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	expected := int32(3423279)
	result := solve1(puzzleInput)
	assert.Equal(t, expected, result)
}

func Test2(t *testing.T) {
	puzzleInput, err := utils.PreparePuzzleInput("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	expected := int32(5132018)
	result := solve2(puzzleInput)
	assert.Equal(t, expected, result)
}
