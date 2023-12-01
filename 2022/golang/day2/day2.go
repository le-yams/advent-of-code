package day2

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type hand int

const (
	rock     hand = 1
	paper    hand = 2
	scissors hand = 3
)

type result int

const (
	lost result = 0
	draw result = 3
	win  result = 6
)

func Run() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()

		otherHand, _ := parseOtherHand(rune(line[0]))
		expectedResult, _ := parseExpectedResult(rune(line[2]))
		myHand := whatShouldIPlay(otherHand, expectedResult)

		lineScore := int(myHand)
		lineScore += score(otherHand, myHand)
		totalScore += lineScore
		fmt.Println(line, "scores:", lineScore)
	}
	fmt.Println("total score:", totalScore)
}

func parseOtherHand(code rune) (hand, error) {
	switch code {
	case 'A':
		return rock, nil
	case 'B':
		return paper, nil
	case 'C':
		return scissors, nil
	}
	return -1, errors.New("unknown hand")
}

func parseMyHand(code rune) (hand, error) {
	switch code {
	case 'X':
		return rock, nil
	case 'Y':
		return paper, nil
	case 'Z':
		return scissors, nil
	}
	return -1, errors.New("unknown hand")
}

func parseExpectedResult(code rune) (result, error) {
	switch code {
	case 'X':
		return lost, nil
	case 'Y':
		return draw, nil
	case 'Z':
		return win, nil
	}
	return -1, errors.New("unknown result")
}

func whatShouldIPlay(otherHand hand, expectedResult result) hand {
	if expectedResult == draw {
		return otherHand
	}
	if expectedResult == win {
		return winningHand(otherHand)
	}

	return losingHand(otherHand)
}

func winningHand(otherHand hand) hand {
	wHand := (otherHand + 1) % 4
	if wHand == 0 {
		wHand += 1
	}
	return wHand
}

func losingHand(otherHand hand) hand {
	lHand := otherHand - 1
	if lHand == 0 {
		lHand = 3
	}
	return lHand
}

func score(otherHand hand, myHand hand) int {
	if myHand == otherHand {
		return int(draw)
	}

	if myHand == winningHand(otherHand) {
		return int(win)
	}
	return int(lost)
}
