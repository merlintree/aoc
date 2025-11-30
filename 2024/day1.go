package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseDay1Data(data string) ([]int, []int) {

	lines := strings.Split(strings.TrimSpace(data), "\n")

	var left []int
	var right []int

	for _, line := range lines {
		parts := strings.Fields(line)
		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	return left, right
}

func day1part1(data string) int {

	left, right := parseDay1Data(data)

	sort.Ints(left)
	sort.Ints(right)

	var distance int

	for i, _ := range left {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	return distance

}

func day1part2(data string) int {

	left, right := parseDay1Data(data)

	var totalSimilarity int

	for l, _ := range left {
		similarity := 0
		for r, _ := range right {
			if right[r] == left[l] {
				similarity++
			}
		}
		totalSimilarity += similarity * left[l]
	}

	return totalSimilarity
}
