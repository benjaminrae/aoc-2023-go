package day1

import (
	"path"
	"reflect"
	"testing"
)

type parseLineTestCase struct {
	arg1     string
	expected []int
}

var parseLineTestCases = []parseLineTestCase{
	{"1abc2", []int{1, 2}},
	{"pqr3stu8vwx", []int{3, 8}},
	{"a1b2c3d4e5f", []int{1, 2, 3, 4, 5}},
	{"treb7uchet", []int{7}},
}

type combineFirstAndLastTestCase struct {
	arg1     []int
	expected int
}

var combineFirstAndLastCases = []combineFirstAndLastTestCase{
	{[]int{1, 2}, 12},
	{[]int{3, 8}, 38},
	{[]int{1, 2, 3, 4, 5}, 15},
	{[]int{7}, 77},
}

func TestParseLine(t *testing.T) {
	for _, test := range parseLineTestCases {
		if output := ParseLine(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestCombineFirstAndLast(t *testing.T) {
	for _, test := range combineFirstAndLastCases {
		if output, _ := CombineFirstAndLast(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

type testSolveCase struct {
	arg1     string
	expected int
}

var testSolveCases = []testSolveCase{
	{"day1_test.txt", 142},
	{"day1_test2.txt", 281},
}

func TestSolve(t *testing.T) {
	var inputDirectory = "../../input"

	for _, test := range testSolveCases {
		if output := Solve(path.Join(inputDirectory, test.arg1)); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
