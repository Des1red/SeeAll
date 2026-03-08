package sources

import (
	"sort"
	"sync"
	"time"

	"SeeAll/internal/devmode"
	"SeeAll/internal/model"
)

func FetchByType(t string) ([]model.Post, error) {

	var filtered []Source
	for _, s := range GetSources() {
		if s.Type == t {
			filtered = append(filtered, s)
		}
	}

	var wg sync.WaitGroup
	resultsChan := make(chan []model.Post, len(filtered))

	for _, s := range filtered {

		wg.Add(1)

		go func(src Source) {
			defer wg.Done()

			start := time.Now()
			posts, err := src.Fetch()
			devmode.RecordSource(src.Name, time.Since(start), len(posts))
			if err != nil || len(posts) == 0 {
				return
			}

			resultsChan <- posts

		}(s)
	}

	wg.Wait()
	close(resultsChan)

	var results []model.Post

	for posts := range resultsChan {
		results = append(results, posts...)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Time > results[j].Time
	})

	seen := make(map[string]bool)
	var dedup []model.Post

	for _, p := range results {
		if seen[p.URL] {
			continue
		}
		seen[p.URL] = true
		dedup = append(dedup, p)
	}

	return dedup, nil
}
