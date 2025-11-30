package main

import (
	"strconv"
	"strings"
)

func parseDay2Data(inputData string) [][]int {
	lines := strings.Split(strings.TrimSpace(inputData), "\n")

	data := make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		levels := make([]int, len(parts))

		for j, part := range parts {
			num, _ := strconv.Atoi(part)
			levels[j] = num
		}
		data[i] = levels
	}
	return data
}

func day2part1(inputData string) int {
	data := parseDay2Data(inputData)
	var safeCount int

	for _, levels := range data {
		isUnsafe := checkIsUnsafe(levels)
		if !isUnsafe {
			safeCount++
		}
	}
	return safeCount
}

func day2part2(inputData string) int {

	data := parseDay2Data(inputData)
	var safeCount int
	var unsafeList [][]int

	for _, levels := range data {
		isUnsafe := checkIsUnsafe(levels)
		if isUnsafe {
			unsafeList = append(unsafeList, levels)
			continue
		}
		safeCount++
	}

	for _, levels := range unsafeList {
		for i := 0; i < len(levels); i++ {
			updatedLevels := []int{}
			updatedLevels = append(updatedLevels, levels[:i]...)
			updatedLevels = append(updatedLevels, levels[i+1:]...)
			isUnsafe := checkIsUnsafe(updatedLevels)
			if !isUnsafe {
				safeCount++
				break
			}
		}
	}
	return safeCount
}

func checkIsUnsafe(level []int) bool {
	isUnsafe := false
	isIncreasing := false
	for j := 1; j < len(level); j++ {
		distance := level[j] - level[j-1]
		if distance == 0 || distance < -3 || distance > 3 {
			isUnsafe = true
			break
		}
		if j == 1 && distance > 0 {
			isIncreasing = true
		} else if (!isIncreasing && distance > 0) || (isIncreasing && distance < 0) {
			isUnsafe = true
			break
		}
	}
	return isUnsafe
}
