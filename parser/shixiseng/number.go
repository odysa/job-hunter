package shixiseng

import "strings"

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
