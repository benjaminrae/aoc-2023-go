package main

import (
	"adventofcode/days/day1"
	"fmt"
	"path"
)

func main() {
	var inputDirectory = "input"

	fmt.Println(day1.Solve(path.Join(inputDirectory, "day1.txt")))
}
