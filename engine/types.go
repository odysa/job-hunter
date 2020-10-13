package engine

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
type ParseFunc func(contents []byte, url string) ParseResult

func NilParser(contents []byte, url string) ParseResult {
	return ParseResult{}
}