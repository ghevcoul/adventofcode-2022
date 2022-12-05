package day03

import (
	"fmt"

	"ghevcoul/aoc22/utils"
)

// TODO: If we need sets again, move this to utils and make generic
func SetIntersection(setA map[rune]struct{}, setB map[rune]struct{}) map[rune]struct{} {
	intersection := make(map[rune]struct{})

	// iterate through the shorter set
	if len(setA) > len(setB) {
		setA, setB = setB, setA
	}

	for value := range setA {
		if _, ok := setB[value]; ok {
			intersection[value] = struct{}{}
		}
	}
	return intersection
}

func Day3() {
	fmt.Println("* * * * * * * * * * Day 03 * * * * * * * * * *")
	dataPath := "inputs/day03"
	reader, err := utils.Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	prioritySum := 0
	for line := range reader {
		midpoint := len(line) / 2

		// Populate sets of the items in each rucksack
		sack1Set := make(map[rune]struct{})
		for _, item := range line[:midpoint] {
			sack1Set[item] = struct{}{}
		}
		sack2Set := make(map[rune]struct{})
		for _, item := range line[midpoint:] {
			sack2Set[item] = struct{}{}
		}
		// Find the common item between the sets
		var commonItem rune
		for item := range SetIntersection(sack1Set, sack2Set) {
			commonItem = item
		}
		if commonItem > 97 {
			prioritySum += int(commonItem) - 96
		} else {
			prioritySum += int(commonItem) - 38
		}
	}

	fmt.Println("The total priority of the common items is ", prioritySum)
}
