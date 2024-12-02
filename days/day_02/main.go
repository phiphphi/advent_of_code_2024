package day_02

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

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	count := 0

	for _, line := range input {
		nums := strings.Split(line, " ")
		order := 0 // 1 if increasing, -1 if decreasing
		good := true

	line_loop:
		for i := 0; i < len(nums); i++ {
			if i == 0 {
				continue
			}

			num, _ := strconv.Atoi(nums[i])
			prev, _ := strconv.Atoi(nums[i-1])

			if num > prev {
				if order == -1 {
					good = false
					break line_loop
				}

				order = 1
			} else if num < prev {
				if order == 1 {
					good = false
					break line_loop
				}

				order = -1
			} else {
				good = false
				break line_loop
			}

			diff := 0

			if order == 1 {
				diff = num - prev
			} else if order == -1 {
				diff = (num - prev) * -1
			}

			if diff < 1 || diff > 3 {
				good = false
				break line_loop
			}
		}

		if good {
			count += 1
		}
	}

	return strconv.Itoa(count)
}

func checkNums(num int, prev int, order int) (int, bool) {
	good := true

	fmt.Printf("Order: %v", order)

	if num > prev {
		if order == -1 {
			good = false
		}

		order = 1
	} else if num < prev {
		if order == 1 {
			good = false
		}

		order = -1
	} else {
		good = false
	}

	diff := 0

	if order == 1 {
		diff = num - prev
	} else if order == -1 {
		diff = (num - prev) * -1
	}

	if diff < 1 || diff > 3 {
		good = false
	}

	fmt.Printf("checked %v against %v, %v\n", num, prev, good)

	return order, good
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	count := 0

	for _, line := range input {
		nums := strings.Split(line, " ")
		good := false

		// Test copies without one.
	line_loop:
		for i := 0; i < len(nums); i++ {
			copy_good := true
			order := 0 // 1 if increasing, -1 if decreasing
			copy := make([]string, 0, len(nums)-1)
			copy = append(copy, nums[:i]...)
			copy = append(copy, nums[i+1:]...)
			fmt.Printf("%v\n", copy)

			for j := 1; j < len(copy); j++ {
				num, _ := strconv.Atoi(copy[j])
				prev, _ := strconv.Atoi(copy[j-1])

				order, copy_good = checkNums(num, prev, order)

				fmt.Println(num)
				if !copy_good {
					break
				}
			}

			fmt.Println()

			if copy_good {
				good = true
				break line_loop
			}
		}

		fmt.Printf("status: %v\n\n\n", good)

		if good {
			count += 1
		}
	}

	return strconv.Itoa(count)
}
