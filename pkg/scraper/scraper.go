package scraper

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func Scrap(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error fetching URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error parsing HTML: %v", err)
	}
	links := make([]string, 0)
	var parseLinks func(*html.Node)
	parseLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseLinks(c)
		}
	}
	parseLinks(doc)
	return links, nil
}
