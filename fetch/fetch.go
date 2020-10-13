package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Microsecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error, status code:%d", res.StatusCode)
	}
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}
