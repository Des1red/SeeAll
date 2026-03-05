package sources

import (
	"time"

	"SeeAll/internal/http"
	"SeeAll/internal/model"
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

		image := item.Thumbnail

		if image == "" {
			image = item.Enclosure.Link
		}

		id := item.GUID
		if id == "" {
			id = item.Link
		}

		post := model.Post{
			ID:     id,
			Title:  item.Title,
			URL:    item.Link,
			Image:  cleanImageURL(image),
			Source: source,
			Time:   t.Unix(),
		}

		if post.Title == "" || post.Time == 0 {
			continue
		}

		posts = append(posts, post)
	}

	return posts, nil
}
