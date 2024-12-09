package day_09

import (
	"fmt"
	"slices"
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

// Keep moving the pointer until we find a non-zero value.
func movePointer(line string, pointer int, increment int) (int, int) {
	value := 0

	for {
		pointer += increment
		value = int(line[pointer] - '0')

		if value > 0 {
			break
		}
	}

	return pointer, value
}

func computeChecksum(line string) int {
	count := 0

	dataIndex := 0

	startPointer := 0
	startPointerValue := int(line[startPointer] - '0')

	endPointer := len(line) - 1
	endPointerValue := int(line[endPointer] - '0')
	for {
		// Move pointers when they run out.
		if startPointerValue == 0 {
			startPointer, startPointerValue = movePointer(line, startPointer, 1)
		}

		if endPointerValue == 0 {
			endPointer, endPointerValue = movePointer(line, endPointer, -2)
		}

		if startPointer >= endPointer {
			// Special case here: what if startPointer and endPointer are working on the same block?

			if startPointer == endPointer {
				fileDataId := startPointer / 2
				diff := slices.Min([]int{startPointerValue, endPointerValue})

				for i := 0; i < diff; i++ {
					count += fileDataId * dataIndex
					dataIndex += 1
				}
			}

			break
		}

		// Counting file data.
		if startPointer%2 == 0 {
			fileDataId := startPointer / 2

			count += dataIndex * fileDataId

			startPointerValue -= 1
		} else { // Counting through empty data from the end.
			fileDataId := endPointer / 2

			count += dataIndex * fileDataId

			startPointerValue -= 1
			endPointerValue -= 1
		}

		dataIndex += 1
	}

	return count
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	line := input[0]

	return strconv.Itoa(computeChecksum(line))
}

// Keep moving the pointer until we find a non-zero value.
func movePointer(line string, pointer int, increment int) (int, int) {
	value := 0

	for {
		pointer += increment
		value = int(line[pointer] - '0')

		if value > 0 {
			break
		}
	}

	return pointer, value
}

func computeChecksum2(line string) int {
	count := 0

	data := ""

	dataIndex := 0

	startPointer := 0
	startPointerValue := int(line[startPointer] - '0')

	endPointer := len(line) - 1
	endPointerValue := int(line[endPointer] - '0')
	for {
		// Move pointers when they run out.
		if startPointerValue == 0 {
			startPointer, startPointerValue = movePointer(line, startPointer, 1)
		}

		// TODO: search for end pointer value gap to fill
		if endPointerValue == 0 {
			endPointer, endPointerValue = movePointer(line, endPointer, -2)
		}

		if startPointer >= endPointer {
			// Special case here: what if startPointer and endPointer are working on the same block?
			fmt.Printf("break case: %v %v %v %v\n", startPointer, endPointer, startPointerValue, endPointerValue)

			if startPointer == endPointer {
				fileDataId := startPointer / 2
				diff := slices.Min([]int{startPointerValue, endPointerValue})

				for i := 0; i < diff; i++ {
					data += strconv.Itoa(fileDataId)
					count += fileDataId * dataIndex
					dataIndex += 1
				}
			}

			break
		}

		fmt.Printf("%v %v %v %v\n", startPointer, endPointer, startPointerValue, endPointerValue)

		// Counting file data.
		if startPointer%2 == 0 {
			fileDataId := startPointer / 2

			fmt.Printf("file data add %v, total %v\n", dataIndex*fileDataId, count)
			count += dataIndex * fileDataId

			data += strconv.Itoa(fileDataId)

			startPointerValue -= 1
		} else { // Counting through empty data from the end.
			fileDataId := endPointer / 2

			fmt.Printf("file empty add %v, total %v\n", dataIndex*fileDataId, count)
			count += dataIndex * fileDataId

			data += strconv.Itoa(fileDataId)

			startPointerValue -= 1
			endPointerValue -= 1
		}

		dataIndex += 1
	}

	fmt.Println(data)

	return count
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	line := input[0]

	return strconv.Itoa(computeChecksum2(line))
}
