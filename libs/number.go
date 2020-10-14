package libs

import (
	"strconv"
	"strings"
)

// convert unicode to real integer
func GetNumber(codes string, convertor func(string) int) int {

	result, index := 0, 1
	res := strings.Split(codes, "\\u")

	for i := len(res) - 1; i >= 0; i-- {
		s := res[i]
		if len(s) == 0 {
			continue
		}
		result += convertor(s) * index
		index *= 10
	}
	return result
}

func ParseSalary(s string, parser func(string) int) (low int, high int) {
	// unicode string must be converted first
	s = strconv.QuoteToASCII(s)
	s = strings.Split(s, "/")[0]

	if !strings.Contains(s, "-") {
		num := GetNumber(s, parser)
		low, high = num, num
		return
	}
	salaries := strings.Split(s, "-")

	low = GetNumber(salaries[0], parser)
	high = GetNumber(salaries[1], parser)

	return
}
