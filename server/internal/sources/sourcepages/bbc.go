package sourcepages

import (
	"encoding/xml"
	"net/http"
	"time"

	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

const BBC_RSS = "https://feeds.bbci.co.uk/news/world/rss.xml"
const BBC_MAX = 50

type bbcRSS struct {
	Channel struct {
		Items []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchBBC() ([]model.Post, error) {

	resp, err := http.Get(BBC_RSS)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss bbcRSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range rss.Channel.Items {

		if i >= BBC_MAX {
			break
		}

		t, _ := time.Parse(time.RFC1123Z, item.PubDate)

		posts = append(posts, sources.NormalizeNews(
			item.Link,
			item.Title,
			item.Link,
			"BBC",
			t.Unix(),
		))
	}

	return posts, nil
}

func init() {
	sources.RegisterSource(sources.Source{
		Name:  "BBC",
		Type:  "daily",
		Fetch: fetchBBC,
	})
}
