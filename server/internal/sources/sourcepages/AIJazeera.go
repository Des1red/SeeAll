package sourcepages

import (
	"encoding/xml"
	"net/http"
	"time"

	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

const ALJAZEERA_RSS = "https://www.aljazeera.com/xml/rss/all.xml"
const ALJAZEERA_MAX = 50

type aljazeeraRSS struct {
	Channel struct {
		Items []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
			Guid    string `xml:"guid"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchAlJazeera() ([]model.Post, error) {

	resp, err := http.Get(ALJAZEERA_RSS)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss aljazeeraRSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range rss.Channel.Items {

		if i >= ALJAZEERA_MAX {
			break
		}

		t, _ := time.Parse(time.RFC1123Z, item.PubDate)

		posts = append(posts, sources.NormalizeNews(
			item.Guid,
			item.Title,
			item.Link,
			"Al Jazeera",
			t.Unix(),
		))
	}

	return posts, nil
}

func init() {
	sources.RegisterSource(sources.Source{
		Name:  "Al Jazeera",
		Type:  "daily",
		Fetch: fetchAlJazeera,
	})
}
