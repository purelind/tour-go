package search

import (
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, SearchTerm string, results chan<- *Result) {
	searchResults, er := matcher.Search(feed, SearchTerm)
	if err != nil {
		log.Printf(err)
		return
	}
	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
