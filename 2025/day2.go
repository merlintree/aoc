package main

import (
	"strconv"
	"strings"
)

func day2part1(inputData string) int {
	data := strings.Split(strings.TrimSpace(inputData), ",")
	var sumInvalid int

	for _, r := range data {
		bounds := strings.Split(r, "-")
		lowerBound, _ := strconv.Atoi(bounds[0])
		upperBound, _ := strconv.Atoi(bounds[1])

		for i := lowerBound; i <= upperBound; i++ {
			s := strconv.Itoa(i)
			// Single digits cannot have repeating patterns
			if len(s) == 1 {
				continue
			}
			// Only Ids with an even number of digits have two repeating patterns
			if len(s)%2 == 0 {
				right := s[len(s)/2:]
				left := s[:len(s)/2]
				if right == left {
					sumInvalid += i
				}
			}
		}
	}
	return sumInvalid
}

func day2part2(inputData string) int {
	data := strings.Split(strings.TrimSpace(inputData), ",")
	var sumInvalid int

	for _, r := range data {
		bounds := strings.Split(r, "-")
		lowerBound, _ := strconv.Atoi(bounds[0])
		upperBound, _ := strconv.Atoi(bounds[1])

		for i := lowerBound; i <= upperBound; i++ {
			s := strconv.Itoa(i)
			// Increase element size until > len(s)/2
			for size := 1; size < len(s); size++ {
				var relSize int
				elements := []string{}

				if len(s)%size != 0 {
					continue
				}
				for pos := 0; pos < len(s); pos += size {
					relSize += size
					if relSize > len(s) {
						relSize = len(s)
					}
					elements = append(elements, s[pos:relSize])
				}
				invalid := true
				// Check if all elements are the same
				for e := range elements {
					if elements[e] != elements[0] {
						invalid = false
						break
					}
				}
				if invalid {
					sumInvalid += i
					break
				}
			}
		}
	}
	return sumInvalid
}
