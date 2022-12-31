package day09

import (
	"fmt"
	"ghevcoul/aoc22/utils"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	// Right positive x, Left negative x
	x int
	// Up positive y, Down negative y
	y int
}

func (position *Position) toString() string {
	return fmt.Sprintf("%d, %d", position.x, position.y)
}

func updateTail(head *Position, tail *Position) {
	xDiff := head.x - tail.x
	yDiff := head.y - tail.y
	if math.Abs(float64(xDiff)) <= 1 && math.Abs(float64(yDiff)) <= 1 {
		// return if head and tail are touching
		return
	} else if xDiff > 0 && yDiff == 0 {
		// Move tail RIGHT toward head
		tail.x++
	} else if xDiff < 0 && yDiff == 0 {
		// move tail LEFT toward head
		tail.x--
	} else if xDiff == 0 && yDiff > 0 {
		// move tail UP toward head
		tail.y++
	} else if xDiff == 0 && yDiff < 0 {
		// move tail DOWN toward head
		tail.y--
	} else {
		// Diagonally more than 2 units away
		// Move tail x and y toward head
		if xDiff > 0 {
			tail.x++
		} else {
			tail.x--
		}
		if yDiff > 0 {
			tail.y++
		} else {
			tail.y--
		}
	}
}

func processLine(input string, knots []*Position, visited map[string]struct{}) {
	splitInput := strings.Split(input, " ")
	direction := splitInput[0]
	count, _ := strconv.Atoi(splitInput[1])

	head := knots[0]	
	for i := 0; i < count; i++ {
		switch direction {
		case "R":
			head.x++
		case "L":
			head.x--
		case "U":
			head.y++
		case "D":
			head.y--
		}
		for j := 1; j < len(knots); j++ {
			updateTail(knots[j-1], knots[j])
		}
		visited[knots[len(knots)-1].toString()] = struct{}{}
	}
}

func runSimulation(path string, numKnots int) {
	reader, err := utils.Readlines(path)
	if err != nil {
		fmt.Println(err)
	}

	knots := make([]*Position, numKnots)
	for i := 0; i < numKnots; i++ {
		knot := new(Position)
		knot.x = 0
		knot.y = 0
		knots[i] = knot
	}

	visited := make(map[string]struct{})
	visited[knots[numKnots-1].toString()] = struct{}{}

	for line := range reader {
		processLine(line, knots, visited)
	}

	fmt.Println("After simulating the rope with", numKnots, "knots, the tail visited", len(visited), "total positions")
}

func Day09() {
	fmt.Println("* * * * * * * * * * Day 09 * * * * * * * * * *")
	dataPath := "inputs/day09"
	runSimulation(dataPath, 2)
	runSimulation(dataPath, 10)
	fmt.Println()
}