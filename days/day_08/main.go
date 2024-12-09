package day_08

import (
	"fmt"
	"strconv"
	"strings"
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

func buildMappings(input []string) map[string][]string {
	mappings := make(map[string][]string)

	for row, line := range input {
		for col, char := range line {
			if char == '.' {
				continue
			}

			index := fmt.Sprintf("%d,%d", row, col)

			strChar := string(char)
			spots, ok := mappings[strChar]
			if !ok {
				spots = make([]string, 0)
			}
			spots = append(spots, index)
			mappings[strChar] = spots
		}
	}

	return mappings
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func getCombinations(indexes []string) [][]string {
	var combinations [][]string

	for i := 0; i < len(indexes); i++ {
		for j := i + 1; j < len(indexes); j++ {
			combinations = append(combinations, []string{indexes[i], indexes[j]})
		}
	}

	return combinations
}

// Helper function to parse coordinates from a string "x,y"
func parseCoordinates(coord string) (int, int) {
	parts := strings.Split(coord, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}

// Helper function to generate possible antinodes
func generateAntinodes(x1, y1, x2, y2 int, vector1, vector2 []int) [][]int {
	return [][]int{
		{x1 - vector1[0], y1 - vector1[1]},
		{x2 - vector2[0], y2 - vector2[1]},
	}
}

func antiNodeValid(antiNode []int, antiNodeLocations map[string]bool, dimensions []int) bool {
	// Invalid if out of bounds or on another antenna.
	x, y := antiNode[0], antiNode[1]
	spot := fmt.Sprintf("%d,%d", x, y)
	onAntinodeSpot := antiNodeLocations[spot]

	return !(x < 0 || x >= dimensions[0] || y < 0 || y >= dimensions[1] || onAntinodeSpot)
}

func countAntinodes(mapping map[string][]string, dimensions []int) int {
	count := 0
	antiNodeLocations := make(map[string]bool)

	for _, values := range mapping {
		combinations := getCombinations(values)

		for _, combination := range combinations {
			x1, y1 := parseCoordinates(combination[0])
			x2, y2 := parseCoordinates(combination[1])

			vector1 := []int{x2 - x1, y2 - y1}
			vector2 := []int{x1 - x2, y1 - y2}

			possibleAntinodes := generateAntinodes(x1, y1, x2, y2, vector1, vector2)

			for _, possibleAntinode := range possibleAntinodes {
				if antiNodeValid(possibleAntinode, antiNodeLocations, dimensions) {
					antiNodeLocations[fmt.Sprintf("%d,%d", possibleAntinode[0], possibleAntinode[1])] = true
					count++
				}
			}
		}
	}

	return count
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	mappings := buildMappings(input)
	dimensions := []int{len(input), len(input[0])}

	return strconv.Itoa(countAntinodes(mappings, dimensions))
}

func generateResonantAntinodes(x1, y1, x2, y2 int, vector1, vector2 []int, dimensions []int) [][]int {
	antiNodes := make([][]int, 0)

	antiNodes = loopAntinodes(antiNodes, x1, y1, vector1, dimensions)
	antiNodes = loopAntinodes(antiNodes, x2, y2, vector2, dimensions)

	return antiNodes
}

func loopAntinodes(antiNodes [][]int, x, y int, vector, dimensions []int) [][]int {
	increment := 0
	for {
		xNode := x - (vector[0] * increment)
		yNode := y - (vector[1] * increment)

		if xNode < 0 || xNode >= dimensions[0] || yNode < 0 || yNode >= dimensions[1] {
			break
		}

		antiNode := []int{xNode, yNode}
		antiNodes = append(antiNodes, antiNode)
		increment++
	}

	return antiNodes
}

func countResonantAntinodes(mapping map[string][]string, dimensions []int) int {
	count := 0
	antiNodeLocations := make(map[string]bool)

	for _, values := range mapping {
		combinations := getCombinations(values)

		for _, combination := range combinations {
			x1, y1 := parseCoordinates(combination[0])
			x2, y2 := parseCoordinates(combination[1])

			vector1 := []int{x2 - x1, y2 - y1}
			vector2 := []int{x1 - x2, y1 - y2}

			possibleAntinodes := generateResonantAntinodes(x1, y1, x2, y2, vector1, vector2, dimensions)

			for _, possibleAntinode := range possibleAntinodes {
				if antiNodeValid(possibleAntinode, antiNodeLocations, dimensions) {
					antiNodeLocations[fmt.Sprintf("%d,%d", possibleAntinode[0], possibleAntinode[1])] = true
					count++
				}
			}
		}
	}

	return count
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	mappings := buildMappings(input)
	dimensions := []int{len(input), len(input[0])}

	return strconv.Itoa(countResonantAntinodes(mappings, dimensions))
}
