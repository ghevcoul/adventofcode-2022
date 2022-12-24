package day08

import (
	"fmt"
	"ghevcoul/aoc22/utils"
	"strconv"
	"strings"
)

func GenerateMatrixFromInput(path string) [][]int {
	reader, err := utils.Readlines(path)
	if err != nil {
		fmt.Println(err)
		return make([][]int, 0)
	}

	// TODO: Figure out the length of the arrays
	matrix := make([][]int, 0)

	for line := range reader {
		vals := strings.Split(line, "")
		row := make([]int, len(vals))
		for i := 0; i < len(vals); i++ {
			height, _ := strconv.Atoi(vals[i])
			row[i] = height
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func CountVisible(heights [][]int) int {
	visible := 0
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			// Increment visible and exit early if on an edge
			if i == 0 || i == len(heights)-1 || j == 0 || j == len(heights[i])-1 {
				visible++
				continue
			}

			// North is decreasing i
			north := true
			for n := i-1; n >= 0; n-- {
				if heights[n][j] >= heights[i][j] {
					north = false
					break
				}
			}
			// West is decreasing j
			west := true
			for w := j-1; w >= 0; w-- {
				if heights[i][w] >= heights[i][j] {
					west = false
					break
				}
			}
			// South is increasing i
			south := true
			for s := i+1; s < len(heights); s++ {
				if heights[s][j] >= heights[i][j] {
					south = false
					break
				}
			}
			// East is increasing j
			east := true
			for e := j+1; e < len(heights[i]); e++ {
				if heights[i][e] >= heights[i][j] {
					east = false
					break
				}
			}
			// Check if visible in any direction
			if north || east || south || west {
				visible++
			}
		}
	}
	return visible
}

func FindMostScenic(heights [][]int) int {
	mostScenic := 0
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			// If on an edge, one scenic score will be zero
			// Assume there's better out there and exit early
			if i == 0 || i == len(heights)-1 || j == 0 || j == len(heights[i])-1 {
				continue
			}

			// North is decreasing i
			northScore := 0
			for n := i-1; n >= 0; n-- {
				northScore++
				if heights[n][j] >= heights[i][j] {
					break
				}
			}
			// West is decreasing j
			westScore := 0
			for w := j-1; w >= 0; w-- {
				westScore++
				if heights[i][w] >= heights[i][j] {
					break
				}
			}
			// South is increasing i
			southScore := 0
			for s := i+1; s < len(heights); s++ {
				southScore++
				if heights[s][j] >= heights[i][j] {
					break
				}
			}
			// East is increasing j
			eastScore := 0
			for e := j+1; e < len(heights[i]); e++ {
				eastScore++
				if heights[i][e] >= heights[i][j] {
					break
				}
			}
			score := northScore * eastScore * southScore * westScore
			if score > mostScenic {
				mostScenic = score
			}
		}
	}
	return mostScenic
}

func Day08() {
	fmt.Println("* * * * * * * * * * Day 08 * * * * * * * * * *")
	dataPath := "inputs/day08"
	heightMap := GenerateMatrixFromInput(dataPath)

	totalVisible := CountVisible(heightMap)
	fmt.Println("The total trees visible are", totalVisible)

	mostScenic := FindMostScenic(heightMap)
	fmt.Println("The most scenic tree has a score of", mostScenic)
}