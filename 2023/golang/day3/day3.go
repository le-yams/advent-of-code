package day3

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var testInput string

func RunPart1() {
	fmt.Println(part1(input))
}
