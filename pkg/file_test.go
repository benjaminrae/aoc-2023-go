package pkg

import "testing"

func TestGetInputFileContents(t *testing.T) {
	var filename = "day1_test.txt"

	var expected = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\n"

	if output := GetInputFileContents(filename); output != expected {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
}
