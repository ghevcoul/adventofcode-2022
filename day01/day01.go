package day01

import (
	"fmt"
	"sort"
	"strconv"

	"ghevcoul/aoc22/utils"
)

func Day1() {
	fmt.Println("* * * * * * * * * * Day 01 * * * * * * * * * *")
	dataPath := "inputs/day01"
	reader, err := utils.Readlines(dataPath)
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
	fmt.Println()
}
