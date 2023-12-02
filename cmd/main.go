package main

import (
	"adventofcode/days/day1"
	"adventofcode/days/day2"
	"fmt"
	"path"
)

func main() {
	var inputDirectory = "input"

	fmt.Println(day1.Solve(path.Join(inputDirectory, "day1.txt")))
	fmt.Println(day2.Solve(path.Join(inputDirectory, "day2.txt")))
}
