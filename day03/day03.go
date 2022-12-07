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

func PopulateSet(sack string) map[rune]struct{} {
	set := make(map[rune]struct{})
	for _, item := range sack {
		set[item] = struct{}{}
	}
	return set
}

func ItemToPriority(item rune) int {
	var priority int

	if item > 97 {
		priority = int(item) - 96
	} else {
		priority = int(item) - 38
	}

	return priority
}

func Day3a(dataPath string) {
	reader, err := utils.Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	prioritySum := 0
	for line := range reader {
		midpoint := len(line) / 2

		// Populate sets of the items in each rucksack
		sack1Set := PopulateSet(line[:midpoint])
		sack2Set := PopulateSet(line[midpoint:])
		// Find the common item between the sets
		var commonItem rune
		for item := range SetIntersection(sack1Set, sack2Set) {
			commonItem = item
		}
		prioritySum += ItemToPriority(commonItem)
	}

	fmt.Println("The total priority of the common items is ", prioritySum)
}

func Day3b(dataPath string) {
	reader, err := utils.Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	prioritySum := 0
	currentGroups := make([]string, 0, 3)
	for line := range reader {
		currentGroups = append(currentGroups, line)
		if len(currentGroups) == 3 {
			elf1Set := PopulateSet(currentGroups[0])
			elf2Set := PopulateSet(currentGroups[1])
			elf3Set := PopulateSet(currentGroups[2])
			commonItems := SetIntersection(SetIntersection(elf1Set, elf2Set), elf3Set)

			var commonItem rune
			for item := range commonItems {
				commonItem = item
			}
			prioritySum += ItemToPriority(commonItem)
			// Clear the current groups list
			currentGroups = make([]string, 0, 3)
		}
	}

	fmt.Println("The total priority of the elves' badges is ", prioritySum)
}

func Day3() {
	fmt.Println("* * * * * * * * * * Day 03 * * * * * * * * * *")
	dataPath := "inputs/day03"
	Day3a(dataPath)
	Day3b(dataPath)
	fmt.Println()
}
