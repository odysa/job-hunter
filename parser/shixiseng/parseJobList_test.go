package shixiseng

import (
	"job-hunter/fetch"
	"testing"
)

func TestParseJobList(t *testing.T) {
	const urlLen = 20
	resp, err := fetch.Fetch("https://www.shixiseng.com/interns?page=1&type=intern&city=%E5%85%A8%E5%9B%BD")
	if err != nil {
		panic(err)
	}
	result := ParseJobList(resp, "")
	if len(result.Requests) != urlLen {
		t.Errorf("parseJobListError, length should be 20, but got %d", len(result.Requests))
	}
}
