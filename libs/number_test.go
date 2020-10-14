package libs

import (
	"strconv"
	"strings"
	"testing"
)

func ConvertNumber(code string) int {
	numberMap := map[string]int{
		"f0e5": 0,
		"e0d8": 1,
		"e5d1": 2,
		"f313": 3,
		"e609": 4,
		"eb5b": 5,
		"f2e1": 6,
		"eb83": 7,
		"f4c4": 8,
		"eb46": 9,
	}
	return numberMap[strings.ToLower(code)]
}

func TestGetNumber(t *testing.T) {
	num := GetNumber(strconv.QuoteToASCII("\uE0D8\uF0E5\uF0E5"), ConvertNumber)
	if num != 100 {
		t.Errorf("Value should be 100, but got %d", num)
	}
}

func TestParseSalary(t *testing.T) {
	low, high := ParseSalary("\uE0D8\uF0E5\uF0E5-\uE0D8\uE5D1\uF0E5/天", ConvertNumber)

	if low != 100 && high != 120 {
		t.Errorf("Low should be 100 and high should be 120, but got %d,%d", low, high)
	}
}
