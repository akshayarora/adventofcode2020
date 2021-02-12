package main

import "bufio"
import _ "errors"
import "fmt"
import "os"
import "regexp"
import "strings"
import "strconv"

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
var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

type Passport struct {
	Data map[string]Field
}
type Field struct {
	Key   string
	Value string
}

func (p *Passport) IsValid() bool {
	for _, key := range requiredFields {
		field := p.Data[key]
		if !field.IsValid() {
			return false
		}
	}
	return true
}

func (f *Field) IsValid() bool {
	switch key := f.Key; key {
	case "byr":
		i, _ := strconv.Atoi(f.Value)
		return i >= 1920 && i <= 2002
	case "iyr":
		i, _ := strconv.Atoi(f.Value)
		return i >= 2010 && i <= 2020
	case "eyr":
		i, _ := strconv.Atoi(f.Value)
		return i >= 2020 && i <= 2030
	case "hgt":
		r := regexp.MustCompile(`^(\d+)(cm|in)$`)
		components := r.FindStringSubmatch(f.Value)
		if len(components) != 3 {
			return false
		}
		i, _ := strconv.Atoi(components[1])
		switch components[2] {
		case "cm":
			return i >= 150 && i <= 193
		case "in":
			return i >= 59 && i <= 76
		}
		return false
	case "hcl":
		r := regexp.MustCompile(`^(#[0-9a-f]{6})$`)
		components := r.FindStringSubmatch(f.Value)
		if len(components) != 2 {
			return false
		}
		return true
	case "ecl":
		return stringInSlice(f.Value, eyeColors)
	case "pid":
		b, err := regexp.MatchString(`^[0-9]{9}$`, f.Value)
		check(err)
		return b
	case "cid":
		return true
	}
	return false
}

func ParsePassports(rawData []string) []Passport {
	var passports []Passport
	data := make(map[string]Field)
	for _, line := range rawData {
		if len(line) == 0 {
			passports = append(passports, Passport{data})
			data = make(map[string]Field)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				kv := strings.Split(field, ":")
				data[kv[0]] = Field{kv[0], kv[1]}
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
