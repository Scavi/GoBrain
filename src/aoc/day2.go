package aoc

func solveDay2A(input []int) int {
	run := true
	currentPos := 0
	result := 0
	for run {
		operation := input[currentPos]
		lookupA := input[currentPos+1]
		lookupB := input[currentPos+2]
		switch operation {
		case 1:
			result = input[lookupA] + input[lookupB]
		case 2:
			result = input[lookupA] * input[lookupB]
		case 99:
			run = false
		default:
			panic("Oh no!")
		}
		input[currentPos+3%len(input)] = result
		currentPos += 4
	}
	return result
}

func solveDay2B(input []int) int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			input[1] = i
			input[2] = j
			result := solveDay2A(input)
			if result == 19690720 {
				return 100*input[1] + input[2]
			}
		}
	}
	panic("whaaaat?")
}
