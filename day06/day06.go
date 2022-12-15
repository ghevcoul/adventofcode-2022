package day06

import (
	"fmt"
	"strings"

	"ghevcoul/aoc22/utils"
)

func Day6() {
	fmt.Println("* * * * * * * * * * Day 06 * * * * * * * * * *")
	dataPath := "inputs/day06"
	reader, err := utils.Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// There should be only one line in this input file
	var line string
	for loopLine := range reader {
		line = loopLine
		break
	}

	// Loop through the line, starting with char 3 until finding a group of 4 unique characters
	splitLine := strings.Split(line, "")
	var marker int
	for i := 4; i <= len(splitLine); i++ {
		set := make(map[string]struct{})
		for _, char := range splitLine[i-4 : i] {
			set[char] = struct{}{}
		}
		if len(set) == 4 {
			marker = i
			break
		}
	}
	fmt.Println("The first start-of-packet marker is at character", marker)

	// Loop through the line again, starting with char 13 until finding a group of 14 unique characters
	for i := 14; i <= len(splitLine); i++ {
		set := make(map[string]struct{})
		for _, char := range splitLine[i-14 : i] {
			set[char] = struct{}{}
		}
		if len(set) == 14 {
			marker = i
			break
		}
	}
	fmt.Println("The first start-of-packet marker is at character", marker)
	fmt.Println()
}
