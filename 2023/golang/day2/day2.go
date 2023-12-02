package day2

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var maxRed = 12
var maxGreen = 13
var maxBlue = 14

func Run() {
	games := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		games += computeGameValue(line)
	}

	fmt.Println("games:", games)
}

func computeGameValue(line string) int {
	gameInfo, subsets, _ := strings.Cut(line, ":")
	gameId, _ := strings.CutPrefix(gameInfo, "Game ")
	gameValue, _ := strconv.Atoi(gameId)
	for _, subset := range strings.Split(subsets, ";") {
		if subsetNotPossible(subset) {
			return 0
		}
	}
	return gameValue
}

func subsetNotPossible(subset string) bool {
	for _, cubes := range strings.Split(subset, ",") {
		cubes = strings.TrimSpace(cubes)
		countStr, color, _ := strings.Cut(cubes, " ")
		count, _ := strconv.Atoi(countStr)
		switch color {
		case "blue":
			if count > maxBlue {
				return true
			}
		case "red":
			if count > maxRed {
				return true
			}
		case "green":
			if count > maxGreen {
				return true
			}
		}
	}
	return false
}
