package sourcepages

import (
	"encoding/xml"
	"net/http"
	"time"

	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

const NPR_RSS = "https://feeds.npr.org/1001/rss.xml"
const NPR_MAX = 50

type nprRSS struct {
	Channel struct {
		Items []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
			Guid    string `xml:"guid"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchNPR() ([]model.Post, error) {

	resp, err := http.Get(NPR_RSS)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss nprRSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range rss.Channel.Items {

		if i >= NPR_MAX {
			break
		}

		t, _ := time.Parse(time.RFC1123Z, item.PubDate)

		posts = append(posts, sources.NormalizeNews(
			item.Guid,
			item.Title,
			item.Link,
			"NPR",
			t.Unix(),
		))
	}

	return posts, nil
}

func init() {
	sources.RegisterSource(sources.Source{
		Name:  "NPR",
		Type:  "daily",
		Fetch: fetchNPR,
	})
}
