package main

import (
	"fmt"
	"strings"
	"time"
)

func day4part1(inputData string) {

	defer timer(time.Now(), "day4part1")

	var data [][]string
	rows := strings.Split(strings.TrimSpace(inputData), "\n")

	for _, d := range rows {
		data = append(data, strings.Split(d, ""))
	}

	numRows := len(data) - 1
	numCols := len(data[0]) - 1

	var accessibleRolls int

	for row := 0; row <= numRows; row++ {
		for col := 0; col <= numCols; col++ {
			var adjacentRolls int
			if (col == 0 && row == 0) ||
				(col == numCols && row == numRows) ||
				(col == 0 && row == numRows) ||
				(col == numCols && row == 0) {

				if data[row][col] == "@" {
					accessibleRolls++
				}
				continue
			}

			var testGrid []string

			if data[row][col] == "@" {
				if row == 0 {
					testGrid = []string{
						data[row][col-1],
						data[row][col+1],
						data[row+1][col-1],
						data[row+1][col],
						data[row+1][col+1],
					}
				} else if col == 0 {
					testGrid = []string{
						data[row-1][col],
						data[row+1][col],
						data[row-1][col+1],
						data[row][col+1],
						data[row+1][col+1],
					}
				} else if row == numRows {
					testGrid = []string{
						data[row][col-1],
						data[row][col+1],
						data[row-1][col-1],
						data[row-1][col],
						data[row-1][col+1],
					}

				} else if col == numCols {
					testGrid = []string{
						data[row-1][col],
						data[row+1][col],
						data[row-1][col-1],
						data[row][col-1],
						data[row+1][col-1],
					}
				} else {
					testGrid = []string{
						data[row-1][col-1],
						data[row-1][col],
						data[row-1][col+1],
						data[row][col-1],
						data[row][col+1],
						data[row+1][col-1],
						data[row+1][col],
						data[row+1][col+1],
					}
				}
				for _, t := range testGrid {
					if t == "@" {
						adjacentRolls++
					}
				}

				if adjacentRolls < 4 {
					accessibleRolls++
				}
			}
		}
	}
	fmt.Println(accessibleRolls)
}

func day4part2(inputData string) {

	defer timer(time.Now(), "day4part2")

	var data [][]string
	rows := strings.Split(strings.TrimSpace(inputData), "\n")

	for _, d := range rows {
		data = append(data, strings.Split(d, ""))
	}

	numRows := len(data) - 1
	numCols := len(data[0]) - 1

	var accessedRolls int
	var accessedOnIteration int

	isAccessibleRolls := true

	for isAccessibleRolls {
		accessedOnIteration = 0
		for row := 0; row <= numRows; row++ {
			for col := 0; col <= numCols; col++ {

				var adjacentRolls int

				if (col == 0 && row == 0) ||
					(col == numCols && row == numRows) ||
					(col == 0 && row == numRows) ||
					(col == numCols && row == 0) {

					if data[row][col] == "@" {
						accessedRolls++
						accessedOnIteration++
						data[row][col] = "."
					}
					continue
				}

				var testGrid []string

				if data[row][col] == "@" {
					if row == 0 {
						testGrid = []string{
							data[row][col-1],
							data[row][col+1],
							data[row+1][col-1],
							data[row+1][col],
							data[row+1][col+1],
						}
					} else if col == 0 {
						testGrid = []string{
							data[row-1][col],
							data[row+1][col],
							data[row-1][col+1],
							data[row][col+1],
							data[row+1][col+1],
						}
					} else if row == numRows {
						testGrid = []string{
							data[row][col-1],
							data[row][col+1],
							data[row-1][col-1],
							data[row-1][col],
							data[row-1][col+1],
						}

					} else if col == numCols {
						testGrid = []string{
							data[row-1][col],
							data[row+1][col],
							data[row-1][col-1],
							data[row][col-1],
							data[row+1][col-1],
						}
					} else {
						testGrid = []string{
							data[row-1][col-1],
							data[row-1][col],
							data[row-1][col+1],
							data[row][col-1],
							data[row][col+1],
							data[row+1][col-1],
							data[row+1][col],
							data[row+1][col+1],
						}
					}
					for _, t := range testGrid {
						if t == "@" {
							adjacentRolls++
						}
					}
					if adjacentRolls < 4 {
						accessedRolls++
						accessedOnIteration++
						data[row][col] = "."
					}
				}
			}
		}
		if accessedOnIteration == 0 {
			isAccessibleRolls = false
		}
	}
	fmt.Println(accessedRolls)
}
