package sources

import (
	"encoding/xml"
	"io"
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
}

type genericRSS struct {
	Channel struct {
		Items []rssItem `xml:"item"`
	} `xml:"channel"`
}

type atomEntry struct {
	Title   string `xml:"title"`
	ID      string `xml:"id"`
	Updated string `xml:"updated"`

	Links []struct {
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr"`
	} `xml:"link"`

	Summary string `xml:"summary"`
	Content struct {
		Body string `xml:",innerxml"`
	} `xml:"content"`

	MediaThumbnail struct {
		URL string `xml:"url,attr"`
	} `xml:"media:thumbnail"`

	MediaContents []struct {
		URL string `xml:"url,attr"`
	} `xml:"media:content"`
}
type atomFeed struct {
	XMLName xml.Name    `xml:"feed"`
	Entries []atomEntry `xml:"entry"`
}

var imgRegex = regexp.MustCompile(`(?i)<img[^>]+src="([^"]+)"`)

func FetchRSS(url string, source string, max int, browser bool) ([]model.Post, error) {

	var resp *http.Response
	var err error

	if browser {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("User-Agent", "Mozilla/5.0")

		resp, err = httpClient.Do(req)

	} else {

		resp, err = httpClient.Get(url)

	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss genericRSS
	if err := xml.Unmarshal(data, &rss); err == nil && len(rss.Channel.Items) > 0 {
		return parseRSS(rss.Channel.Items, source, max), nil
	}

	var atom atomFeed
	if err := xml.Unmarshal(data, &atom); err == nil && len(atom.Entries) > 0 {
		return parseAtom(atom.Entries, source, max), nil
	}

	return nil, nil
}

func parseRSS(items []rssItem, source string, max int) []model.Post {

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

	return posts
}

func parseAtom(entries []atomEntry, source string, max int) []model.Post {

	var posts []model.Post

	for i, e := range entries {
		if i >= max {
			break
		}
		var link string

		for _, l := range e.Links {
			if l.Rel == "alternate" || l.Rel == "" {
				link = l.Href
				break
			}
		}
		if link == "" && len(e.Links) > 0 {
			link = e.Links[0].Href
		}
		if e.Title == "" || link == "" {
			continue
		}
		t, _ := time.Parse(time.RFC3339, e.Updated)

		post := NormalizeNews(
			e.ID,
			e.Title,
			link,
			source,
			t.Unix(),
		)

		post.Image = extractAtomImage(e)

		posts = append(posts, post)
	}

	return posts
}

func extractAtomImage(e atomEntry) string {

	for _, m := range e.MediaContents {
		if m.URL != "" {
			return cleanImageURL(m.URL)
		}
	}

	if e.MediaThumbnail.URL != "" {
		return cleanImageURL(e.MediaThumbnail.URL)
	}

	if e.Summary != "" {
		match := imgRegex.FindStringSubmatch(e.Summary)
		if len(match) > 1 {
			return cleanImageURL(match[1])
		}
	}

	if e.Content.Body != "" {
		match := imgRegex.FindStringSubmatch(e.Content.Body)
		if len(match) > 1 {
			return cleanImageURL(match[1])
		}
	}

	return ""
}

func extractImage(item rssItem) string {

	// media:content (multiple)
	for _, m := range item.MediaContents {
		if m.URL != "" {
			return cleanImageURL(m.URL)
		}
	}
	// media:group
	for _, m := range item.MediaGroup.Contents {
		if m.URL != "" {
			return cleanImageURL(m.URL)
		}
	}
	image := item.MediaThumbnail.URL

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

	return cleanImageURL(image)
}
