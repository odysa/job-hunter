package shixiseng

import (
	"job-hunter/fetch"
	"testing"
)

func TestParseJob(t *testing.T) {
	resp, err := fetch.Fetch("https://www.shixiseng.com/intern/inn_onwoc69o51rg?pcm=pc_SearchList")
	if err != nil {
		panic(err)
	}
	res := ParseJob(resp, "")
	if len(res.Items) == 0 {
		t.Errorf("Result should not be empty")
	}
}
