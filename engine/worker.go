package engine

import (
	"job-hunter/fetch"
	"log"
)

func worker(request Request) (ParseResult, error) {
	resp, err := fetch.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetch err, url: %s", request.Url)
		return ParseResult{}, err
	}
	return request.ParseFunc(resp, request.Url), nil
}
