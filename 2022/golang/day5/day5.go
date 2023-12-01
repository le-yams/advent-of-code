package day5

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type craneCommand struct {
	move int
	from int
	to   int
}

var craneCommandRegex, _ = regexp.Compile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")

func parseCraneCommand(line string) *craneCommand {
	submatches := craneCommandRegex.FindStringSubmatch(line)
	move, _ := strconv.Atoi(submatches[1])
	from, _ := strconv.Atoi(submatches[2])
	to, _ := strconv.Atoi(submatches[3])
	return &craneCommand{
		move: move,
		from: from,
		to:   to,
	}
}

var stacksCount = 9

func Run() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	stacks := parseInitialStacks(scanner)
	for scanner.Scan() {
		command := parseCraneCommand(scanner.Text())
		moveCrates(command, stacks)
	}

	topCrates := ""
	for stackId := 0; stackId < stacksCount; stackId++ {
		cratesStack := stacks[stackId]
		topCrates = fmt.Sprintf("%s%c", topCrates, cratesStack.Peek())
	}
	fmt.Println("topCrates:", topCrates)
}

func parseInitialStacks(scanner *bufio.Scanner) map[int]*stack.Stack {
	var stackDef []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		stackDef = append(stackDef, line)
	}
	slices.Reverse(stackDef)
	stacks := map[int]*stack.Stack{}
	for stackId := 0; stackId < stacksCount; stackId++ {
		stacks[stackId] = &stack.Stack{}
	}
	for rowId := 1; rowId < len(stackDef); rowId++ {
		row := stackDef[rowId]
		for stackId := 0; stackId < stacksCount; stackId++ {
			index := (stackId * 4) + 1
			crate := row[index]
			if crate != ' ' {
				stacks[stackId].Push(crate)
			}
		}
	}
	return stacks
}

func moveCrates(command *craneCommand, stacks map[int]*stack.Stack) {
	sourceStack := stacks[command.from-1]
	targetStack := stacks[command.to-1]

	for i := 0; i < command.move; i++ {
		crate := sourceStack.Pop()
		fmt.Printf("moving crate '%c' from %d to %d\n", crate, command.from, command.to)
		targetStack.Push(crate)
	}
}
