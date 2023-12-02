package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rexsimiloluwah/adventofcode23/utils"
)

// Game configuration
var redCubes = 12
var greenCubes = 13
var blueCubes = 14

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "1 or 2")
	flag.Parse()

	var ans int
	var err error

	if part == 1 {
		ans = part1()
	} else {
		ans = part2()
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("Answer: %v", ans)
}

func part1() int {
	games, err := utils.ReadInputFile("./input.txt")

	if err != nil {
		panic(err)
	}

	sumID := 0

	for _, game := range games {
		parsedGame, id, _ := parseGame(game)
		possible := processInputPart1(parsedGame)

		if possible {
			sumID += id
		}
	}

	return sumID
}

func part2() int {
	games, err := utils.ReadInputFile("./input.txt")

	if err != nil {
		panic(err)
	}

	sumPower := 0

	for _, game := range games {
		parsedGame, _, _ := parseGame(game)
		setPower := processInputPart2(parsedGame)
		sumPower += setPower
	}

	return sumPower
}

// Process a game for part 1
func processInputPart1(input string) bool {
	// Split the game into sets
	sets := strings.Split(input, ";")

	// Check if each set is possible
	possible := true

	for _, set := range sets {
		r, g, b := parseSet(set)

		if b > blueCubes || g > greenCubes || r > redCubes {
			possible = false
			break
		}
	}

	return possible
}

// Process a game for part 2
func processInputPart2(input string) int {
	// Split the game into sets
	sets := strings.Split(input, ";")

	// Fewest number of cubes required for each color
	rSets, gSets, bSets := []int{}, []int{}, []int{}

	for _, set := range sets {
		r, g, b := parseSet(set)

		rSets = append(rSets, r)
		gSets = append(gSets, g)
		bSets = append(bSets, b)
	}

	rMin, _ := utils.Max(rSets)
	gMin, _ := utils.Max(gSets)
	bMin, _ := utils.Max(bSets)

	return rMin * gMin * bMin
}

// Parse the game to extract the game ID and the sets
func parseGame(input string) (string, int, error) {
	splitGame := strings.Split(input, ":")
	game := strings.Trim(splitGame[1], " ")
	gameID, _ := strconv.Atoi(strings.Trim(strings.Split(splitGame[0], "Game")[1], " "))
	return game, gameID, nil
}

// Parse a game set to extract the number of red, green, and blue cubes respectively
func parseSet(set string) (int, int, int) {
	pattern := `(\d+) (blue|red|green)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(set, -1)

	blueCount := 0
	redCount := 0
	greenCount := 0

	for _, match := range matches {
		count, _ := strconv.Atoi(match[1])
		color := match[2]
		switch color {
		case "blue":
			blueCount += count
		case "red":
			redCount += count
		case "green":
			greenCount += count
		}
	}

	return redCount, greenCount, blueCount
}
