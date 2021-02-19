package main

import "bufio"
import "fmt"
import "os"
import "strconv"

func main() {
	rawData := LoadInput("day8.input")
	program := CreateProgram(rawData)
	fmt.Printf("Program created with %d lines of code\n", len(program.Instructions))
	program.Execute()
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

type Program struct {
	Instructions []string
	Executed     map[int]bool
	Accumulator  int
}

func CreateProgram(rawData []string) Program {
	program := Program{rawData, make(map[int]bool, len(rawData)), 0}
	for i, _ := range rawData {
		program.Executed[i] = false
	}
	return program
}

func (p *Program) Execute() {
	fmt.Println("  ## Executing Program ##")
	pointer := 0

	for true {
		row := p.Instructions[pointer]
		if p.Executed[pointer] {
			fmt.Printf("Infinite loop detected. Accumulator is %d\n", p.Accumulator)
			break
		}
		instruction := row[0:3]
		input, err := strconv.Atoi(row[4:])
		check(err)
		fmt.Printf("    [%d] Performing %s / %d\n", pointer, instruction, input)
		p.Executed[pointer] = true
		switch instruction {
		case "nop":
			pointer++
		case "acc":
			p.Accumulator += input
			pointer++
		case "jmp":
			pointer += input
		}
	}
}
