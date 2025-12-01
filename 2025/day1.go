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
		// Rotating right
		if direction == "R" {
			// Partial right rotation passes zero
			if (position + distance) > 100 {
				countZero += 1
			}
			position = (position + distance) % 100
		}
		// Rotating left
		if direction == "L" {
			// Partial left rotation passes zero
			if distance > position {
				if position != 0 {
					countZero += 1
				}
				position = (position - distance) + 100
			} else {
				position = (position - distance)
			}
		}
		// Count full rotations
		countZero += fullRotations
		// Finally count zero if turning the dial lands on zero
		if position == 0 {
			countZero++
		}
	}
	return countZero
}
