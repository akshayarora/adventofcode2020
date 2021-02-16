package main

import "bufio"
import "fmt"
import "os"

func main() {
	// rawData := LoadInput("day6.input")
	// groupedData := GroupInput(rawData)
	fmt.Println("done")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadInput(fileName string) []string {
	file, err := os.Open(fileName)
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
