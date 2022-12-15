package day05

import (
	"fmt"
	"ghevcoul/aoc22/utils"
	"strconv"
	"strings"
)

type Operation struct {
	from, to, numCrates int
}

func ParseInput(path string) ([]utils.Stack, []Operation) {
	reader, err := utils.Readlines(path)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	// Determine the starting configuration of the crates
	// and populate the respective stacks with that configuration
	lineBuffer := utils.NewStack()
	for line := range reader {
		if len(line) > 0 {
			lineBuffer.Push(line)
		} else {
			break
		}
	}
	line, _ := lineBuffer.Pop()
	numStacks := len(strings.Split(line, "   "))

	stackList := make([]utils.Stack, numStacks)
	for i := 0; i < numStacks; i++ {
		stackList[i] = *utils.NewStack()
	}
	for !lineBuffer.Empty() {
		line, _ := lineBuffer.Pop()
		lineChars := strings.Split(line, "")
		stackIdx := 0
		linePos := 1
		for {
			if lineChars[linePos] != " " {
				stackList[stackIdx].Push(lineChars[linePos])
			}
			stackIdx++
			linePos += 4
			if linePos >= len(lineChars) {
				break
			}
		}
	}

	// For the remaining lines, parse into Operation objects
	operations := make([]Operation, 1000)
	for line := range reader {
		values := strings.Split(line, " ")
		numCrates, _ := strconv.Atoi(values[1])
		from, _ := strconv.Atoi(values[3])
		to, _ := strconv.Atoi(values[5])
		// When making Operation, shift from and to by 1 so they're 0-indexed
		operations = append(operations, Operation{from - 1, to - 1, numCrates})
	}

	return stackList, operations
}

// Performs Operation assuming only one crate can be moved at a time
func PerformOperationOneAtATime(operation Operation, stacks []utils.Stack) {
	for i := 0; i < operation.numCrates; i++ {
		inHand, _ := stacks[operation.from].Pop()
		stacks[operation.to].Push(inHand)
	}
}

// Perform operations assuming all requested crates can be moved at once
func PerformOperationInBulk(operation Operation, stacks []utils.Stack) {
	internalStack := utils.NewStack()
	for i := 0; i < operation.numCrates; i++ {
		temp, _ := stacks[operation.from].Pop()
		internalStack.Push(temp)
	}
	for !internalStack.Empty() {
		temp, _ := internalStack.Pop()
		stacks[operation.to].Push(temp)
	}
}

func Day5() {
	fmt.Println("* * * * * * * * * * Day 05 * * * * * * * * * *")
	dataPath := "inputs/day05"
	stacks, operations := ParseInput(dataPath)
	for _, operation := range operations {
		PerformOperationOneAtATime(operation, stacks)
	}
	// Print the top of each stack
	fmt.Print("Part1: After performing all operations the top of the stacks is ")
	for _, stack := range stacks {
		fmt.Print(stack.Peek())
	}
	fmt.Println()

	stacks, operations = ParseInput(dataPath)
	for _, operation := range operations {
		PerformOperationInBulk(operation, stacks)
	}
	// Print the top of each stack
	fmt.Print("Part2: After performing all operations the top of the stacks is ")
	for _, stack := range stacks {
		fmt.Print(stack.Peek())
	}
	fmt.Println()
	fmt.Println()
}
