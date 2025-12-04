package day02

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
	"strings"
)

type idRange struct {
	start int
	end   int
}

func parseInput(input string) []idRange {
	input = strings.TrimSpace(input)
	ranges := strings.Split(input, ",")

	idRanges := make([]idRange, len(ranges))

	for i, r := range ranges {
		split := strings.Split(r, "-")

		if len(split) != 2 {
			panic(fmt.Sprintf("Malformed input at range %d: %s", i, r))
		}

		start, err := strconv.Atoi(split[0])
		if err != nil {
			panic(fmt.Sprintf("Malformed input at range %d: %s", i, r))
		}

		end, err := strconv.Atoi(split[1])
		if err != nil {
			panic(fmt.Sprintf("Malformed input at range %d: %s", i, r))
		}

		idRanges[i].start = start
		idRanges[i].end = end
	}

	return idRanges
}

func allSlicesEqual[E comparable](seq iter.Seq[[]E]) bool {
	var firstSlice []E
	foundFirst := false

	for currentSlice := range seq {
		if !foundFirst {
			firstSlice = currentSlice
			foundFirst = true
			continue
		}

		if !slices.Equal(firstSlice, currentSlice) {
			return false
		}
	}
	return true
}

func SolvePartOne(input string) {
	result := 0
	for _, r := range parseInput(input) {
		for i := r.start; i <= r.end; i++ {
			str := strconv.Itoa(i)

			middle := len(str) / 2

			if string(str[:middle]) == string(str[middle:]) {
				result += i
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func SolvePartTwo(input string) {
	result := 0
	for _, r := range parseInput(input) {
		for i := r.start; i <= r.end; i++ {
			str := strings.Split(strconv.Itoa(i), "")

			for j := 1; j <= len(str)/2; j++ {
				chunks := slices.Chunk(str, j)

				if allSlicesEqual(chunks) {
					result += i
					break
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}
