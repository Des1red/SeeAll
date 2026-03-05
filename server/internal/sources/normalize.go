package sources

import (
	"strconv"
	"time"

	"SeeAll/internal/model"
)

func NormalizeHN(id int, title, url string, ts int64, score *int) model.Post {

	if url == "" {
		url = "https://news.ycombinator.com/item?id=" + strconv.Itoa(id)
	}

	return model.Post{
		ID:     strconv.Itoa(id),
		Title:  title,
		URL:    url,
		Source: "HackerNews",
		Time:   ts,
		Score:  score,
	}
}

func NormalizeNews(id, title, url, source string, ts int64) model.Post {

	return model.Post{
		ID:     id,
		Title:  title,
		URL:    url,
		Source: source,
		Time:   ts,
		Score:  nil,
	}
}

func NormalizeReddit(id, title, permalink string, created int64, score *int) model.Post {

	url := "https://reddit.com" + permalink

	return model.Post{
		ID:     id,
		Title:  title,
		URL:    url,
		Source: "Reddit",
		Time:   created,
		Score:  score,
	}
}

func NormalizeLobsters(guid, title, link, pubDate string) model.Post {

	var ts int64

	if pubDate != "" {

		// Try RFC1123 with timezone offset
		t, err := time.Parse(time.RFC1123Z, pubDate)

		// Some RSS feeds use RFC1123 without numeric zone
		if err != nil {
			t, err = time.Parse(time.RFC1123, pubDate)
		}

		if err == nil {
			ts = t.Unix()
		}
	}

	return model.Post{
		ID:     "lobsters-" + guid,
		Title:  title,
		URL:    link,
		Source: "Lobsters",
		Time:   ts,
		Score:  nil,
	}
}
