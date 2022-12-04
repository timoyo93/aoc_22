package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	res := getInputData("input2")
	result := solution(res)
	for _, v := range result {
		fmt.Printf("%d, %d\n", v.Key, v.Value)
	}
	fmt.Println("-------------------------")
	secondResult := solutionSecondPart(result)
	fmt.Printf("Total calories: %d", secondResult)
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
	fileLines = append(fileLines, "")
	_ = readFile.Close()
	return fileLines
}

func solution(data []string) PairList {
	calories := 0
	elfCount := 0
	dict := make(map[int]int)
	for _, line := range data {
		if line == "" {
			dict[elfCount] = calories
			elfCount++
			calories = 0
		}
		if line != "" {
			c, _ := strconv.Atoi(line)
			calories = calories + c
		}
	}
	return sortDataByValue(dict)
}

func solutionSecondPart(list PairList) int {
	p := list[list.Len()-3 : list.Len()]
	sum := 0
	for _, v := range p {
		sum = sum + v.Value
	}
	return sum
}

func sortDataByValue(sortdata map[int]int) PairList {
	p := make(PairList, len(sortdata))
	i := 0
	for k, v := range sortdata {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)

	return p
}
