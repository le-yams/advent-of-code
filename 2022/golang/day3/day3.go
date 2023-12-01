package day3

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	priorities := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		middleIndex := len(rucksack) / 2
		compartment1 := rucksack[:middleIndex]
		compartment2 := rucksack[middleIndex:]

		sharedItems := findSharedItems(compartment1, compartment2)

		priorities += itemsPriorities(sharedItems)
	}
	fmt.Println("priorities:", priorities)
}

func findSharedItems(compartment1 string, compartment2 string) []rune {

	var sharedItems []rune
	for _, item := range distinctItems(compartment1) {
		if strings.ContainsRune(compartment2, item) {
			sharedItems = append(sharedItems, item)
		}
	}
	return sharedItems
}

func itemsPriorities(sharedItems []rune) int {
	priorities := 0
	for _, item := range sharedItems {
		if item >= 'a' && item <= 'z' {
			priorities += int(item) - int('a') + 1
		} else {
			priorities += int(item) - int('A') + 27
		}
	}
	return priorities
}

func distinctItems(compartment string) []rune {
	keys := make(map[rune]bool)
	var distinct []rune
	for _, item := range compartment {
		if _, value := keys[item]; !value {
			keys[item] = true
			distinct = append(distinct, item)
		}
	}
	return distinct
}
