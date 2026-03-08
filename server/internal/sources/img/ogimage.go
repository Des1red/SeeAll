package img

import (
	"SeeAll/internal/model"
	"io"
	"net/http"
	"regexp"
	"sync"
	"time"
)

var ogImageRegex = regexp.MustCompile(`(?i)<meta[^>]+property=["']og:image["'][^>]+content=["']([^"']+)`)

var ogClient = &http.Client{
	Timeout: 3 * time.Second,
}

func fetchOGImageUncached(url string) string {

	resp, err := ogClient.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	reader := io.LimitReader(resp.Body, 32768)

	body, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}

	match := ogImageRegex.FindSubmatch(body)
	if len(match) > 1 {
		return CleanImageURL(string(match[1]))
	}

	return ""
}

var (
	ogCache   = make(map[string]string)
	ogCacheMu sync.RWMutex
)

func fetchOGImage(url string) string {
	ogCacheMu.RLock()
	if v, ok := ogCache[url]; ok {
		ogCacheMu.RUnlock()
		return v
	}
	ogCacheMu.RUnlock()

	result := fetchOGImageUncached(url)

	ogCacheMu.Lock()
	ogCache[url] = result
	ogCacheMu.Unlock()

	return result
}

func EnrichWithOGImages(posts []model.Post) {
	var wg sync.WaitGroup

	for i := range posts {
		if posts[i].Image != "" {
			continue
		}

		wg.Add(1)

		go func(i int, url string) {
			defer wg.Done()
			posts[i].Image = fetchOGImage(url)
		}(i, posts[i].URL)
	}

	wg.Wait()
}
