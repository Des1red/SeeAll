package rss

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources/img"
	"SeeAll/internal/sources/normalizer"
	"encoding/xml"
	"time"
)

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
	ItunesImage string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image"`
}
type atomFeed struct {
	XMLName xml.Name    `xml:"feed"`
	Entries []atomEntry `xml:"entry"`
}

func parseAtom(entries []atomEntry, source string, max int) []model.Post {
	model.Usage.Atom++
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

		post := normalizer.NormalizeNews(
			e.ID,
			e.Title,
			link,
			source,
			t.Unix(),
		)

		post.Image = extractAtomImage(e)
		posts = append(posts, post)
	}
	img.EnrichWithOGImages(posts)

	return posts
}
