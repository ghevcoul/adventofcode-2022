package day05

import (
	"errors"
	"fmt"
	"ghevcoul/aoc22/utils"
	"strconv"
	"strings"
)

/*
 * Implement a stack data structure for use in this challenge
 * If it becomes necessary for other challenges, can move it into utils
 */
type Stack struct {
	top int  // The index of the element at the top of the stack
	stackArray []string  // An array holding the stack elements
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.top = -1
	stack.stackArray = make([]string, 100)
	return stack
}

func (stack *Stack) Empty() bool {
	return stack.top == -1
}

func (stack *Stack) Size() int {
	return stack.top + 1
}

func (stack *Stack) Push(data string) {
	stack.top++
	stack.stackArray[stack.top] = data
}

func (stack *Stack) Pop() (string, error) {
	if stack.top < 0 {
		return "", errors.New("Stack is empty!")
	}
	value := stack.stackArray[stack.top]
	stack.top--
	return value, nil
}

func (stack *Stack) Peek() string {
	return stack.stackArray[stack.top]
}

/*
 * End of Stack implementation
 */

type Operation struct {
	from, to, numCrates int
}

func ParseInput(path string) ([]Stack, []Operation) {
	reader, err := utils.Readlines(path)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	// Determine the starting configuration of the crates
	// and populate the respective stacks with that configuration 
	lineBuffer := NewStack()
	for line := range reader {
		if len(line) > 0 {
			lineBuffer.Push(line)
		} else {
			break
		}
	}
	line, _ := lineBuffer.Pop()
	numStacks := len(strings.Split(line, "   "))
	
	stackList := make([]Stack, numStacks)
	for i := 0; i < numStacks; i++ {
		stackList[i] = *NewStack()
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
func PerformOperationOneAtATime(operation Operation, stacks []Stack) {
	for i := 0; i < operation.numCrates; i++ {
		inHand, _ := stacks[operation.from].Pop()
		stacks[operation.to].Push(inHand)
	}
}

// Perform operations assuming all requested crates can be moved at once
func PerformOperationInBulk(operation Operation, stacks []Stack) {
	internalStack := NewStack()
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
