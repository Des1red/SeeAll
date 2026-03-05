package sources

import (
	"time"

	"SeeAll/internal/http"
	"SeeAll/internal/model"
)

type jsonRSS struct {
	Items []struct {
		GUID      string `json:"guid"`
		Title     string `json:"title"`
		Link      string `json:"link"`
		PubDate   string `json:"pubDate"`
		Thumbnail string `json:"thumbnail"`
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

		post := model.Post{
			ID:     item.GUID,
			Title:  item.Title,
			URL:    item.Link,
			Image:  item.Thumbnail,
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
