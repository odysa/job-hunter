package zhilian

import (
	"github.com/PuerkitoBio/goquery"
	"job-hunter/engine"
	"strings"
)

func ParseCity(doc *goquery.Document, url string) engine.ParseResult {
	result := engine.ParseResult{}
	doc.Find("li.cities-show__list--li").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Find("a.cities-show__list--href").Attr("href")
		if ok {
			url = strings.Replace(url, "/", "", 2)
			result.Requests = append(result.Requests, engine.Request{
				Url:       url,
				ParseFunc: engine.NilParser,
			})
		}
	})
	return result
}
