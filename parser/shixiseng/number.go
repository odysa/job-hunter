package shixiseng

import "strings"

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
