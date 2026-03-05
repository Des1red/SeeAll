package sourcepages

import (
	"time"

	"SeeAll/internal/http"
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

const (
	REDDIT_HOT = "https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/worldnews/.rss"
	REDDIT_NEW = "https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/worldnews/new/.rss"
	REDDIT_MAX = 50
)

type redditResp struct {
	Items []struct {
		GUID    string `json:"guid"`
		Title   string `json:"title"`
		Link    string `json:"link"`
		PubDate string `json:"pubDate"`
	} `json:"items"`
}

func fetchReddit(url string) ([]model.Post, error) {

	var data redditResp
	err := http.FetchJSON(url, &data)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range data.Items {

		if i >= REDDIT_MAX {
			break
		}

		t, _ := time.Parse(time.RFC1123Z, item.PubDate)

		posts = append(posts, model.Post{
			ID:     "reddit-" + item.GUID,
			Title:  item.Title,
			URL:    item.Link,
			Source: "Reddit",
			Time:   t.Unix(),
		})
	}

	return posts, nil
}

func init() {

	sources.RegisterSource(sources.Source{
		Name: "Reddit",
		Type: "live",
		Fetch: func() ([]model.Post, error) {
			return fetchReddit(REDDIT_HOT)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "RedditNew",
		Type: "live",
		Fetch: func() ([]model.Post, error) {
			return fetchReddit(REDDIT_NEW)
		},
	})
}
