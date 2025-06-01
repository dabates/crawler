package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)
import "net/url"

func normalizeURL(sourceUrl string) string {
	pieces, err := url.Parse(sourceUrl)
	if err != nil {
		log.Fatal(err)
	}

	path := pieces.Path
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return fmt.Sprintf("%s%s", pieces.Host, path)
}
func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, err
	}

	reader := strings.NewReader(htmlBody)

	// Parse
	doc, err := html.Parse(reader)
	if err != nil {
		return []string{}, err
	}

	var urls []string

	// recursive func to traverse the nodes

	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					u, err := url.Parse(a.Val)
					if err != nil {
						continue
					}

					// Make them relative
					resolvedURL := baseURL.ResolveReference(u)
					urls = append(urls, resolvedURL.String())
					// done here
					break
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	visitNode(doc)

	if len(urls) == 0 {
		return []string{}, nil
	}

	return urls, nil
}
