package main

import "fmt"

func main() {
	fmt.Printf("Code is %d\n", Parse(0, 127, "FBFBBFF"))
	fmt.Printf("Code is %d\n", Parse(0, 7, "RLR"))
}

func Parse(min, max int, code string) int {
	if max-min == 1 {
		if CodeIsHigh(code[0]) {
			return max
		} else {
			return min
		}
	}
	mid := (max-min+1)/2 + min
	if CodeIsHigh(code[0]) {
		return Parse(mid, max, code[1:])
	} else {
		return Parse(min, mid-1, code[1:])
	}
	return -1
}

var Front = byte('F')
var Back = byte('B')
var Right = byte('R')
var Left = byte('L')

func CodeIsHigh(b byte) bool {
	return b == Back || b == Right
}

func CodeIsLow(b byte) bool {
	return b == Front || b == Left
}
