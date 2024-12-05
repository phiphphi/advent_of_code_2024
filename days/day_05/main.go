package day_05

import (
	"fmt"
	"slices"
	"sort"
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

func checkLineValid(mapping map[string][]string, line string) bool {
	parts := strings.Split(line, ",")
	nextParts, ok := mapping[parts[0]]
	if !ok {
		return false
	}

	for i := 1; i < len(parts); i++ {
		if !slices.Contains(nextParts, parts[i]) {
			return false
		}

		potentialNextParts, _ := mapping[parts[i]]
		nextParts = potentialNextParts
	}

	return true
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	count := 0
	mapping := make(map[string][]string)

	for _, line := range input {
		if strings.Contains(line, "|") {
			// Handle building mapping of pages.
			parts := strings.Split(line, "|")
			mapping[parts[0]] = append(mapping[parts[0]], parts[1])
		} else if strings.Contains(line, ",") {
			if checkLineValid(mapping, line) {
				parts := strings.Split(line, ",")
				middle, _ := strconv.Atoi(parts[len(parts)/2])
				count += middle
			}
		}
	}

	return strconv.Itoa(count)
}

func sortInvalidLine(mapping map[string][]string, line string) []string {

	return make([]string, 0)
}

func buildPriorityList(mapping map[string][]string, parts []string) map[string]int {
	priorityList := make(map[string]int)

	// Count for each part, how many dependencies it has in parts.
	for _, part := range parts {
		count := 0
		dependencies := mapping[part]

		for _, dependency := range dependencies {
			if slices.Contains(parts, dependency) {
				count += 1
			}
		}

		priorityList[part] = count
	}

	return priorityList
}

func getPriorityListValue(priorityList map[string]int, key string) int {
	value, ok := priorityList[key]

	// Default to 0 if not in map, this means it has no dependencies and can go at the end.
	if !ok {
		return 0
	}

	return value
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	count := 0
	mapping := make(map[string][]string)

	for _, line := range input {
		if strings.Contains(line, "|") {
			// Handle building mapping of pages.
			parts := strings.Split(line, "|")
			mapping[parts[0]] = append(mapping[parts[0]], parts[1])
		} else if strings.Contains(line, ",") {
			if !checkLineValid(mapping, line) {
				parts := strings.Split(line, ",")
				priorityList := buildPriorityList(mapping, parts)

				sort.Slice(parts, func(i, j int) bool {
					return getPriorityListValue(priorityList, parts[i]) > getPriorityListValue(priorityList, parts[j])
				})

				middle, _ := strconv.Atoi(parts[len(parts)/2])

				count += middle
			}
		}
	}

	fmt.Println(mapping)

	return strconv.Itoa(count)
}
