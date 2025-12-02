package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

func main() {

	inputDir := "inputs"
	var data []string

	for day := 1; day <= 2; day++ {

		inputFile := filepath.Join(inputDir, fmt.Sprintf("day%d.txt", day))

		inputData, err := readInputFile(inputFile)
		if err != nil {
			inputDataRaw := getInput(day)
			writeInputFile(inputFile, inputDir, inputDataRaw)
		}
		inputData, err = readInputFile(inputFile)
		data = append(data, inputData)
	}

	fmt.Println(day1part1(data[0]))
	fmt.Println(day1part2(data[0]))
	fmt.Println(day2part1(data[1]))
	fmt.Println(day2part2(data[1]))
}
