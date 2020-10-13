package parser

import (
	"job-hunter/engine"
	"regexp"
)

var urlRe = regexp.MustCompile(`href="(https://www.shixiseng.com/intern/([^?]+)\?pcm=pc_SearchList)"`)

func ParseJobList(contents []byte, url string) engine.ParseResult {
	result := engine.ParseResult{}
	matches := urlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: engine.NilParser,
		})
		result.Items = append(result.Items, engine.Item{
			Url:  string(m[1]),
			Id:   string(m[2]),
			Data: nil,
		})
	}
	return result
}
