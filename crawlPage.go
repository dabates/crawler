package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURLpieces, err := url.Parse(rawBaseURL)
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
	if _, ok := pages[normalizedURL]; ok {
		fmt.Printf("already seen this url (%s, %s), incrementing count and returning\n", rawCurrentURL, normalizedURL)
		pages[normalizedURL]++
		return
	}

	pages[normalizedURL] = 1
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}
	urls, err := getURLsFromHTML(htmlBody, rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, thisUrl := range urls {
		crawlPage(rawBaseURL, thisUrl, pages)
	}
}
