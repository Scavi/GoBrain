package utils

import (
	"bufio"
	"log"
	"os"
	"path"
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
		log.Fatal(err)
	}
	inputPath := path.Join(path.Dir(dir), "../res/aoc/"+testFileName)
	puzzleInput, err := ReadFile(inputPath)
	return puzzleInput, err
}
