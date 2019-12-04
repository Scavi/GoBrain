package aoc

import (
	"regexp"
	"strconv"
)

func generate(from, to int) []int {
	var output []int
	for i := from; i < to; i++ {
		output = append(output, i)
	}
	return output
}

func solveDay4(from, to int, replaceIllegal bool) int {
	var DigitsIncreasing = regexp.MustCompile("^(0*1*2*3*4*5*6*7*8*9*)$")
	var TwoAdjacentDigits = regexp.MustCompile("(11|22|33|44|55|66|77|88|99)")
	var Illegal = regexp.MustCompile("(1{3,6}|2{3,6}|3{3,6}|4{3,6}|5{3,6}|6{3,6}|7{3,6}|8{3,6}|9{3,6})")
	passwordCount := 0
	possiblePasswords := generate(from, to)
	for _, possiblePassword := range possiblePasswords {
		pw := strconv.Itoa(possiblePassword)
		if !DigitsIncreasing.MatchString(pw) {
			continue
		}
		if replaceIllegal {
			pw = Illegal.ReplaceAllString(pw, "")
		}
		if !TwoAdjacentDigits.MatchString(pw) {
			continue
		}
		passwordCount += 1
	}
	return passwordCount
}
