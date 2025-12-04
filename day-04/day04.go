package day04

import (
	"fmt"
	"strings"
)

func parseInput(input string) [][]string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := [][]string{}

	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func SolvePartOne(input string) {
	grid := parseInput(input)
	maxRow := len(grid)
	maxCol := len(grid[0])
	result := 0

	for row := range maxRow {
		for col := range maxCol {
			if grid[row][col] != "@" {
				continue
			}

			nei := [][]int{
				{row - 1, col},
				{row + 1, col},
				{row - 1, col - 1},
				{row + 1, col + 1},
				{row, col + 1},
				{row, col - 1},
				{row + 1, col - 1},
				{row - 1, col + 1},
			}

			count := 0

			for _, n := range nei {
				r := n[0]
				c := n[1]

				if r < 0 || r >= maxRow || c < 0 || c >= maxCol {
					continue
				}

				if grid[r][c] == "@" {
					count += 1
				}

				if count == 4 {
					break
				}
			}

			if count < 4 {
				result += 1
				// fmt.Printf("Marked: x=%d y=%d\n", row, col)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func SolvePartTwo(input string) {
	grid := parseInput(input)
	maxRow := len(grid)
	maxCol := len(grid[0])
	result := 0

	wasRemoved := false

	for {
		for row := range maxRow {
			for col := range maxCol {
				if grid[row][col] != "@" {
					continue
				}

				nei := [][]int{
					{row - 1, col},
					{row + 1, col},
					{row - 1, col - 1},
					{row + 1, col + 1},
					{row, col + 1},
					{row, col - 1},
					{row + 1, col - 1},
					{row - 1, col + 1},
				}

				count := 0

				for _, n := range nei {
					r := n[0]
					c := n[1]

					if r < 0 || r >= maxRow || c < 0 || c >= maxCol {
						continue
					}

					if grid[r][c] == "@" {
						count += 1
					}

					if count == 4 {
						break
					}
				}

				if count < 4 {
					result += 1
					grid[row][col] = "x"
					wasRemoved = true
					// fmt.Printf("Marked: x=%d y=%d\n", row, col)
				}

			}
		}

		if !wasRemoved {
			break
		}
		wasRemoved = false
	}

	fmt.Printf("Part 2: %d\n", result)
}
