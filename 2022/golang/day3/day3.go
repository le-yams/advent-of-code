package day3

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	priorities := 0
	for scanner.Scan() {
		rucksack1 := scanner.Text()
		scanner.Scan()
		rucksack2 := scanner.Text()
		scanner.Scan()
		rucksack3 := scanner.Text()

		badge, _ := findSharedItem(distinctItems(rucksack1), distinctItems(rucksack2), distinctItems(rucksack3))

		priorities += itemsPriorities([]rune{badge})
	}
	fmt.Println("priorities:", priorities)
}

func findSharedItem(items1 []rune, items2 []rune, items3 []rune) (rune, error) {
	for _, item := range items1 {
		if strings.ContainsRune(string(items2), item) &&
			strings.ContainsRune(string(items3), item) {
			return item, nil
		}
	}
	return -1, errors.New("no shared item found")
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
