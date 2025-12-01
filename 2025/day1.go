package main

import (
	"math"
	"strconv"
	"strings"
)

func day1part1(inputData string) int {
	var countZero int
	position := 50
	rotations := strings.Fields(inputData)

	for _, rot := range rotations {
		direction := rot[:1]
		distance, _ := strconv.Atoi(rot[1:])
		distance = distance % 100

		if direction == "R" {
			position = (position + distance) % 100
		} else if distance > position {
			position = (position - distance) + 100
		} else {
			position = (position - distance)
		}

		if position == 0 {
			countZero++
		}
	}
	return countZero
}

func day1part2(inputData string) int {
	var countZero int
	position := 50
	rotations := strings.Fields(inputData)

	for _, rot := range rotations {
		direction := rot[:1]
		distance, _ := strconv.Atoi(rot[1:])
		fullRotations := int(math.Abs(math.Floor(float64(distance / 100))))
		distance = distance % 100

		// All full rotations pass zero
		countZero += fullRotations

		if direction == "R" {
			// Partial right rotation passing zero
			if (position + distance) > 100 {
				countZero++
			}
			position = (position + distance) % 100
		}

		if direction == "L" {
			// Partial left rotation passing zero
			if distance > position {
				if position != 0 {
					countZero++
				}
				position = (position - distance) + 100
			} else {
				position = (position - distance)
			}
		}

		// Landing on zero
		if position == 0 {
			countZero++
		}
	}
	return countZero
}
