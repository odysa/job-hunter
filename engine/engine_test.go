package engine

import (
	"job-hunter/libs"
	"regexp"
	"strconv"
	"testing"
)

func TestGenerateUrlPage(t *testing.T) {
	const limit = 100
	g := GenerateUrlPage(Request{
		Url:       "http://a.com?",
		ParseFunc: NilParser,
	})
	pageRe := regexp.MustCompile(`.*page=([0-9]+).*`)
	for i := 1; i <= limit; i++ {
		url := g()
		s := libs.ExtractString([]byte(url.Url), pageRe)
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if num != i {
			t.Errorf("page should be %d,but got %d", i, num)
		}
	}
}
