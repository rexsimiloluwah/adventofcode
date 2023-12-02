package utils

import (
	"bufio"
	"os"
)

// Read the input file
func ReadInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var inputs []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return inputs, nil
}
