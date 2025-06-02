package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Println("starting crawl of:", args[0])

	//fmt.Println(getHTML(args[0]))
	urls := make(map[string]int)
	crawlPage(args[0], args[0], urls)

	// print the map
	for thisUrl, count := range urls {
		fmt.Printf("%s: %d\n", thisUrl, count)
	}
}
