package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/Imnarka/GoSearch/pkg/scraper"
)

func main() {
	searchWord := flag.String("s", "", "Слово для поиска")
	flag.Parse()
	urls := []string{"https://go.dev", "https://golang.org"}
	var allLinks []string
	for _, url := range urls {
		links, err := scraper.Scrap(url)
		if err != nil {
			log.Fatalf("Ошибка при сканировании %s: %v", url, err)
		}
		allLinks = append(allLinks, links...)
	}
	if *searchWord != "" {
		filteredLinks := filterLinksByWord(allLinks, *searchWord)
		fmt.Println("Ссылки, содержащие слово:", *searchWord)
		for _, link := range filteredLinks {
			fmt.Println(link)
		}
	} else {
		fmt.Println("Все найденные ссылки:")
		for _, link := range allLinks {
			fmt.Println(link)
		}
	}
}

func filterLinksByWord(links []string, word string) []string {
	var filtered []string
	for _, link := range links {
		if strings.Contains(link, word) {
			filtered = append(filtered, link)
		}
	}
	return filtered
}
