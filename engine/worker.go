package engine

import (
	"job-hunter/fetch"
	"time"
)

var rateLimiter = time.Tick(300 * time.Microsecond)

func worker(request Request) (ParseResult, error) {
	<-rateLimiter

	doc, err := fetch.Fetch(request.Url)
	if err != nil {
		return ParseResult{}, err
	}

	return request.ParseFunc(doc, request.Url), nil
}
