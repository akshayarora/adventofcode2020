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
	Instructions        []string
	Executed            map[int]bool
	Accumulator         int
	Pointer             int
	AlteredInstruction  int
	AlteredInstructions map[int]bool
}

func CreateProgram(rawData []string) Program {
	program := Program{
		Instructions:        rawData,
		Executed:            make(map[int]bool, len(rawData)),
		Accumulator:         0,
		Pointer:             0,
		AlteredInstruction:  0,
		AlteredInstructions: make(map[int]bool, 0)}
	program.Reset()
	return program
}

func (p *Program) Execute() {
	fmt.Println("  ## Executing Program ##")
	for !p.Finished() {
		row := p.Instructions[p.Pointer]
		if p.Executed[p.Pointer] {
			fmt.Printf("Infinite loop detected at %d. Accumulator is %d.\n", p.Pointer, p.Accumulator)
			p.FlipNextInstructionAndRestart()
			break
		}
		instruction, value := ParseInstruction(row)
		fmt.Printf("    [%d] Performing %s / %d\n", p.Pointer, instruction, value)
		p.Executed[p.Pointer] = true
		switch instruction {
		case "nop":
			p.Pointer++
		case "acc":
			p.Accumulator += value
			p.Pointer++
		case "jmp":
			p.Pointer += value
		}
	}
}

func (p *Program) Finished() bool {
	finished := p.Pointer >= len(p.Instructions)
	if finished {
		fmt.Printf("  ## Execution Completed, Accumulator: %d ##\n", p.Accumulator)
	}
	return finished
}

func (p *Program) Reset() {
	// Reset to default state
	p.Accumulator = 0
	p.Pointer = 0
	for i, _ := range p.Instructions {
		p.Executed[i] = false
	}
}

func ParseInstruction(code string) (string, int) {
	instruction := code[0:3]
	value, err := strconv.Atoi(code[4:])
	check(err)
	return instruction, value
}

func CreateCode(instruction string, value int) string {
	return fmt.Sprintf("%s %d", instruction, value)
}

func (p *Program) CanFlipInstruction(pointer int) bool {
	instruction, value := ParseInstruction(p.Instructions[pointer])
	switch instruction {
	case "acc":
		return false
	case "nop":
		if value == 0 {
			return false
		}
	case "jmp":
	}
	return true
}

func (p *Program) FlipInstructionAt(pointer int) bool {
	instruction, value := ParseInstruction(p.Instructions[pointer])
	fmt.Printf("  Altering [%d] %s, %d\n", pointer, instruction, value)
	switch instruction {
	case "nop":
		p.Instructions[pointer] = CreateCode("jmp", value)
	case "jmp":
		p.Instructions[pointer] = CreateCode("nop", value)
	case "acc":
		return false
	}
	return true
}

func (p *Program) FlipNextInstructionAndRestart() {
	p.Reset()
	if p.AlteredInstruction > -1 {
		p.FlipInstructionAt(p.AlteredInstruction)
	}
	for true {
		row := p.Instructions[p.Pointer]
		instruction, value := ParseInstruction(row)
		p.Executed[p.Pointer] = true
		tried, ok := p.AlteredInstructions[p.Pointer]
		if p.CanFlipInstruction(p.Pointer) && (!tried || !ok) {
			fmt.Printf(" ## Flipping [%d] %s\n", p.Pointer, row)
			p.FlipInstructionAt(p.Pointer)
			p.AlteredInstruction = p.Pointer
			p.AlteredInstructions[p.Pointer] = true
			break
		}
		switch instruction {
		case "nop":
			p.Pointer++
		case "acc":
			p.Accumulator += value
			p.Pointer++
		case "jmp":
			p.Pointer += value
		}
	}
	p.Reset()
	p.Execute()
}
