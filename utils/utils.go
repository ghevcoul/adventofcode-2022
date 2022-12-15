package utils

import (
	"bufio"
	"errors"
	"os"
)

// An iterator that returns one line of a file at a time
// Based on/stolen from https://bbengfort.github.io/2016/12/yielding-functions-for-iteration-golang/
func Readlines(path string) (<-chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl, nil
}

/*
 * Implement a stack data structure
 * Currently implemented to hold Strings
 */
type Stack struct {
	top        int      // The index of the element at the top of the stack
	stackArray []string // An array holding the stack elements
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
		return "", errors.New("Stack is empty")
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
