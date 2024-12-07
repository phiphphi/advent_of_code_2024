package day_06

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

func parseMap(input []string) (map[string]bool, []int) {
	obstacles := make(map[string]bool)
	var start []int

	for row, line := range input {
		for col, char := range line {
			if char == '#' {
				key := fmt.Sprintf("%d,%d", row, col)
				obstacles[key] = true
			} else if char == '^' {
				start = []int{row, col}
			}
		}
	}

	return obstacles, start
}

func walkMap(input []string, obstacles map[string]bool, start []int) map[string]bool {
	visited := make(map[string]bool)
	currentDir := 0
	directions := [][]int{
		[]int{-1, 0},
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
	}

	return walk(input, obstacles, visited, directions, currentDir, start)
}

func walk(input []string, obstacles map[string]bool, visited map[string]bool, directions [][]int, currentDir int, pos []int) map[string]bool {
	spot := fmt.Sprintf("%d,%d", pos[0], pos[1])
	visited[spot] = true

	direction := directions[currentDir%4]
	nextPos := []int{pos[0] + direction[0], pos[1] + direction[1]}
	nextSpot := fmt.Sprintf("%d,%d", nextPos[0], nextPos[1])

	// Exit if we're leaving the map.
	if nextPos[0] < 0 || nextPos[0] >= len(input) || nextPos[1] < 0 || nextPos[1] >= len(input[0]) {
		return visited
	}

	// Rotate if we're walking into an obstacle.
	if obstacles[nextSpot] {
		currentDir += 1
		nextPos = pos
	}

	return walk(input, obstacles, visited, directions, currentDir, nextPos)
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	obstacles, start := parseMap(input)

	visited := walkMap(input, obstacles, start)

	return strconv.Itoa(len(visited))
}

func walkMapLoops(input []string, obstacles map[string]bool, start []int) int {
	currentDir := 0
	visited := make(map[string]bool)
	directions := [][]int{
		[]int{-1, 0},
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
	}

	return walkLoops(input, obstacles, directions, currentDir, start, 0, visited)
}

func walkLoops(input []string, obstacles map[string]bool, directions [][]int, currentDir int, pos []int, count int, visited map[string]bool) int {
	direction := directions[currentDir%4]
	nextPos := []int{pos[0] + direction[0], pos[1] + direction[1]}
	nextSpot := fmt.Sprintf("%d,%d", nextPos[0], nextPos[1])

	spot := fmt.Sprintf("%d,%d", pos[0], pos[1])
	visited[spot] = true

	// Exit if we're leaving the map.
	if nextPos[0] < 0 || nextPos[0] >= len(input) || nextPos[1] < 0 || nextPos[1] >= len(input[0]) {
		return count
	}

	// Rotate if we're walking into an obstacle.
	if obstacles[nextSpot] {
		currentDir += 1
		direction := directions[currentDir%4]
		nextPos = []int{pos[0] + direction[0], pos[1] + direction[1]}
		nextSpot = fmt.Sprintf("%d,%d", nextPos[0], nextPos[1])
	}

	// Don't add obstacles for the original location or locations we've visited.
	if !visited[nextSpot] {
		count += checkLoopSetup(input, obstacles, directions, currentDir, pos, nextPos)
	}

	return walkLoops(input, obstacles, directions, currentDir, nextPos, count, visited)
}

func checkLoopSetup(input []string, obstacles map[string]bool, directions [][]int, currentDir int, pos []int, nextPos []int) int {
	// Check for a loop if we put an obstacle in the next position.
	nextSpot := fmt.Sprintf("%d,%d", nextPos[0], nextPos[1])
	obstaclesCopy := make(map[string]bool)
	for k, v := range obstacles {
		obstaclesCopy[k] = v
	}
	obstaclesCopy[nextSpot] = true

	// Make a sub-visited map for checking loops.
	visited := make(map[string]bool)

	if checkForLoop(visited, input, obstaclesCopy, directions, currentDir, pos, currentDir, pos) {
		return 1
	}

	return 0
}

func checkForLoop(visited map[string]bool, input []string, obstacles map[string]bool, directions [][]int, currentDir int, pos []int, startDir int, start []int) bool {
	direction := directions[currentDir%4]
	nextPos := []int{pos[0] + direction[0], pos[1] + direction[1]}
	nextSpot := fmt.Sprintf("%d,%d", nextPos[0], nextPos[1])

	// Exit if we've looped, same position and same direction.
	posAndDirection := fmt.Sprintf("%d,%d,%d", pos[0], pos[1], direction)
	if visited[posAndDirection] {
		return true
	}
	visited[posAndDirection] = true

	// Exit if we're leaving the map.
	if nextPos[0] < 0 || nextPos[0] >= len(input) || nextPos[1] < 0 || nextPos[1] >= len(input[0]) {
		return false
	}

	// Rotate if we're walking into an obstacle.
	if obstacles[nextSpot] {
		currentDir += 1
		nextPos = pos
	}

	return checkForLoop(visited, input, obstacles, directions, currentDir, nextPos, startDir, start)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	obstacles, start := parseMap(input)

	return strconv.Itoa(walkMapLoops(input, obstacles, start))
}
