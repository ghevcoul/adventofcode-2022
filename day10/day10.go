package day10

import (
	"fmt"
	"ghevcoul/aoc22/utils"
	"strconv"
	"strings"
)

func Contains(values []int, target int) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func Day10() {
	fmt.Println("* * * * * * * * * * Day 10 * * * * * * * * * *")
	dataPath := "inputs/day10"

	reader, err := utils.Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
	}

	cyclesToCheck := []int{20, 60, 100, 140, 180, 220}

	cycle := 1
	xRegister := 1
	temp := 0
	processingAddx := false
	cumSigStrength := 0
	pixelPos := 0
	for {
		if Contains(cyclesToCheck, cycle) {
			cumSigStrength += cycle * xRegister
		}

		// Determine the state of the current pixel
		if pixelPos - xRegister <= 1 && pixelPos - xRegister >= -1 {
			// Draw a lit pixel
			fmt.Print("#")
		} else {
			// Draw a dark pixel
			fmt.Print(".")
		}

		if !processingAddx {
			line := <- reader
			if len(line) == 0 {
				// We've exhausted the commands
				break
			}
			splitLine := strings.Split(line, " ")
			if splitLine[0] == "noop" {
				// don't do anything
			} else if splitLine[0] == "addx" {
				intVal, _ := strconv.Atoi(splitLine[1])
				temp = intVal
				processingAddx = true
			}
		} else {
			xRegister += temp
			processingAddx = false
		}

		if cycle % 40 == 0 {
			// Move to the next row of pixels
			fmt.Println()
			pixelPos = 0
		} else {
			pixelPos++
		}
		cycle++
	}
	fmt.Println("The cumulative signal strength is", cumSigStrength)
	fmt.Println()
}
