package rss

import (
	"time"

	"SeeAll/internal/http"
	"SeeAll/internal/model"
	"SeeAll/internal/sources/img"
	"SeeAll/internal/sources/normalizer"
)

type jsonRSS struct {
	Items []struct {
		GUID    string `json:"guid"`
		Title   string `json:"title"`
		Link    string `json:"link"`
		PubDate string `json:"pubDate"`

		Thumbnail string `json:"thumbnail"`

		Enclosure struct {
			Link string `json:"link"`
		} `json:"enclosure"`
	} `json:"items"`
}

func FetchJSONRSS(url string, source string, max int) ([]model.Post, error) {
	model.Usage.JSONRSS++

	var data jsonRSS
	err := http.FetchJSON(url, &data)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range data.Items {

		if i >= max {
			break
		}

		t, _ := time.Parse(time.RFC1123Z, item.PubDate)

		id := item.GUID
		if id == "" {
			id = item.Link
		}

		image := item.Thumbnail
		if image == "" {
			image = item.Enclosure.Link
		}

		post := normalizer.NormalizeNews(
			id,
			item.Title,
			item.Link,
			source,
			t.Unix(),
		)

		post.Image = img.CleanImageURL(image)

		if post.Title == "" || post.Time == 0 {
			continue
		}

		posts = append(posts, post)
	}

	// Fan out OG fetches only for posts that need it
	img.EnrichWithOGImages(posts)
	return posts, nil
}
