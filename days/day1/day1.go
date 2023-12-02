package day1

import (
	"adventofcode/pkg"
	"fmt"
	"log"
	"strconv"
	"strings"
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

type Positions struct {
	First int
	Last  int
}

type Lookup map[string]Positions

func GetLookup(line string) Lookup {
	var lookup = Lookup{
		"one":   {-1, -1},
		"two":   {-1, -1},
		"three": {-1, -1},
		"four":  {-1, -1},
		"five":  {-1, -1},
		"six":   {-1, -1},
		"seven": {-1, -1},
		"eight": {-1, -1},
		"nine":  {-1, -1},
		"1":     {-1, -1},
		"2":     {-1, -1},
		"3":     {-1, -1},
		"4":     {-1, -1},
		"5":     {-1, -1},
		"6":     {-1, -1},
		"7":     {-1, -1},
		"8":     {-1, -1},
		"9":     {-1, -1},
	}

	for digit, _ := range lookup {
		var first = strings.Index(line, digit)
		var last = strings.LastIndex(line, digit)

		lookup[digit] = Positions{first, last}
	}

	return lookup
}

func TextToInt(text string) (num int, err error) {
	var textDigitMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9}

	num, err = strconv.Atoi(text)

	if err != nil {
		if val, ok := textDigitMap[text]; ok {
			return val, nil
		}

		return 0, err
	}

	return num, nil
}

func GetFirstAndLastFromLookup(lookup Lookup) (first string, last string) {
	var firstPosition = -1
	var lastPosition = -1
	var firstDigit = ""
	var lastDigit = ""

	for digit, positions := range lookup {
		if firstPosition == -1 && positions.First != -1 {
			firstPosition = positions.First
			first = digit
		}

		if positions.First != -1 && positions.First < firstPosition {
			firstPosition = positions.First
			firstDigit = digit
		}

		if lastPosition == -1 && positions.Last != -1 {
			lastPosition = positions.Last
			last = digit
		}

		if positions.Last != -1 && positions.Last > lastPosition {
			lastPosition = positions.Last
			lastDigit = digit
		}

	}

	return firstDigit, lastDigit
}

func GetFirstFromLookup(lookup Lookup) string {
	var firstPosition = -1
	var firstDigit = ""

	for digit, positions := range lookup {
		if firstPosition == -1 && positions.First != -1 {
			firstPosition = positions.First
			firstDigit = digit
		}

		if positions.First != -1 && positions.First < firstPosition {
			firstPosition = positions.First
			firstDigit = digit
		}
	}

	return firstDigit
}

func GetLastFromLookup(lookup Lookup) string {
	var lastPosition = -1
	var lastDigit = ""

	for digit, positions := range lookup {
		if lastPosition == -1 && positions.Last != -1 {
			lastPosition = positions.Last
			lastDigit = digit
		}

		if positions.Last != -1 && positions.Last > lastPosition {
			lastPosition = positions.Last
			lastDigit = digit
		}
	}

	return lastDigit
}

func Solve(inputFile string) int {
	var input = pkg.GetFileContents(inputFile)

	var lines = strings.Split(input, "\n")

	var lineResults = []int{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var lookup = GetLookup(line)

		var first = GetFirstFromLookup(lookup)
		var last = GetLastFromLookup(lookup)

		var firstNum, err = TextToInt(first)

		if err != nil {
			log.Fatal(err)
			continue
		}

		var lastNum, err2 = TextToInt(last)

		if err2 != nil {
			log.Fatal(err2)
			continue
		}

		var combined, err3 = CombineFirstAndLast([]int{firstNum, lastNum})

		if err3 != nil {
			log.Fatal(err3)
			continue
		}

		lineResults = append(lineResults, combined)

	}

	var sum = 0

	for _, num := range lineResults {
		sum += num
	}

	return sum
}
