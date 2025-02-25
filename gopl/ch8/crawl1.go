package main

import (
	"os"
	"fmt"
	"log"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	// start with command param
	go func() { worklist <- os.Args[1:] }()

	// crawl web in parallel
	seen := make(map[string]bool)
	for list := range worklist {
		if != seen[link] {
			seen[link] = true
			go func(link string) {
				worklist <- crawl(link)
			}(link)
		}
	}
}
