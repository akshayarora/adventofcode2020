package main

import "bufio"
import "errors"
import "fmt"
import "os"
import "strconv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data := LoadInput()
	validate(data)
}

func validate(data []int) {
	for2 := 928896
	for3 := 295668576
	a, b := FindTuple(data, 2020)
	c, d, e := FindTriple(data, 2020)
	if a*b != for2 {
		panic(fmt.Sprintf("Tuple for2 check failed, expected %d, got %d", for2, a*b))
	}
	fmt.Printf("%d * %d = %d\n", a, b, for2)
	// 1312, 708
	if c*d*e != for3 {
		panic("Triple for3 check failed")
	}
	fmt.Printf("%d * %d * %d  = %d\n", c, d, e, for3)
	// 798, 664, 558
	numbers, err := FindNumbersForSum(data, 0, 2, 2020)
	check(err)
	fmt.Println(numbers)
	a = numbers[0]
	b = numbers[1]
	if a*b != for2 {
		panic(fmt.Sprintf("Tuple for2 check failed, expected %d, got %d", for2, a*b))
	}
	numbers, err = FindNumbersForSum(data, 0, 3, 2020)
	c = numbers[0]
	d = numbers[1]
	e = numbers[2]
	check(err)
	fmt.Println(numbers)
	if c*d*e != for3 {
		panic("Triple for3 check failed")
	}
}

func LoadInput() []int {
	file, err := os.Open("day1.input")
	check(err)
	defer file.Close()
	var array []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		check(err)
		array = append(array, i)
	}
	err = scanner.Err()
	check(err)
	return array
}

func FindNumbersForSum(data []int, startingIndex int, n int, sum int) ([]int, error) {
	var result []int
	for i := startingIndex; i < len(data)-(n-1); i++ {
		num := data[i]
		//fmt.Printf("For %d, trying with %d...\n", sum, num)
		if n == 1 {
			if num == sum {
				//fmt.Printf("   appending %d\n", num)
				result = append(result, num)
				//fmt.Println(result)
				return result, nil
			}
		} else {
			nums, err := FindNumbersForSum(data, i+1, n-1, sum-num)
			if err == nil {
				result = append(result, num)
				result = append(result, nums...)
				//fmt.Println(result)
				return result, nil
			}
		}
	}
	return nil, errors.New("Matching numbers not found")
}

func FindTuple(data []int, total int) (int, int) {
	var num1, num2 int
	for i := 0; i < len(data)-1; i++ {
		num1 = int(data[i])
		for j := i + 1; j < len(data); j++ {
			num2 = int(data[j])
			if num1+num2 == 2020 {
				return num1, num2
			}
		}
	}
	return 0, 0
}

func FindTriple(data []int, total int) (int, int, int) {
	var num1, num2, num3 int
	for i := 0; i < len(data)-2; i++ {
		num1 = int(data[i])
		for j := i + 1; j < len(data)-1; j++ {
			num2 = int(data[j])
			for k := j + 1; k < len(data); k++ {
				num3 = int(data[k])
				if num1+num2+num3 == 2020 {
					return num1, num2, num3
				}
			}
		}
	}
	return 0, 0, 0
}
