package fetch

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	doc, err := Fetch("https://www.sina.com.cn")
	if err != nil {
		panic(err)
	}
	title := doc.Find("title").First().Text()
	if title != "新浪首页" {
		t.Errorf("fetch failed")
	}
}

func TestFetchWithCookies(t *testing.T) {
	doc, err := FetchWithCookies("https://sou.zhaopin.com/?jl=530&kw=Java%E5%BC%80%E5%8F%91&kt=3", "zhaopin.com")
	if err != nil {
		panic(err)
	}
	len := doc.Find("a.not-login__position").Text()
	fmt.Println(len)
}