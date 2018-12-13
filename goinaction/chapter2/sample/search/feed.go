package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(data.File)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var feed []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
