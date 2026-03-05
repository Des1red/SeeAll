package sourcepages

import (
	"SeeAll/internal/http"
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

const LOBSTERS_RSS = "https://api.rss2json.com/v1/api.json?rss_url=https://lobste.rs/rss"
const LOBSTERS_MAX = 50

type lobstersResp struct {
	Items []struct {
		GUID    string `json:"guid"`
		Title   string `json:"title"`
		Link    string `json:"link"`
		PubDate string `json:"pubDate"`
	} `json:"items"`
}

func fetchLobsters() ([]model.Post, error) {

	var data lobstersResp
	err := http.FetchJSON(LOBSTERS_RSS, &data)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range data.Items {

		if i >= LOBSTERS_MAX {
			break
		}

		post := sources.NormalizeLobsters(
			item.GUID,
			item.Title,
			item.Link,
			item.PubDate,
		)

		if post.Title == "" || post.Time == 0 {
			continue
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func init() {
	sources.RegisterSource(sources.Source{
		Name:  "Lobsters",
		Type:  "live",
		Fetch: fetchLobsters,
	})
}
