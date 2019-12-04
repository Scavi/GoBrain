package aoc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay4TestA(t *testing.T) {
	result := solveDay4(147981, 691423, false)
	assert.Equal(t, 1790, result)
}

func TestDay4TestB(t *testing.T) {
	result := solveDay4(147981, 691423, true)
	assert.Equal(t, 1206, result)
}
