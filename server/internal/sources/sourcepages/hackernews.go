package sourcepages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func isTech(title string) bool {
	keywords := []string{
		"AI", "LLM", "OpenAI", "GPU", "CPU",
		"Linux", "Rust", "Go", "JavaScript",
		"programming", "developer", "database",
		"security", "hacker", "kernel",
	}

	for _, k := range keywords {
		if strings.Contains(strings.ToLower(title), strings.ToLower(k)) {
			return true
		}
	}

	return false
}

func fetchHNFiltered(techOnly bool) ([]model.Post, error) {

	resp, err := http.Get(HN_TOP)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ids []int
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("hn topstories status %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}

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

			resp, err := http.Get(HN_ITEM + strconv.Itoa(id) + ".json")
			if err != nil {
				return
			}
			defer resp.Body.Close()

			var item hnItem
			if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
				return
			}

			if item.Title == "" || item.Time == 0 {
				return
			}

			if techOnly && !isTech(item.Title) {
				return
			}

			if !techOnly && isTech(item.Title) {
				return
			}

			post := sources.NormalizeHN(item.ID, item.Title, item.URL, item.Time, item.Score)

			if post.URL != "" && post.Image == "" {
				post.Image = sources.FetchOGImage(post.URL)
			}

			mu.Lock()
			posts = append(posts, post)
			mu.Unlock()

		}(id)
	}

	wg.Wait()

	return posts, nil
}

func fetchHNTech() ([]model.Post, error) {
	return fetchHNFiltered(true)
}

func fetchHNGeneral() ([]model.Post, error) {
	return fetchHNFiltered(false)
}

func init() {

	sources.RegisterSource(sources.Source{
		Name:  "HackerNews",
		Type:  model.AudienceTech,
		Fetch: fetchHNTech,
	})

	sources.RegisterSource(sources.Source{
		Name:  "HackerNews",
		Type:  model.AudienceGeneral,
		Fetch: fetchHNGeneral,
	})

}
