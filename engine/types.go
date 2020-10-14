package engine

import (
	"github.com/PuerkitoBio/goquery"
)

// Item of result
type Item struct {
	Id   string
	Url  string
	Data interface{}
}

type Request struct {
	Url       string
	ParseFunc ParseFunc
}

type ParseResult struct {
	Items    []Item
	Requests []Request
}
type ParseFunc func(g *goquery.Document, url string) ParseResult

func NilParser(g *goquery.Document, url string) ParseResult {
	return ParseResult{}
}
