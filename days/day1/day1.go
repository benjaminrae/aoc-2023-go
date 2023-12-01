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

func CombineFirstAndLast(numbers []int) (combined int, err error) {
	var first = numbers[0]
	var last = numbers[len(numbers)-1]

	var joined = fmt.Sprintf("%v%v", first, last)

	num, err := strconv.Atoi(joined)

	if err != nil {
		return 0, err
	}

	return num, nil
}

func Solve() {
	var line1 = "1abc2"
	var line2 = "pqr3stu8vwx"
	var line3 = "a1b2c3d4e5f"
	var line4 = "treb7uchet"

	fmt.Println(CombineFirstAndLast(ParseLine(line1)))
	fmt.Println(CombineFirstAndLast(ParseLine(line2)))
	fmt.Println(CombineFirstAndLast(ParseLine(line3)))
	fmt.Println(CombineFirstAndLast(ParseLine(line4)))

}
