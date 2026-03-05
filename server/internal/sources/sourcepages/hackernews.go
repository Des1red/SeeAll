package sourcepages

import (
	"encoding/json"
	"net/http"
	"sync"

	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

const (
	HN_TOP         = "https://hacker-news.firebaseio.com/v0/topstories.json"
	HN_ITEM        = "https://hacker-news.firebaseio.com/v0/item/"
	HN_MAX         = 100
	HN_CONCURRENCY = 10
)

type hnItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Time  int64  `json:"time"`
	Score *int   `json:"score"`
}

func fetchHN() ([]model.Post, error) {

	resp, err := http.Get(HN_TOP)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ids []int
	json.NewDecoder(resp.Body).Decode(&ids)

	if len(ids) > HN_MAX {
		ids = ids[:HN_MAX]
	}

	var posts []model.Post
	var mu sync.Mutex
	var wg sync.WaitGroup

	sem := make(chan struct{}, HN_CONCURRENCY)

	for _, id := range ids {

		wg.Add(1)
		sem <- struct{}{}

		go func(id int) {
			defer wg.Done()
			defer func() { <-sem }()

			resp, err := http.Get(HN_ITEM + string(rune(id)) + ".json")
			if err != nil {
				return
			}
			defer resp.Body.Close()

			var item hnItem
			err = json.NewDecoder(resp.Body).Decode(&item)
			if err != nil {
				return
			}

			post := sources.NormalizeHN(item.ID, item.Title, item.URL, item.Time, item.Score)

			mu.Lock()
			posts = append(posts, post)
			mu.Unlock()

		}(id)
	}

	wg.Wait()

	return posts, nil
}

func init() {
	sources.RegisterSource(sources.Source{
		Name:  "HackerNews",
		Type:  "live",
		Fetch: fetchHN,
	})
}
