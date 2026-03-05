package sourcepages

import (
	"encoding/xml"
	"net/http"
	"time"

	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

const REUTERS_RSS = "https://www.reutersagency.com/feed/?best-topics=world&post_type=best"
const REUTERS_MAX = 50

type reutersRSS struct {
	Channel struct {
		Items []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
			Guid    string `xml:"guid"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchReuters() ([]model.Post, error) {

	resp, err := http.Get(REUTERS_RSS)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss reutersRSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range rss.Channel.Items {

		if i >= REUTERS_MAX {
			break
		}

		t, _ := time.Parse(time.RFC1123Z, item.PubDate)

		posts = append(posts, sources.NormalizeNews(
			item.Guid,
			item.Title,
			item.Link,
			"Reuters",
			t.Unix(),
		))
	}

	return posts, nil
}

func init() {
	sources.RegisterSource(sources.Source{
		Name:  "Reuters",
		Type:  "daily",
		Fetch: fetchReuters,
	})
}
