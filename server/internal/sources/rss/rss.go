package rss

import (
	"net/http"
	"time"

	"SeeAll/internal/model"
	"SeeAll/internal/sources/img"
	"SeeAll/internal/sources/normalizer"
)

var httpClient = &http.Client{
	Timeout: 8 * time.Second,
}

type rssItem struct {
	Title          string `xml:"title"`
	Link           string `xml:"link"`
	Guid           string `xml:"guid"`
	PubDate        string `xml:"pubDate"`
	Description    string `xml:"description"`
	ContentEncoded string `xml:"content:encoded"`

	MediaContents []struct {
		URL  string `xml:"url,attr"`
		Type string `xml:"type,attr"`
	} `xml:"media:content"`
	MediaGroup struct {
		Contents []struct {
			URL  string `xml:"url,attr"`
			Type string `xml:"type,attr"`
		} `xml:"media:content"`
	} `xml:"media:group"`
	MediaThumbnail struct {
		URL string `xml:"url,attr"`
	} `xml:"media:thumbnail"`

	Enclosure struct {
		URL string `xml:"url,attr"`
	} `xml:"enclosure"`
	ItunesImage string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image"`
}

type genericRSS struct {
	Channel struct {
		Items []rssItem `xml:"item"`
	} `xml:"channel"`
}

func parseRSS(items []rssItem, source string, max int) []model.Post {
	model.Usage.RSS++
	var posts []model.Post

	for i, item := range items {

		if i >= max {
			break
		}

		if item.Title == "" || item.Link == "" {
			continue
		}

		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			t, _ = time.Parse(time.RFC1123, item.PubDate)
		}

		id := item.Guid
		if id == "" {
			id = item.Link
		}

		post := normalizer.NormalizeNews(
			id,
			item.Title,
			item.Link,
			source,
			t.Unix(),
		)

		post.Image = extractImage(item) // inline only, no HTTP

		posts = append(posts, post)
	}

	// Fan out OG fetches only for posts that need it
	img.EnrichWithOGImages(posts)

	return posts
}
