package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"aoc26/day-01"
	"aoc26/day-02"
	"aoc26/day-03"
	"aoc26/day-04"
	"aoc26/day-05"
	"aoc26/day-06"
	"aoc26/day-07"
	"aoc26/day-08"
	"aoc26/day-09"
	"aoc26/day-10"
	"aoc26/day-11"
	"aoc26/day-12"
)

func main() {
	day := flag.Int("day", 1, "The day to run")
	test := flag.Bool("test", false, "Whether to use the small test input")
	flag.Parse()

	var inputPath string

	if *test {
		inputPath = filepath.Join("inputs", fmt.Sprintf("%d_test.txt", *day))
	} else {
		inputPath = filepath.Join("inputs", fmt.Sprintf("%d.txt", *day))
	}

	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading input for day %d: %v\n", *day, err)
		os.Exit(1)
	}
	input := string(inputBytes)

	fmt.Printf("--- Running Day %d ---\n", *day)

	switch *day {
	case 1:
		day01.SolvePartOne(input)
		day01.SolvePartTwo(input)
	case 2:
		day02.SolvePartOne(input)
		day02.SolvePartTwo(input)
	case 3:
		day03.SolvePartOne(input)
		day03.SolvePartTwo(input)
	case 4:
		day04.SolvePartOne(input)
		day04.SolvePartTwo(input)
	case 5:
		day05.SolvePartOne(input)
		day05.SolvePartTwo(input)
	case 6:
		day06.SolvePartOne(input)
		day06.SolvePartTwo(input)
	case 7:
		day07.SolvePartOne(input)
		day07.SolvePartTwo(input)
	case 8:
		day08.SolvePartOne(input)
		day08.SolvePartTwo(input)
	case 9:
		day09.SolvePartOne(input)
		day09.SolvePartTwo(input)
	case 10:
		day10.SolvePartOne(input)
		day10.SolvePartTwo(input)
	case 11:
		day11.SolvePartOne(input)
		day11.SolvePartTwo(input)
	case 12:
		day12.SolvePartOne(input)
		day12.SolvePartTwo(input)
	default:
		fmt.Printf("Day %d is not supported!\n", *day)
	}
}
