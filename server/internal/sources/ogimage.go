package sources

import (
	"io"
	"net/http"
	"regexp"
)

var ogImageRegex = regexp.MustCompile(`(?i)<meta[^>]+property=["']og:image["'][^>]+content=["']([^"']+)`)

func FetchOGImage(url string) string {

	resp, err := http.Get(url)
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
		return cleanImageURL(string(match[1]))
	}

	return ""
}
