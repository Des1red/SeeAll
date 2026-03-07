package img

import (
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

func FetchOGImage(url string) string {
	ogCacheMu.RLock()
	if v, ok := ogCache[url]; ok {
		ogCacheMu.RUnlock()
		return v
	}
	ogCacheMu.RUnlock()

	result := fetchOGImageUncached(url) // your existing logic

	ogCacheMu.Lock()
	ogCache[url] = result
	ogCacheMu.Unlock()

	return result
}
