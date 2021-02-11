package main

import "bufio"
import _ "errors"
import "fmt"
import "os"
import "strings"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data := LoadInput()
	passports := ParsePassports(data)
	valid := CountValidPassports(passports)
	fmt.Printf("Parsed %d passports\n", len(passports))
	fmt.Printf("Valid passports = %d\n", valid)
}

func LoadInput() []string {
	file, err := os.Open("day4.input")
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

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var optionalFields = []string{"cid"}

type Passport struct {
	Data map[string]string
}

func (p *Passport) IsValid() bool {
	for _, field := range requiredFields {
		if len(p.Data[field]) <= 0 {
			return false
		}
	}
	return true
}

func ParsePassports(rawData []string) []Passport {
	var passports []Passport
	data := make(map[string]string)
	for _, line := range rawData {
		if len(line) == 0 {
			passports = append(passports, Passport{data})
			data = make(map[string]string)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				kv := strings.Split(field, ":")
				data[kv[0]] = kv[1]
			}
		}
	}
	passports = append(passports, Passport{data})
	return passports
}

func CountValidPassports(passports []Passport) int {
	count := 0
	for _, passport := range passports {
		if passport.IsValid() {
			count++
		}
	}
	return count
}
