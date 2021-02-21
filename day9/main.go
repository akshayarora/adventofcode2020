package main

import "bufio"
import "fmt"
import "os"
import "strconv"

func main() {
	rawData := LoadInput("day9.input")
	data := ConvertToInt(rawData)
	FindWeakness(25, data)
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

func ConvertToInt(input []string) []int {
	var output = make([]int, len(input))
	for i, s := range input {
		j, err := strconv.Atoi(s)
		check(err)
		output[i] = j
	}
	return output
}

func FindWeakness(preamble int, data []int) {
	for i := preamble; i < len(data); i++ {
		j := data[i]
		slice := data[i-preamble : i]
		k, l, ok := FindComponentsForSum(slice, j)
		if !ok {
			fmt.Printf("Number %d doesn't match\n", j)
			break
		} else {
			fmt.Printf("[%d/%d] %d = %d + %d\n", i, len(data), j, k, l)
		}
	}
}

func FindComponentsForSum(data []int, sum int) (int, int, bool) {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == sum {
				return data[i], data[j], true
			}
		}
	}
	return 0, 0, false
}
