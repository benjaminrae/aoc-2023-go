package day1

import (
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

func TestParseLine(t *testing.T) {
	for _, test := range parseLineTestCases {
		if output := ParseLine(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
