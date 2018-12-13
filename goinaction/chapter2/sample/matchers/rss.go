package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/purelind/tour-go/goinaction/chapter2/sample/search"
)

type (
	item struct {
		XMLName xml.Name
		PubDate string
		Title string
		Description string
		Link string
		GUID string
		GeoRssPoint string
	}

	image struct {
		XMLName xml.Name
		URL string
		Title string
		Link string
	}

	channel struct {
		XMLName xml.Name
		Title string
		Description string
		Link string
		PubDate string
		LastBuildDate string
		TTL string
		Language string
		ManagingEditor string
		WebMaster string
		Image image
		Item []item
	}

	rssDocument struct {
		XMLName xml.Name
		Channel channel
	}
)

type rssMatcher struct {

}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search looks at the document for the specified search term.
func (m rssMatcher) Search(feed *search.Feed, searchTerm string)  ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// Retrieve the data to search
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// Check the title for the search term.
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field: "Title",
				Content: channelItem.Title,
			})
		}

		// Check the description for the search term.
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field: "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}


// retrieve performs a HTTP Get request for the rss feed and decodes the results.
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	// Retrieve the rss feed document from the web.
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
