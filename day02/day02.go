package day02

import (
	"fmt"
	"strings"

	"ghevcoul/aoc22/utils"
)

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
	reader, err := utils.Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalScore := 0
	for line := range reader {
		totalScore += RoundScorePart1(line)
	}

	fmt.Println("Part 1: Your total score would be", totalScore)

	reader, err = utils.Readlines(dataPath)
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
