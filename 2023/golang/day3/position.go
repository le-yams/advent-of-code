package day3

import (
	"fmt"
	"slices"
)

type position struct {
	row    int
	column int
}

func (p *position) String() string {
	return fmt.Sprintf("position [%d,%d]", p.row, p.column)
}

func containsOneOf(symbols []*position, other []*position) bool {
	fmt.Println("\ncomparing", symbols, "\nwith", other)
	for _, symbol := range symbols {
		for _, o := range other {
			if symbol.row == o.row && symbol.column == o.column {
				fmt.Println(symbol, " found in", other, " contains:", slices.Contains(other, symbol))
				return true
			}
		}
		//if slices.Contains(other, symbol) {
		//	return true
		//}

	}
	return false
}
