package day1

import (
	"fmt"
	"strconv"
)

func ParseLine(line string) []int {
	var numbers = []int{}

	for _, char := range line {
		num, err := strconv.Atoi(string(char))

		if err != nil {
			continue
		}

		numbers = append(numbers, num)
	}

	return numbers
}

func Solve() {
	var line1 = "1abc2"
	var line2 = "pqr3stu8vwx"
	var line3 = "a1b2c3d4e5f"
	var line4 = "treb7uchet"

	fmt.Println(ParseLine(line1))
	fmt.Println(ParseLine(line2))
	fmt.Println(ParseLine(line3))
	fmt.Println(ParseLine(line4))
}
