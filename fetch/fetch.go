package fetch

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	localCookie "github.com/odysa/go-local-cookie"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(1000 * time.Microsecond)

func FetchWithCookies(url string,domain string) (*goquery.Document, error) {
	<-rateLimiter

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("user-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	req.Header.Add("referer","https://"+domain)
	req.Header.Add("accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	cookies, err := getCookie(domain)
	if err == nil {
		for _, x := range cookies {
			req.AddCookie(&x)
		}
	}

	res, err := client.Do(req)
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

func Fetch(url string) (*goquery.Document, error) {
	<-rateLimiter
	res, err := http.Get(url)

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

func getCookie(url string) ([]http.Cookie, error) {
	cookies, err := localCookie.GetCookiesByDomain(url)
	if err != nil {
		return nil, err
	}
	return cookies, err
}
