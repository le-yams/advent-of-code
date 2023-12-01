package day1

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

type digitWord struct {
	value rune
	label string
}

var digitWords = []*digitWord{
	{'1', "one"},
	{'2', "two"},
	{'3', "three"},
	{'4', "four"},
	{'5', "five"},
	{'6', "six"},
	{'7', "seven"},
	{'8', "eight"},
	{'9', "nine"},
}

func Run() {
	calibration := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		lineCalibration, _ := calibrationValue(line)
		fmt.Println("line:", line, "\ncalibration:", lineCalibration)
		calibration += lineCalibration
	}
	fmt.Println("total calibration:", calibration)
}

func calibrationValue(line string) (int, error) {
	calibrationStr := string([]rune{firstNumber(line), lastNumber(line)})
	return strconv.Atoi(calibrationStr)
}

func firstNumber(s string) rune {
	chars := []rune(s)
	for index := 0; index < len(chars); index++ {
		char := chars[index]
		if unicode.IsNumber(char) {
			return char
		}
		digit := parseDigitWord(s, index)
		if digit != nil {
			return digit.value
		}
	}
	return '0'
}

func lastNumber(s string) rune {
	chars := []rune(s)
	for index := len(chars) - 1; index >= 0; index-- {
		char := chars[index]
		if unicode.IsNumber(char) {
			return char
		}
		digit := parseDigitWord(s, index)
		if digit != nil {
			return digit.value
		}
	}
	return '0'
}

func parseDigitWord(text string, index int) *digitWord {
	for _, digit := range digitWords {
		regex, _ := regexp.Compile(fmt.Sprintf("^.{%d}%s", index, digit.label))
		if regex.MatchString(text) {
			return digit
		}
	}
	return nil
}
