package rss

import (
	"SeeAll/internal/sources/img"
	"regexp"
	"strings"
)

var (
	imgRegex    = regexp.MustCompile(`(?i)<img[^>]+src="([^"]+)"`)
	figureRegex = regexp.MustCompile(`(?i)<figure[^>]*>.*?<img[^>]+src="([^"]+)"`)
	badPatterns = []string{
		"1x1", "pixel", "tracker", "beacon",
		"spacer", "blank", "transparent", "ad.",
	}
)

func isUsableImage(u string) bool {
	if u == "" {
		return false
	}
	lower := strings.ToLower(u)
	for _, bad := range badPatterns {
		if strings.Contains(lower, bad) {
			return false
		}
	}
	// skip tiny tracking gifs/pngs
	if strings.HasSuffix(lower, ".gif") {
		return false
	}
	return true
}

func firstUsable(candidates []string) string {
	for _, c := range candidates {
		c = img.CleanImageURL(c)
		if isUsableImage(c) {
			return c
		}
	}
	return ""
}

func extractAtomImage(e atomEntry) string {
	var candidates []string

	for _, m := range e.MediaContents {
		candidates = append(candidates, m.URL)
	}
	if e.MediaThumbnail.URL != "" {
		candidates = append(candidates, e.MediaThumbnail.URL)
	}
	if e.ItunesImage != "" { // ← add this
		candidates = append(candidates, e.ItunesImage)
	}
	// Atom enclosure links
	for _, l := range e.Links {
		if l.Rel == "enclosure" && l.Href != "" {
			candidates = append(candidates, l.Href)
		}
	}

	// figure > img before bare img (usually higher quality)
	for _, body := range []string{e.Summary, e.Content.Body} {
		if body == "" {
			continue
		}
		if m := figureRegex.FindStringSubmatch(body); len(m) > 1 {
			candidates = append(candidates, m[1])
		}
		if m := imgRegex.FindStringSubmatch(body); len(m) > 1 {
			candidates = append(candidates, m[1])
		}
	}

	return firstUsable(candidates)
}

func extractImage(item rssItem) string {
	var candidates []string

	for _, m := range item.MediaContents {
		candidates = append(candidates, m.URL)
	}
	for _, m := range item.MediaGroup.Contents {
		candidates = append(candidates, m.URL)
	}
	if item.MediaThumbnail.URL != "" {
		candidates = append(candidates, item.MediaThumbnail.URL)
	}
	if item.Enclosure.URL != "" {
		candidates = append(candidates, item.Enclosure.URL)
	}
	if item.ItunesImage != "" {
		candidates = append(candidates, item.ItunesImage)
	}

	for _, body := range []string{item.Description, item.ContentEncoded} {
		if body == "" {
			continue
		}
		if m := figureRegex.FindStringSubmatch(body); len(m) > 1 {
			candidates = append(candidates, m[1])
		}
		if m := imgRegex.FindStringSubmatch(body); len(m) > 1 {
			candidates = append(candidates, m[1])
		}
	}

	return firstUsable(candidates)
}
