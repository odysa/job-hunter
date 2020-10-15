package zhilian

import (
	"job-hunter/fetch"
	"testing"
)

func TestParseCity(t *testing.T) {
	doc, err := fetch.Fetch("https://www.zhaopin.com/citymap")
	if err != nil {
		panic(err)
	}
	result := ParseCity(doc, "")
	if len(result.Requests) != 531 {
		t.Errorf("length should be 531, but got %d", len(result.Requests))
	}
}
