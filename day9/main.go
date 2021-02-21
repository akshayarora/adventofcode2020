package main

import "bufio"
import "fmt"
import "os"
import "strconv"

var part1 = 731031916

func main() {
	rawData := LoadInput("day9.input")
	data := ConvertToInt(rawData)
	weakness := FindWeakness(25, data)
	if weakness != part1 {
		fmt.Printf("Weakness expected %d, found %d\n", part1, weakness)
		return
	}
	fmt.Printf("Number %d weakness\n", weakness)
	encWeakness := FindEncryptionWeakness(data, weakness)
	fmt.Printf("Encryption Weakness = %d\n", encWeakness)
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

func FindWeakness(preamble int, data []int) int {
	for i := preamble; i < len(data); i++ {
		j := data[i]
		slice := data[i-preamble : i]
		k, l, ok := FindComponentsForSum(slice, j)
		if !ok {
			return j
		} else {
			fmt.Printf("[%d/%d] %d = %d + %d\n", i, len(data), j, k, l)
		}
	}
	return -1
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

func FindEncryptionWeakness(data []int, weakness int) int {
	for i := 0; i < len(data)-1; i++ {
		min := data[i]
		max := data[i]
		sum := data[i]
		for j := i + 1; j < len(data); j++ {
			sum += data[j]
			if min > data[j] {
				min = data[j]
			}
			if max < data[j] {
				max = data[j]
			}
			if weakness == sum {
				return min + max
			}
		}
	}
	return -1
}
