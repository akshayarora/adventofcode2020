package main

import "bufio"
import _ "errors"
import "fmt"
import "os"
import "regexp"
import "strconv"
import "strings"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	passwords := LoadInput()
	//Part1(passwords)
	Validate(passwords)
}

func Validate(passwords []string) {
	count1 := 0
	count2 := 0
	r := regexp.MustCompile(`^(\d+)-(\d+)\s([a-z]):\s(.*)$`)
	for _, line := range passwords {
		components := r.FindStringSubmatch(line)
		if len(components) != 5 {
			fmt.Printf("Line with value `%s` did not find a match\n", line)
			continue
		}
		min, _ := strconv.Atoi(components[1])
		max, _ := strconv.Atoi(components[2])
		letter := components[3]
		password := components[4]
		// Part 1
		if IsValidPart1(min, max, letter, password) {
			count1++
		}
		// Part 2
		if IsValidPart2(min, max, letter, password) {
			count2++
		}
	}
	if count1 != 454 {
		panic(fmt.Sprintf("Count for part 1 is wrong. Expected 454, received %d", count1))
	}
	if count2 != 649 {
		panic(fmt.Sprintf("Count for part 2 is wrong. Expected 649, received %d", count2))
	}
	fmt.Printf("%d passwords are valid in Part 1\n", count1)
	fmt.Printf("%d passwords are valid in Part 2\n", count2)
}

func IsValidPart1(min, max int, letter, password string) bool {
	n := strings.Count(password, letter)
	return n >= min && n <= max
}

func IsValidPart2(min, max int, letter, password string) bool {
	if min > len(password) || max > len(password) {
		return false
	}
	a := password[min-1] == letter[0]
	b := password[max-1] == letter[0]
	return (a || b) && !(a && b)
}

func LoadInput() []string {
	file, err := os.Open("day2.input")
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
