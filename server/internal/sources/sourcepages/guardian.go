package sourcepages

import (
	"encoding/xml"
	"net/http"
	"time"

	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

const GUARDIAN_RSS = "https://www.theguardian.com/world/rss"
const GUARDIAN_MAX = 50

type guardianRSS struct {
	Channel struct {
		Items []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchGuardian() ([]model.Post, error) {

	resp, err := http.Get(GUARDIAN_RSS)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss guardianRSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range rss.Channel.Items {

		if i >= GUARDIAN_MAX {
			break
		}

		t, _ := time.Parse(time.RFC1123Z, item.PubDate)

		posts = append(posts, sources.NormalizeNews(
			item.Link,
			item.Title,
			item.Link,
			"Guardian",
			t.Unix(),
		))
	}

	return posts, nil
}

func init() {
	sources.RegisterSource(sources.Source{
		Name:  "Guardian",
		Type:  "daily",
		Fetch: fetchGuardian,
	})
}
