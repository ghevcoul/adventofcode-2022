package day04

import (
	"fmt"
	"strconv"
	"strings"


	"ghevcoul/aoc22/utils"
)

// Takes two integer ranges and determines whether one fully contains the other
func RangesFullyContained(aStart int, aEnd int, bStart int, bEnd int) bool {
	return (aStart <= bStart && aEnd >= bEnd) || (bStart <= aStart && bEnd >= aEnd)
}

// Takes two integer ranges and determines whether there is any overlap
// Any number in common counts as overlap, so 1-4 and 4-6 overlap
func RangesOverlap(aStart int, aEnd int, bStart int, bEnd int) bool {
	return (aStart <= bStart && bStart <= aEnd) || (aStart <= bEnd && bEnd <= aEnd) || (bStart <= aStart && aStart <= bEnd) || (bStart <= aEnd && aEnd <= bEnd)
}

func Day4() {
	fmt.Println("* * * * * * * * * * Day 04 * * * * * * * * * *")
	dataPath := "inputs/day04"
	reader, err := utils.Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	containedWork := 0
	overlappingWork := 0
	for line := range reader {
		elves := strings.Split(line, ",")
		elf1 := strings.Split(elves[0], "-")
		elf1start, _ := strconv.Atoi(elf1[0])
		elf1end, _ := strconv.Atoi(elf1[1])
		elf2 := strings.Split(elves[1], "-")
		elf2start, _ := strconv.Atoi(elf2[0])
		elf2end, _ := strconv.Atoi(elf2[1])

		if RangesFullyContained(elf1start, elf1end, elf2start, elf2end) {
			containedWork++
		}
		if RangesOverlap(elf1start, elf1end, elf2start, elf2end) {
			overlappingWork++
		}
	}

	fmt.Println("There are ", containedWork, " groups with one range fully contained within the other.")
	fmt.Println("There are ", overlappingWork, " groups with overlapping ranges.")
	fmt.Println()
}
