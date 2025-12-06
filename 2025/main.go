package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

func readInputFile(inputFile string) (string, error) {

	_, err := os.Stat(inputFile)
	if err != nil {
		return "", err
	}

	inputData, err := os.ReadFile(inputFile)

	return string(inputData), err
}

func writeInputFile(inputFile string, inputDir string, inputData []byte) {
	_ = os.MkdirAll(inputDir, os.ModePerm)
	_ = os.WriteFile(inputFile, inputData, os.ModePerm)
}

func getInput(day int) []byte {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file")
	}

	sessionCookie := os.Getenv("SESSION_COOKIE")

	if sessionCookie == "" {
		log.Fatal("SESSION_COOKIE environment variable not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", day)

	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie})

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func timer(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {

	inputDir := "inputs"
	var data []string

	for day := 1; day <= 5; day++ {

		inputFile := filepath.Join(inputDir, fmt.Sprintf("day%d.txt", day))

		inputData, err := readInputFile(inputFile)
		if err != nil {
			inputDataRaw := getInput(day)
			writeInputFile(inputFile, inputDir, inputDataRaw)
		}
		inputData, err = readInputFile(inputFile)
		data = append(data, inputData)
	}

	//day1part1(data[0])
	//day1part2(data[0])
	//day2part1(data[1])
	//day2part2(data[1])
	//day3part1(data[2])
	//day3part2(data[2])
	//day4part1(data[3])
	//day4part2(data[3])
	//day5part1(data[4])
	day5part2(data[4])
	//day6part1(data[5])
	//day6part2(data[5])
	//day7part1(data[6])
	//day7part2(data[6])
	//day8part1(data[7])
	//day8part2(data[7])
	//day9part1(data[8])
	//day9part2(data[8])
	//day10part1(data[9])
	//day10part2(data[9])
	//day11part1(data[10])
	//day11part2(data[10])
	//day12part1(data[11])
	//day12part2(data[11])
}
