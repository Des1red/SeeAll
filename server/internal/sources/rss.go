package sources

import (
	"encoding/xml"
	"net/http"
	"regexp"
	"time"

	"SeeAll/internal/model"
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

	MediaContent struct {
		URL string `xml:"url,attr"`
	} `xml:"media:content"`

	MediaThumbnail struct {
		URL string `xml:"url,attr"`
	} `xml:"media:thumbnail"`

	Enclosure struct {
		URL string `xml:"url,attr"`
	} `xml:"enclosure"`
}

type genericRSS struct {
	Channel struct {
		Items []rssItem `xml:"item"`
	} `xml:"channel"`
}

var imgRegex = regexp.MustCompile(`(?i)<img[^>]+src="([^"]+)"`)

func FetchRSS(url string, source string, max int) ([]model.Post, error) {

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss genericRSS

	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range rss.Channel.Items {

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

		image := extractImage(item)

		post := NormalizeNews(
			id,
			item.Title,
			item.Link,
			source,
			t.Unix(),
		)

		post.Image = image

		posts = append(posts, post)
	}

	return posts, nil
}

func extractImage(item rssItem) string {

	image := item.MediaContent.URL

	if image == "" {
		image = item.MediaThumbnail.URL
	}

	if image == "" {
		image = item.Enclosure.URL
	}

	if image == "" && item.Description != "" {
		match := imgRegex.FindStringSubmatch(item.Description)
		if len(match) > 1 {
			image = match[1]
		}
	}

	if image == "" && item.ContentEncoded != "" {
		match := imgRegex.FindStringSubmatch(item.ContentEncoded)
		if len(match) > 1 {
			image = match[1]
		}
	}

	cleanImageURL(image)

	return image
}
