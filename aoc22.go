package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TODO: Figure out modules so each day can be put into its own file

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
	fmt.Println("* * * * * * * * * * Day 01 * * * * * * * * * *")
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
	fmt.Println()
}

// Each round string will be formatted as "A Y"
// A == X == rock, 1 pt
// B == Y == paper, 2 pt
// C == Z == scissors, 3 pt
func RoundScorePart1(round string) int {
	values := strings.Split(round, " ")
	opponent := values[0]
	you := values[1]

	score := 0
	switch you {
	case "X":
		score += 1
		switch opponent {
		case "A": // draw
			score += 3
		case "C": // win
			score += 6
		}
	case "Y":
		score += 2
		switch opponent {
		case "B": // draw
			score += 3
		case "A": // win
			score += 6
		}
	case "Z":
		score += 3
		switch opponent {
		case "C": // draw
			score += 3
		case "B": // win
			score += 6
		}
	}
	return score
}

// Each round string will be formatted as "A Y"
// A == opponent rock, 1 pt
// B == opponent paper, 2 pt
// C == opponent scissors, 3 pt
// X == you should lose
// Y == you should draw
// Z == you should win
func RoundScorePart2(round string) int {
	values := strings.Split(round, " ")
	opponent := values[0]
	outcome := values[1]

	score := 0
	switch outcome {
	case "X": // lose
		score += 0
		// Determine how many points you get for what you throw
		if opponent == "A" {
			score += 3
		} else if opponent == "B" {
			score += 1
		} else if opponent == "C" {
			score += 2
		}
	case "Y": // draw
		score += 3
		// Determine how many points you get for what you throw
		if opponent == "A" {
			score += 1
		} else if opponent == "B" {
			score += 2
		} else if opponent == "C" {
			score += 3
		}
	case "Z": // win
		score += 6
		// Determine how many points you get for what you throw
		if opponent == "A" {
			score += 2
		} else if opponent == "B" {
			score += 3
		} else if opponent == "C" {
			score += 1
		}
	}
	return score
}

func Day2() {
	fmt.Println("* * * * * * * * * * Day 02 * * * * * * * * * *")
	dataPath := "inputs/day02"
	reader, err := Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalScore := 0
	for line := range reader {
		totalScore += RoundScorePart1(line)
	}

	fmt.Println("Part 1: Your total score would be", totalScore)

	reader, err = Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalScore = 0
	for line := range reader {
		totalScore += RoundScorePart2(line)
	}

	fmt.Println("Part 2: Your total score would be", totalScore)
	fmt.Println()
}

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
	reader, err := Readlines(dataPath)
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

func main() {
	// Day1()
	// Day2()
	Day3()
}
