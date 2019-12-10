package utils

import (
	"bufio"
	"os"
	"path"
	"strconv"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func PreparePuzzleInput(testFileName string) ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	inputPath := path.Join(path.Dir(dir), "../res/aoc/"+testFileName)
	puzzleInput, err := ReadFile(inputPath)
	return puzzleInput, err
}

func ConvertStringToInt(input []string) ([]int, error) {
	var output []int
	for _, current := range input {
		converted, err := strconv.Atoi(current)
		if err != nil {
			return nil, err
		}
		output = append(output, converted)
	}
	return output, nil
}

func CountOccurrenceInString(input string, lookup string) int {
	if len(input) < len(lookup) {
		return 0
	}
	count := 0
	for i := 0; i < len(input)-len(lookup); i++ {
		if input[i:i+len(lookup)] == lookup {
			count += 1
		}
	}
	return count
}
