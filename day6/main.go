package main

import "bufio"
import "fmt"
import "os"

func main() {
	rawData := LoadInput()
	groupedData := GroupInput(rawData)
	CountSums(groupedData)
	fmt.Printf("%d groups\n", len(groupedData))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadInput() []string {
	file, err := os.Open("day6.input")
	check(err)
	defer file.Close()
	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		check(err)
		data = append(data, line)
	}
	err = scanner.Err()
	check(err)
	return data
}

func GroupInput(rawData []string) [][]string {
	var allData [][]string
	var data []string
	for _, line := range rawData {
		if len(line) == 0 {
			allData = append(allData, data)
			data = []string{}
		} else {
			data = append(data, line)
		}
	}
	allData = append(allData, data)
	return allData
}

func UniqueCount(group []string) int {
	var m = make(map[rune]int)
	var all []rune
	peopleCount := len(group)
	for _, line := range group {
		for _, char := range line {
			i := m[char]
			m[char] = 1 + i
			if m[char] == peopleCount {
				all = append(all, char)
			}
		}
	}
	fmt.Printf("Group %s has %d\n", group, len(all))
	return len(all)
}

func CountSums(groupedData [][]string) {
	sum := 0
	for _, group := range groupedData {
		sum = sum + UniqueCount(group)
	}
	fmt.Printf("Total Sum: %d\n", sum)
}
