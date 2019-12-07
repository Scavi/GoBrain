package aoc

func solveDay2A(input []int) int {
	return executeInstructions(input, 0)
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
