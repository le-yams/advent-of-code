package day3

import (
	"fmt"
	"strconv"
)

type part struct {
	value       string
	row         int
	startColumn int
	endColumn   int
}

func (p *part) intValue() int {
	intValue, _ := strconv.Atoi(p.value)
	return intValue
}

func (p *part) String() string {
	return fmt.Sprintf("part %s at [%d,%d-%d]", p.value, p.row, p.startColumn, p.endColumn)
}

func (p *part) isNear(symbols []*position) bool {
	var nearSymbols []*position

	// before part
	nearSymbols = append(nearSymbols, &position{
		row:    p.row - 1,
		column: p.startColumn - 1,
	})

	nearSymbols = append(nearSymbols, &position{
		row:    p.row,
		column: p.startColumn - 1,
	})

	nearSymbols = append(nearSymbols, &position{
		row:    p.row + 1,
		column: p.startColumn - 1,
	})

	// part
	for i := p.startColumn; i <= p.endColumn-1; i++ {
		nearSymbols = append(nearSymbols, &position{
			row:    p.row - 1,
			column: i,
		})
		nearSymbols = append(nearSymbols, &position{
			row:    p.row + 1,
			column: i,
		})
	}

	// after part
	nearSymbols = append(nearSymbols, &position{
		row:    p.row - 1,
		column: p.endColumn,
	})
	nearSymbols = append(nearSymbols, &position{
		row:    p.row,
		column: p.endColumn,
	})
	nearSymbols = append(nearSymbols, &position{
		row:    p.row + 1,
		column: p.endColumn,
	})

	return containsOneOf(symbols, nearSymbols)
}
