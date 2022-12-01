package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1a")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentCount := 0
	maxCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			val, _ := strconv.Atoi(line)
			currentCount += val
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
			}
			currentCount = 0
		}
	}

	fmt.Println("The most calories carried by an elf is ", maxCount)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
