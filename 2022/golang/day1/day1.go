package day1

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var highestCalories = []int{0, 0, 0}

func Run() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	currentCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			promoteCalories(currentCalories)
			currentCalories = 0
		} else {
			calories, _ := strconv.Atoi(line)
			currentCalories += calories
		}
	}
	promoteCalories(currentCalories)

	fmt.Println("highest calories:", highestCalories)
	sum := highestCalories[0] + highestCalories[1] + highestCalories[2]
	fmt.Println("summed highest calories:", sum)
}

func promoteCalories(calories int) {
	temp := highestCalories
	temp = append(temp, calories)
	sort.Sort(sort.Reverse(sort.IntSlice(temp)))
	highestCalories = temp[:3]
}
