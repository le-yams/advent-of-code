package day3

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

func part1(in string) string {
	scanner := bufio.NewScanner(strings.NewReader(in))
	row := 0
	var parts []*part
	var symbols []*position

	fmt.Println("### Parsing engine map ###")
	for scanner.Scan() {
		line := scanner.Text()
		rowParts, rowSymbols := parseEngineRow(line, row)
		parts = append(parts, rowParts...)
		symbols = append(symbols, rowSymbols...)
		row += 1
	}

	fmt.Println("### Looking for parts near a symbol ###")
	parts = partsNearSymbols(parts, symbols)
	value := 0
	for _, enginePart := range parts {
		value += enginePart.intValue()
	}
	return fmt.Sprintf("%d", value)
}

func parseEngineRow(line string, row int) ([]*part, []*position) {
	var parts []*part
	var symbols []*position

	var currentPart []rune

	for column, char := range line {
		if unicode.IsNumber(char) {
			currentPart = append(currentPart, char)
			continue
		}
		if len(currentPart) > 0 {
			value := string(currentPart)
			newPart := &part{
				value:       value,
				row:         row,
				startColumn: column - len(value),
				endColumn:   column,
			}
			fmt.Println("part found", newPart)
			parts = append(parts, newPart)
		}
		if char != '.' {
			pos := &position{
				row:    row,
				column: column,
			}
			fmt.Println("symbol found at", pos)
			symbols = append(symbols, pos)
		}
		currentPart = []rune{}
	}

	return parts, symbols
}

func partsNearSymbols(parts []*part, symbols []*position) []*part {
	var ps []*part
	for _, p := range parts {
		near := p.isNear(symbols)
		if near {
			fmt.Println(p, "is near a symbol")
			ps = append(ps, p)
		}
	}
	return ps
}
