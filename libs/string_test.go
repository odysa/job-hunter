package libs

import (
	"regexp"
	"testing"
)

func TestExtractString(t *testing.T) {
	pageRe := regexp.MustCompile(`.*page=([0-9]+).*`)
	s := ExtractString([]byte("https://www.shixiseng.com/interns?page=1&type=intern&city=%E5%85%A8%E5%9B%BD"), pageRe)
	if s != "1" {
		t.Errorf("string should be 1,but got %v", len(s))
	}
}
