package day4

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
	scanner := bufio.NewScanner(strings.NewReader(input))

	overlaps := 0
	for scanner.Scan() {
		s1, s2, _ := strings.Cut(scanner.Text(), ",")
		sections1 := parseSections(s1)
		sections2 := parseSections(s2)

		if sections1.fullyContains(sections2) || sections2.fullyContains(sections1) {
			overlaps += 1
		}
	}

	fmt.Println("overlaps:", overlaps)
}

func parseSections(s string) *sections {
	start, end, _ := strings.Cut(s, "-")
	startValue, _ := strconv.Atoi(start)
	endValue, _ := strconv.Atoi(end)
	return &sections{
		start: startValue,
		end:   endValue,
	}
}

type sections struct {
	start int
	end   int
}

func (sections *sections) String() string {
	return fmt.Sprintf("%d-%d", sections.start, sections.end)
}
func (sections *sections) fullyContains(other *sections) bool {
	return sections.start <= other.start && sections.end >= other.end
}
