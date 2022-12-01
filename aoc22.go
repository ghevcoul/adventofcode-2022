package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("inputs/day1")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counts := make([]int, 0, 1000)
	currentCount := 0
	for scanner.Scan() {
		line := scanner.Text()
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

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
