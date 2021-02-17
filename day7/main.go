package main

import "bufio"
import "fmt"
import "os"
import "regexp"
import "strconv"
import "strings"

func main() {
	rawData := LoadInput("day7.input")
	ParseRules(rawData)
	m := make(map[string]Bag)
	color := "shiny gold"
	ParentsForBag(color, &m)
	fmt.Printf("%d bags can contain a %s bag\n", len(m), color)
	count := BagCountWithin(color)
	fmt.Printf("%s bag contains %d other bags\n", color, count)
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

func ParseRules(rawData []string) {
	for _, rule := range rawData {
		ParseRule(rule)
	}
}

func ParseRule(rule string) {
	r := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)
	containerColor := rule[0:strings.Index(rule, " bags")]
	bag := CreateAndGetBag(containerColor)
	matches := r.FindAllStringSubmatch(rule, -1)
	// fmt.Printf("Parsing rule %s\n", rule)
	for _, match := range matches {
		contentRule := ContentRuleFromStringMatch(match)
		*(bag.Contents) = append(*(bag.Contents), contentRule)
		*contentRule.Bag.Parents = append(*contentRule.Bag.Parents, bag)
	}
}

func CreateAndGetBag(color string) Bag {
	if bag, ok := rules[color]; ok {
		return bag
	}
	contents := make([]ContentRule, 0)
	parents := make([]Bag, 0)
	bag := Bag{color, &contents, &parents}
	rules[color] = bag
	return bag
}

func ContentRuleFromStringMatch(rule []string) ContentRule {
	q, err := strconv.Atoi(rule[1])
	check(err)
	bag := CreateAndGetBag(rule[2])
	return ContentRule{&bag, q}
}

var rules = make(map[string]Bag)

type Bag struct {
	Color    string
	Contents *[]ContentRule
	Parents  *[]Bag
}

type ContentRule struct {
	Bag      *Bag
	Quantity int
}

func ParentsForBag(color string, m *map[string]Bag) {
	bag, ok := rules[color]
	if !ok {
		return
	}
	for _, parent := range *bag.Parents {
		(*m)[parent.Color] = parent
		ParentsForBag(parent.Color, m)
	}
}

func BagCountWithin(color string) int {
	count := 0
	_, ok := rules[color]
	if !ok {
		return -1
	}
	return count
}
