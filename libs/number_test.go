package libs

import (
	"strconv"
	"strings"
	"testing"
)

func NumberMapper(code string) int {
	numberMap := map[string]int{
		"e535": 0,
		"f693": 1,
		"e5e7": 2,
		"f19e": 3,
		"f31b": 4,
		"ebae": 5,
		"f307": 6,
		"e689": 7,
		"e221": 8,
		"f661": 9,
	}
	return numberMap[strings.ToLower(code)]
}

func TestGetNumber(t *testing.T) {
	num := GetNumber(strconv.QuoteToASCII("\uF693\uEBAE\uE535"), NumberMapper)
	if num != 150 {
		t.Errorf("Value should be 100, but got %d", num)
	}
}

func TestParseSalary(t *testing.T) {
	low, high := ParseSalary("\uF693\uEBAE\uE535-\uF19E\uE535\uE535/天", NumberMapper)

	if low != 150 && high != 300 {
		t.Errorf("Low should be 100 and high should be 120, but got %d,%d", low, high)
	}
}
