package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func Day1() {
	dataPath := "inputs/day01"
	reader, err := Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	counts := make([]int, 0, 1000)
	currentCount := 0
	for line := range reader {
		if len(line) > 0 {
			val, _ := strconv.Atoi(line)
			currentCount += val
		} else {
			counts = append(counts, currentCount)
			currentCount = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	fmt.Println("The most calories carried by an elf is ", counts[0])
	topThree := counts[0] + counts[1] + counts[2]
	fmt.Println("The total calories carried by the top three elves is", topThree)
}

func main() {
	Day1()
}
