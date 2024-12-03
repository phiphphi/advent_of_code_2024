package day_03

import (
	"fmt"
	"regexp"
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

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	result := 0
	instructions := strings.Join(input, "")

	mulRegex, _ := regexp.Compile("mul\\(([0-9]+),([0-9]+)\\)")

	matches := mulRegex.FindAllStringSubmatch(instructions, -1)

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		fmt.Println(match)
		result += num1 * num2
	}

	fmt.Println()

	return strconv.Itoa(result)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	result := 0
	enabled := true
	instructions := strings.Join(input, "")

	mulRegex, _ := regexp.Compile("mul\\(([0-9]+),([0-9]+)\\)|do\\(\\)|don't\\(\\)")

	matches := mulRegex.FindAllStringSubmatch(instructions, -1)

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			fmt.Println(match)
			result += num1 * num2
		}
	}

	return strconv.Itoa(result)
}
