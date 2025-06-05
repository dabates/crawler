package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *Config) crawlPage(rawCurrentURL string) {
	baseURLpieces, err := url.Parse(cfg.baseURL.String())
	if err != nil {
		log.Fatal(err)
	}
	currentURLpieces, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}

	if currentURLpieces.Host != baseURLpieces.Host {
		fmt.Printf("not same host (%s), returning\n", rawCurrentURL)
		return
	}

	normalizedURL := normalizeURL(rawCurrentURL)
	if _, ok := cfg.pages[normalizedURL]; ok {
		fmt.Printf("already seen this url (%s, %s), incrementing count and returning\n", rawCurrentURL, normalizedURL)
		cfg.pages[normalizedURL]++
		return
	}

	cfg, pages[normalizedURL] = 1
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}
	urls, err := getURLsFromHTML(htmlBody, rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, thisUrl := range urls {
		crawlPage(thisUrl)
	}
}

func (cfg *Config) addPageVisit(normalizedURL string) (isFirst bool) {

}
