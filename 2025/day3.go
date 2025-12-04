package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

func day3part1(inputData string) int {

	defer timer(time.Now(), "day3part1")

	var totalJoltage int
	banks := strings.Fields(inputData)
	for _, bank := range banks {
		var sBank []int
		var maxJoltage int
		// convert to int
		for _, j := range bank {
			joltage, _ := strconv.Atoi(string(j))
			sBank = append(sBank, joltage)
		}
		for i := 0; i < len(sBank)-1; i++ {
			tenth := sBank[i] * 10
			for j := 1; j < len(sBank)-i; j++ {
				joltage := tenth + sBank[i+j]
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}
		totalJoltage += maxJoltage
	}
	return totalJoltage
}

func day3part2(inputData string) int {

	defer timer(time.Now(), "day3part2")

	banks := strings.Fields(strings.TrimSpace(inputData))
	var totalJoltage int

	for _, bank := range banks {
		var sBank []int
		for _, j := range bank {
			joltage, _ := strconv.Atoi(string(j))
			sBank = append(sBank, joltage)
		}
		var output []int

		for len(sBank) > 0 {
			if len(output) == 0 {
				output = append(output, sBank[0])
				sBank = append(sBank[:0], sBank[1:]...)
			}
			for sBank[0] > output[len(output)-1] {
				if len(output)+len(sBank) > 12 {
					output = output[:len(output)-1]
					if len(output) < 1 {
						break
					}
					if output[len(output)-1] > sBank[0] {
						output = append(output, sBank[0])
						sBank = append(sBank[:0], sBank[1:]...)
					}
				} else {
					output = append(output, sBank[0])
					sBank = append(sBank[:0], sBank[1:]...)
					break
				}
				if len(sBank) < 1 {
					break
				}
			}
			if len(sBank) > 0 && len(output) > 0 {
				if len(output) < 12 {
					output = append(output, sBank[0])
					sBank = append(sBank[:0], sBank[1:]...)
				} else {

					if sBank[0] > output[len(output)-1] {
						output[len(output)-1] = sBank[0]
						sBank = append(sBank[:0], sBank[1:]...)
					} else {

						sBank = append(sBank[:0], sBank[1:]...)
					}
				}
			}
		}
		var joltage int
		for i, _ := range output {
			exponent := int(math.Pow10(len(output) - (i + 1)))
			joltage += output[i] * exponent
		}
		totalJoltage += joltage
	}
	return totalJoltage
}
