package wmovie

import (
    "fmt"
    "testing"
    "github.com/PuerkitoBio/goquery"
)

func TestNewDocumentFromUrl(t *testing.T) {
    doc, err := NewDocumentFromUrl("https://www.baidu.com")
    if err != nil {
        t.Error(err)
    }

    doc.Find("title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Text()
		fmt.Printf("Review %d: %s \n", i, band)
	})
}
