package zhilian

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"job-hunter/engine"
)

func ParseJobList(doc *goquery.Document, url string) engine.ParseResult {
	result := engine.ParseResult{}
	doc.Find("div").Length()
	doc.Find("a.contentpile__content__wrapper__item__info").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if ok {
			fmt.Println(url)
		}
	})
	return result
}
