package sources

import (
	"sync"

	"SeeAll/internal/model"
)

func FetchByType(t string) ([]model.Post, error) {

	srcs := GetSources()

	var wg sync.WaitGroup
	resultsChan := make(chan []model.Post, len(srcs))

	for _, s := range srcs {

		if s.Type != t {
			continue
		}

		wg.Add(1)

		go func(src Source) {
			defer wg.Done()

			posts, err := src.Fetch()
			if err != nil {
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

	return results, nil
}
