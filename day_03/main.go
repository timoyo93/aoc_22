package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result := solution(getInputData("input"), CreatePriorities())
	fmt.Println(result)
    secondResult := solutionPartTwo(getInputData("input"), CreatePriorities())
    fmt.Println("--------------------------")
    fmt.Println(secondResult)
}

func getInputData(filename string) []string {
	readFile, err := os.Open(fmt.Sprintf("%s.txt", filename))
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	err = readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	return fileLines
}

func solution(data []string, priorities map[string]int) int {
	score := 0
	for _, line := range data {
		first, second := getFirstAndSecondHalf(line)
		item := findItemInBothCompartments(first, second)
		val, ok := priorities[item]
		if !ok {
			fmt.Println("Item not found in priorities")
		}
		score = score + val
	}
	return score
}

func solutionPartTwo(data []string, priorities map[string]int) int {
	score := 0
	index := len(data) - 2
	for i := 0; i < index; i = i + 3 {
		first := data[i]
		second := data[i+1]
		third := data[i+2]
		item := findItemInElfGroup(first, second, third)
		val, _ := priorities[item]
		score = score + val

	}
	return score
}

func findItemInElfGroup(first, second, third string) string {
	var foundItem string
	for _, f := range first {
		for _, s := range second {
			if string(f) == string(s) {
				for _, t := range third {
					if string(s) == string(t) {
						foundItem = string(t)
						break
					}
				}
				break
			}
		}
	}
	return foundItem
}

func findItemInBothCompartments(first, second string) string {
	var foundItem string
	for _, f := range first {
		for _, s := range second {
			if string(f) == string(s) {
				foundItem = string(s)
				break
			}
		}
	}
	return foundItem
}

func getFirstAndSecondHalf(line string) (first, second string) {
	index := len(line) / 2
	firstHalf := line[0:index]
	secondHalf := line[index:]
	return firstHalf, secondHalf
}
