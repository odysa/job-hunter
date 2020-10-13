package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseJobList(t *testing.T) {
	const urlLen = 20
	resp, err := ioutil.ReadFile("jobListTest.html")
	if err != nil {
		panic(err)
	}
	result := ParseJobList(resp, "")
	if len(result.Requests) != urlLen {
		t.Errorf("parseJobListError, length should be 20, but got %d", len(result.Requests))
	}
}
