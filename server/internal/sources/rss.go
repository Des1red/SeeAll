package sources

import (
	"encoding/xml"
	"net/http"
	"time"

	"SeeAll/internal/model"
)

var httpClient = &http.Client{
	Timeout: 8 * time.Second,
}

type genericRSS struct {
	Channel struct {
		Items []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			Guid    string `xml:"guid"`
			PubDate string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func FetchRSS(url string, source string, max int) ([]model.Post, error) {

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss genericRSS

	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range rss.Channel.Items {

		if i >= max {
			break
		}

		if item.Title == "" || item.Link == "" {
			continue
		}

		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			t, _ = time.Parse(time.RFC1123, item.PubDate)
		}

		id := item.Guid
		if id == "" {
			id = item.Link
		}

		posts = append(posts, NormalizeNews(
			id,
			item.Title,
			item.Link,
			source,
			t.Unix(),
		))
	}

	return posts, nil
}
