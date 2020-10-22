package fetch

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(1000 * time.Microsecond)

func Fetch(url string) (*goquery.Document, error) {
	<-rateLimiter
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Printf("fectch %s", url)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error, status code:%d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
