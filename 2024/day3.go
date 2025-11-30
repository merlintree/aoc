package main

import (
	"regexp"
	"strconv"
)

func day3part1(inputData string) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	b := []byte(inputData)

	matches := re.FindAllIndex(b, -1)

	first_n_re := regexp.MustCompile(`\(\d{1,3}`)
	second_n_re := regexp.MustCompile(`\d{1,3}\)`)

	var output int

	for _, loc := range matches {
		all_match := b[loc[0]:loc[1]]

		first_n_loc := first_n_re.FindIndex(all_match)
		first_n, _ := strconv.Atoi(string(all_match[first_n_loc[0]+1 : first_n_loc[1]]))

		second_n_loc := second_n_re.FindIndex(all_match)
		second_n, _ := strconv.Atoi(string(all_match[second_n_loc[0] : second_n_loc[1]-1]))

		output = output + first_n*second_n

	}

	return output
}
