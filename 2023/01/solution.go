package main

import (
	"flag"
	"fmt"
	"unicode"

	"github.com/rexsimiloluwah/adventofcode23/utils"
)

var digitWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var digitWordMap = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "1 or 2")
	flag.Parse()

	var ans int
	var err error

	if part == 1 {
		ans, err = part1()
	} else {
		ans, err = part2()
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("Answer: %v", ans)
}

// Solve part 1
func part1() (int, error) {
	inputs, err := utils.ReadInputFile("./input.txt")

	if err != nil {
		return 0, err
	}

	finalAnswer := 0
	for _, input := range inputs {
		finalAnswer += processInputPart1(input)
	}

	return finalAnswer, nil
}

// Solve part 2
func part2() (int, error) {
	inputs, err := utils.ReadInputFile("./input.txt")

	if err != nil {
		return 0, err
	}

	finalAnswer := 0
	for _, input := range inputs {
		finalAnswer += processInputPart2(input)
	}

	return finalAnswer, nil
}

// Process each input for part 1 by combining the first and last digit
func processInputPart1(input string) int {
	var firstDigit, lastDigit rune

	for _, c := range input {
		if unicode.IsDigit(c) {
			if firstDigit == 0 {
				firstDigit = c
			}
			lastDigit = c
		}
	}

	result := int(firstDigit-'0')*10 + int(lastDigit-'0')
	return result
}

// Process each input for part 2 by combining the first and last digit
// Here, the digits spelled out with letters are parsed as valid digits
func processInputPart2(input string) int {
	var firstDigit, lastDigit rune

	for i, c := range input {
		if unicode.IsDigit(c) {
			if firstDigit == 0 {
				firstDigit = c
			}
			lastDigit = c
		}

		for _, w := range digitWords {
			if i < len(input)-len(w)+1 && input[i:i+len(w)] == w {
				if firstDigit == 0 {
					firstDigit = digitWordMap[w]
				}
				lastDigit = digitWordMap[w]
			}
		}
	}

	result := int(firstDigit-'0')*10 + int(lastDigit-'0')

	return result
}
