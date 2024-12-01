package day_01

import (
	"fmt"
	"math"
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

func getTwoArrays(input []string) ([]int, []int) {
	var arr1 []int
	var arr2 []int

	for _, line := range input {
		nums := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		arr1 = append(arr1, num1)
		arr2 = append(arr2, num2)
	}

	return arr1, arr2
}

func getDistance(arr1 []int, arr2 []int) int {
	count := float64(0)

	for i := 0; i < len(arr1); i++ {
		count += math.Abs(float64(arr1[i]) - float64(arr2[i]))
	}

	return int(count)
}

func getSimilarity(arr1 []int, arr2 []int) int {
	mapping := make(map[int]int)
	count := 0

	for _, num := range arr2 {
		value, ok := mapping[num]
		if !ok {
			mapping[num] = 1
		} else {
			mapping[num] = value + 1
		}
	}

	for _, num := range arr1 {
		value, _ := mapping[num]
		count += value * num
	}

	return count
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	arr1, arr2 := getTwoArrays(input)

	sort.Sort(sort.IntSlice(arr1))
	sort.Sort(sort.IntSlice(arr2))

	return strconv.Itoa(getDistance(arr1, arr2))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	arr1, arr2 := getTwoArrays(input)

	return strconv.Itoa(getSimilarity(arr1, arr2))
}
