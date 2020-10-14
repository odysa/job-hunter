package shixiseng

import (
	"github.com/PuerkitoBio/goquery"
	"job-hunter/engine"
)

func ParseJobList(doc *goquery.Document, url string) engine.ParseResult {
	result := engine.ParseResult{}
	doc.Find("[href*='intern']").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if ok {
			result.Requests = append(result.Requests, engine.Request{
				Url:       url,
				ParseFunc: ParseJob,
			})
		}
	})
	return result
}
