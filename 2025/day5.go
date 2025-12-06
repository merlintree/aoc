package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day5part1(inputData string) {

	data := strings.Split(inputData, "\n\n")

	rawRanges := strings.Split(strings.TrimSpace(data[0]), "\n")
	var ranges [][]int

	for _, r := range rawRanges {
		strRange := strings.Split(r, "-")
		lowerRange, _ := strconv.Atoi(strRange[0])
		upperRange, _ := strconv.Atoi(strRange[1])
		ranges = append(ranges, []int{lowerRange, upperRange})
	}

	rawIds := strings.Split(strings.TrimSpace(data[1]), "\n")
	var ids []int
	for _, r := range rawIds {
		id, _ := strconv.Atoi(r)
		ids = append(ids, id)
	}

	var countFresh int

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				countFresh++
				break
			}
		}
	}

	fmt.Println(countFresh)

}

func day5part2(inputData string) {
	data := strings.Split(inputData, "\n\n")

	rawRanges := strings.Split(strings.TrimSpace(data[0]), "\n")
	var ranges [][]int

	for _, r := range rawRanges {
		strRange := strings.Split(r, "-")
		lowerRange, _ := strconv.Atoi(strRange[0])
		upperRange, _ := strconv.Atoi(strRange[1])
		ranges = append(ranges, []int{lowerRange, upperRange})
	}

	sort.Slice(ranges, func(a, b int) bool {
		return ranges[a][0] < ranges[b][0]
	})

	mergedRanges := [][]int{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		last := mergedRanges[len(mergedRanges)-1]
		curr := ranges[i]

		if curr[0] <= last[1] {
			mergedRanges[len(mergedRanges)-1][1] = max(curr[1], last[1])
		} else {
			mergedRanges = append(mergedRanges, ranges[i])
		}
	}

	var countIds int

	for _, r := range mergedRanges {
		countIds += (r[1] - r[0]) + 1
	}

	fmt.Println(countIds)
}
