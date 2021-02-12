package main

import "bufio"
import "fmt"
import "os"
import "sort"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadInput() []string {
	file, err := os.Open("day5.input")
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

func main() {
	rawData := LoadInput()
	// ValidateKnownInputs()
	// ComputeHighestSeatId(rawData)
	ComputeMySeatId(rawData)
}

func ComputeHighestSeatId(rawData []string) {
	highSeatId := 0
	for _, code := range rawData {
		row, col := ParseBoardingPass(code)
		seatId := ComputeSeatId(row, col)
		if seatId > highSeatId {
			highSeatId = seatId
		}
	}
	fmt.Printf("High Seat Id %d\n", highSeatId)
}

func ComputeMySeatId(rawData []string) {
	var allSeats []int
	for _, code := range rawData {
		row, col := ParseBoardingPass(code)
		seatId := ComputeSeatId(row, col)
		allSeats = append(allSeats, seatId)
	}
	sort.Ints(allSeats)
	for i, v := range allSeats {
		if v+2 == allSeats[i+1] {
			fmt.Printf("My SeatID is %d", v+1)
			break
		}
	}
}

func ValidateKnownInputs() {
	// Validate via input given
	code := "FBFBBFFRLR"
	row, col := ParseBoardingPass(code)
	seatId := ComputeSeatId(row, col)
	fmt.Printf("%s = %d,%d = %d\n", code, row, col, seatId)

	code = "BFFFBBFRRR"
	row, col = ParseBoardingPass(code)
	seatId = ComputeSeatId(row, col)
	fmt.Printf("%s = %d,%d = %d\n", code, row, col, seatId)

	code = "FFFBBBFRRR"
	row, col = ParseBoardingPass(code)
	seatId = ComputeSeatId(row, col)
	fmt.Printf("%s = %d,%d = %d\n", code, row, col, seatId)

	code = "BBFFBBFRLL"
	row, col = ParseBoardingPass(code)
	seatId = ComputeSeatId(row, col)
	fmt.Printf("%s = %d,%d = %d\n", code, row, col, seatId)
}

func ParseBoardingPass(code string) (int, int) {
	row := ParseCode(0, 127, code[0:7])
	col := ParseCode(0, 7, code[7:])
	return row, col
}

func ParseCode(min, max int, code string) int {
	if max-min == 1 {
		if CodeIsHigh(code[0]) {
			return max
		} else {
			return min
		}
	}
	mid := (max-min+1)/2 + min
	if CodeIsHigh(code[0]) {
		return ParseCode(mid, max, code[1:])
	} else {
		return ParseCode(min, mid-1, code[1:])
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

func ComputeSeatId(row, col int) int {
	return row*8 + col
}

// func ComputeRowColFromSeatId(seatId int) (int, int){
