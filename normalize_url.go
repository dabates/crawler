package main

import (
	"fmt"
	"log"
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
