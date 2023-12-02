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

func Run() {
	games := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		minimalBlue, minimalRed, minimalGreen := computeGameMinimalCubes(line)
		games += minimalBlue * minimalRed * minimalGreen
	}

	fmt.Println("games:", games)
}

func computeGameMinimalCubes(line string) (int, int, int) {
	_, subsets, _ := strings.Cut(line, ":")

	minimalBlue := 0
	minimalRed := 0
	minimalGreen := 0
	for _, subset := range strings.Split(subsets, ";") {
		blue, red, green := parseSubsetCubes(subset)
		minimalBlue = max(blue, minimalBlue)
		minimalRed = max(red, minimalRed)
		minimalGreen = max(green, minimalGreen)
	}
	return minimalBlue, minimalRed, minimalGreen
}

func parseSubsetCubes(subset string) (int, int, int) {
	blue := 0
	red := 0
	green := 0
	for _, cubes := range strings.Split(subset, ",") {
		cubes = strings.TrimSpace(cubes)
		countStr, color, _ := strings.Cut(cubes, " ")
		count, _ := strconv.Atoi(countStr)
		switch color {
		case "blue":
			blue = count
		case "red":
			red = count
		case "green":
			green = count
		}
	}
	return blue, red, green
}
