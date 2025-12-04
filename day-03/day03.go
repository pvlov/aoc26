package day03

import (
	"fmt"
	"strings"
)

func maximizeNDigits(input string, numDigits int) int {
	const zeroByte = '0'
	totalResult := 0

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		L := len(line)

		if L < numDigits {
			panic(fmt.Sprintf("Skipping line: too short for %d digits\n", numDigits))
		}

		currentVal := 0
		searchStart := 0

		for k := range numDigits {

			remainingNeeded := (numDigits - 1) - k

			searchEnd := L - remainingNeeded
			bestIdx := searchStart

			for i := searchStart; i < searchEnd; i++ {
				if line[i] == '9' {
					bestIdx = i
					break
				}

				if line[i] > line[bestIdx] {
					bestIdx = i
				}
			}

			digit := int(line[bestIdx] - zeroByte)
			currentVal = (currentVal * 10) + digit
			searchStart = bestIdx + 1
		}

		totalResult += currentVal
	}

	return totalResult
}

func SolvePartOne(input string) {
	result := maximizeNDigits(input, 2)
	fmt.Printf("Part 1: %d\n", result)
}

func SolvePartTwo(input string) {
	result := maximizeNDigits(input, 12)
	fmt.Printf("Part 2: %d\n", result)
}
