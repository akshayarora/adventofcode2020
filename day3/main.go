package main

import "bufio"
import _ "errors"
import "fmt"
import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	treemap := LoadInput()
	treeMult := Traverse(Point{0, 0}, Point{1, 1}, treemap) *
		Traverse(Point{0, 0}, Point{3, 1}, treemap) *
		Traverse(Point{0, 0}, Point{5, 1}, treemap) *
		Traverse(Point{0, 0}, Point{7, 1}, treemap) *
		Traverse(Point{0, 0}, Point{1, 2}, treemap)
	fmt.Printf("Tree Mult = %d\n", treeMult)
}

func Traverse(origin, path Point, treemap []string) int {
	fmt.Printf("Starting at %o, following path %o\n", origin, path)
	mapWidth := len(treemap[0])
	mapHeight := len(treemap)
	current := origin.Add(&path)
	treesEncountered := 0
	tree := ([]byte("#"))[0]
	for current.Y < mapHeight {
		// fmt.Printf("%d < %d\n", current.Y, mapHeight)
		node := treemap[current.Y][current.X%mapWidth]
		// fmt.Printf("Now at %d,%d, which is %s\n", current.X, current.Y, string(node))
		if node == tree {
			treesEncountered++
		}
		current = current.Add(&path)
	}
	fmt.Printf("Total trees encountered %d\n", treesEncountered)
	return treesEncountered
}

func LoadInput() []string {
	file, err := os.Open("day3.input")
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

type Point struct {
	X int
	Y int
}

func (a *Point) Add(b *Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}
