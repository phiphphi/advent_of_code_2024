package day_07

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

func checkEquation(result int, values []string) bool {
	value, _ := strconv.Atoi(values[0])
	return recursiveWorks(result, value, values[1:], "*") || recursiveWorks(result, value, values[1:], "+")
}

func recursiveWorks(result int, current int, values []string, operator string) bool {
	if len(values) == 0 {
		return false
	}

	nextValue, _ := strconv.Atoi(values[0])

	if operator == "*" {
		current *= nextValue
	} else if operator == "+" {
		current += nextValue
	}

	if current == result && len(values[1:]) == 0 {
		return true
	}

	return recursiveWorks(result, current, values[1:], "*") || recursiveWorks(result, current, values[1:], "+")
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	count := 0

	for _, line := range input {
		values := strings.Split(line, ": ")
		result, _ := strconv.Atoi(values[0])
		numbers := strings.Split(values[1], " ")

		if checkEquation(result, numbers) {
			count += result
		}
	}

	return strconv.Itoa(count)
}

func checkEquation2(result int, values []string) bool {
	value, _ := strconv.Atoi(values[0])
	return recursiveWorks2(result, value, values[1:], "*") ||
		recursiveWorks2(result, value, values[1:], "+") ||
		recursiveWorks2(result, value, values[1:], "||")
}

func recursiveWorks2(result int, current int, values []string, operator string) bool {
	if len(values) == 0 {
		return false
	}

	nextValue, _ := strconv.Atoi(values[0])

	if operator == "*" {
		current *= nextValue
	} else if operator == "+" {
		current += nextValue
	} else if operator == "||" {
		currentStr := strconv.Itoa(current)
		nextStr := strconv.Itoa(nextValue)
		current, _ = strconv.Atoi(currentStr + nextStr)
	}

	if current == result && len(values[1:]) == 0 {
		return true
	}

	return recursiveWorks2(result, current, values[1:], "*") ||
		recursiveWorks2(result, current, values[1:], "+") ||
		recursiveWorks2(result, current, values[1:], "||")
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	count := 0

	for _, line := range input {
		values := strings.Split(line, ": ")
		result, _ := strconv.Atoi(values[0])
		numbers := strings.Split(values[1], " ")

		if checkEquation2(result, numbers) {
			count += result
		}
	}

	return strconv.Itoa(count)
}
