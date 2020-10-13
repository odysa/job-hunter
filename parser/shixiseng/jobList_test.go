package parser

import (
	"job-hunter/fetch"
	"testing"
)

func TestParseJobList(t *testing.T) {
	const urlLen = 20
	resp, err := fetch.Fetch("https://www.shixiseng.com/interns?page=1&type=intern&keyword=IT%E4%BA%92%E8%81%94%E7%BD%91&area=&months=&days=&degree=&official=&enterprise=&salary=-0&publishTime=&sortType=&city=%E5%85%A8%E5%9B%BD&internExtend=")
	if err != nil {
		panic(err)
	}
	result := ParseJobList(resp, "")
	if len(result.Requests) != urlLen {
		t.Errorf("parseJobListError, length should be 20, but got %d", len(result.Requests))
	}
}
