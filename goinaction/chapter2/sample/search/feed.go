package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string
	URI string
	Type string
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}