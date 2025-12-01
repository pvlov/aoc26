package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type rotation struct {
	direction byte
	amount    int
}

func parseInput(input string) []rotation {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	rotations := make([]rotation, len(lines))

	for i, line := range lines {
		direction := line[0]
		amount, err := strconv.Atoi(line[1:])
		if (direction != 'L' && direction != 'R') || err != nil {
			panic(fmt.Sprintf("Malformed input at line %d: %s", i+1, line))
		}

		rotations[i].direction = direction
		rotations[i].amount = amount
	}

	return rotations
}

func SolvePartOne(input string) {
	rotations := parseInput(input)

	result := 0
	dial := 50

	for _, rot := range rotations {
		switch rot.direction {
		case 'L':
			if rot.amount > dial {
				dial = ((dial-rot.amount)%100 + 100) % 100
			} else {
				dial -= rot.amount
			}

		case 'R':
			dial = (dial + rot.amount) % 100
		}

		if dial == 0 {
			result += 1
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func SolvePartTwo(input string) {
	rotations := parseInput(input)

	result := 0
	dial := 50

	for _, rot := range rotations {
		result += rot.amount / 100

		switch rot.direction {
		case 'L':
			if rot.amount > dial {
				if (rot.amount%100) > dial && dial != 0 {
					result += 1
				}

				dial = ((dial-rot.amount)%100 + 100) % 100
			} else {
				dial -= rot.amount
			}

		case 'R':
			if (rot.amount % 100) > 100-dial {
				result += 1
			}

			dial = (dial + rot.amount) % 100
		}

		if dial == 0 {
			result += 1
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}
