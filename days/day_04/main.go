package day_04

import (
	"fmt"
	"strconv"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

func checkMatch(input []string, row int, col int, rowModifier int, colModifier int) bool {
	mapping := "XMAS"

	for key, value := range mapping {
		newRow := row + (key * rowModifier)
		newCol := col + (key * colModifier)

		// Row check
		if newRow < 0 || newRow >= len(input) {
			return false
		}

		// Col check
		if newCol < 0 || newCol >= len(input[newRow]) {
			return false
		}

		// Char match
		line := []rune(input[newRow])
		if line[newCol] != value {
			return false
		}

		if key == 3 {
			fmt.Printf("Match at %v, %v, %v, %v\n", row, col, rowModifier, colModifier)
			return true
		}
	}

	return false
}

func checkMatches(input []string, row int, col int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if checkMatch(input, row, col, i, j) {
				count++
			}
		}
	}

	return count
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	count := 0

	for row, line := range input {
		for col, letter := range line {
			if letter == 'X' {
				count += checkMatches(input, row, col)
			}
		}
	}

	return strconv.Itoa(count)
}

func checkDiagonals(NW rune, NE rune, SW rune, SE rune, top rune, bottom rune) bool {
	return (NW == top && NE == top && SW == bottom && SE == bottom) ||
		(NW == top && SW == top && NE == bottom && SE == bottom) ||
		(SW == top && SE == top && NE == bottom && NW == bottom) ||
		(SE == top && NE == top && NW == bottom && SW == bottom)

}

func checkMasMatch(input []string, row int, col int) bool {
	if row < 1 || row >= len(input)-1 || col < 1 || col >= len(input[row])-1 {
		return false
	}

	NW := rune(input[row-1][col-1])
	NE := rune(input[row-1][col+1])
	SW := rune(input[row+1][col-1])
	SE := rune(input[row+1][col+1])

	return checkDiagonals(NW, NE, SW, SE, 'S', 'M')
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	count := 0

	for row, line := range input {
		for col, letter := range line {
			if letter == 'A' {
				if checkMasMatch(input, row, col) {
					count++
				}
			}
		}
	}

	return strconv.Itoa(count)
}
