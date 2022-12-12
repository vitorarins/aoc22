package day8

import (
	"strconv"
	"strings"
)

func ParseInputTrees(input string) [][]int {

	input = strings.TrimSpace(input)

	// Split the input string into a slice of strings, one for each row of the grid.
	rows := strings.Split(input, "\n")

	// Convert each string into a slice of integers representing the heights of the trees in the row.
	grid := make([][]int, len(rows))
	for i, row := range rows {
		grid[i] = make([]int, len(row))
		for j, r := range row {
			height, _ := strconv.Atoi(string(r))
			grid[i][j] = height
		}
	}

	return grid
}

func GetTrees(input string) (int, int) {
	// Keep track of the maximum scenic score.
	maxScore := 0

	// Parse the input string into a 2D grid of integers.
	grid := ParseInputTrees(input)

	// Count the number of trees that are visible from the edge
	var edgeTrees int

	// Iterate over each row and column in the grid.
	var interiorTrees int
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			// Skip trees that are on the edge of the grid.
			if row == 0 || row == len(grid)-1 || col == 0 || col == len(grid[0])-1 {
				edgeTrees++
				continue
			}

			// Keep track of the distances to the nearest trees that are the same height or taller in each direction.
			distances := [4]int{}

			// Calculate the distances to the nearest trees in the left direction.
			leftVisible := true
			for j := col - 1; j >= 0; j-- {
				if grid[row][j] >= grid[row][col] {
					distances[2] = col - j
					leftVisible = false
					break
				}
				distances[2] += 1
			}

			// Calculate the distances to the nearest trees in the right direction.
			rightVisible := true
			for j := col + 1; j < len(grid[0]); j++ {
				if grid[row][j] >= grid[row][col] {
					distances[3] = j - col
					rightVisible = false
					break
				}
				distances[3] += 1
			}

			// Calculate the distances to the nearest trees in the up direction.
			topVisible := true
			for i := row - 1; i >= 0; i-- {
				if grid[i][col] >= grid[row][col] {
					distances[0] = row - i
					topVisible = false
					break
				}
				distances[0] += 1
			}

			// Calculate the distances to the nearest trees in the down direction.
			bottomVisible := true
			for i := row + 1; i < len(grid); i++ {
				if grid[i][col] >= grid[row][col] {
					distances[1] = i - row
					bottomVisible = false
					break
				}
				distances[1] += 1
			}

			// If the tree is visible from any direction, increment the counter
			if leftVisible || rightVisible || topVisible || bottomVisible {
				interiorTrees++
			}

			// Compute the scenic score for this tree by multiplying together the distances in each direction.
			score := 1
			for _, distance := range distances {
				score *= distance
			}

			// Ignore trees with a scenic score of 0.
			if score == 0 {
				continue
			}

			// Update the maximum scenic score if this tree has a higher score.
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return (edgeTrees + interiorTrees), maxScore
}

func GetVisibleTrees(input string) int {
	vTrees, _ := GetTrees(input)

	return vTrees
}

func GetScenic(input string) int {
	_, score := GetTrees(input)

	return score
}
