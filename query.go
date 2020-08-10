package wmovie

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
    // "github.com/wxnacy/wdl"
)

func NewDocumentFromUrl(url string) (*goquery.Document, error) {
    res, err := http.Get(url)
    defer res.Body.Close()
    if err != nil {
        return nil, err
    }
    // Load the HTML document
    fmt.Println(res.Body)
    doc, err := goquery.NewDocumentFromResponse(res)
    return doc, err
}

func ParseVideo(url string) {

    doc, _ := NewDocumentFromUrl(url)

    doc.Find("ul > li > a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Text()
        href, _ := s.Attr("href")
		fmt.Printf("Review %d: %s - %s \n", i, band, href)
	})
}

func ParseMMItems(url string) ([]map[string]string, error) {

    doc, _ := NewDocumentFromUrl(url)

    items := make([]map[string]string, 0, 40)

    doc.Find("ul > li > a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
        href, _ := s.Attr("href")
		// fmt.Printf("Review %d: %s - %s \n", i, band, href)

        items = append(items, map[string]string{
            "title": s.Text(),
            "url": fmt.Sprintf("https://www.697cf.com%s", href),
        })
	})

    return items, nil
}

func ParseMMItem(url string) string{
    fmt.Println(url)
    doc, _ := NewDocumentFromUrl(url)

    var videoUrl string

    doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
        scriptText := s.Text()
        if strings.HasPrefix(scriptText, "generate_down") {
            splits := strings.Split(scriptText, "\"")
            fmt.Println(splits)
            videoUrl = fmt.Sprintf("http://666.maomixia666.com:888%s", splits[1])
        }

	})
    fmt.Println(videoUrl)
    // wdl.Download(videoUrl)
    // _ = wdl
    return videoUrl
}
