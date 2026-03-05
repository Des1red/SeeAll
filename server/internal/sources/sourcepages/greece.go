package sourcepages

import (
	"encoding/xml"
	"net/http"
	"time"

	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

/* ==============================
   KATHIMERINI
============================== */

const KATHIMERINI_RSS = "https://www.ekathimerini.com/rss/news/"
const KATHIMERINI_MAX = 50

type kathimeriniRSS struct {
	Channel struct {
		Items []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
			Guid    string `xml:"guid"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchKathimerini() ([]model.Post, error) {

	resp, err := http.Get(KATHIMERINI_RSS)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss kathimeriniRSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range rss.Channel.Items {

		if i >= KATHIMERINI_MAX {
			break
		}

		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			t, _ = time.Parse(time.RFC1123, item.PubDate)
		}

		posts = append(posts, sources.NormalizeNews(
			item.Guid,
			item.Title,
			item.Link,
			"Kathimerini",
			t.Unix(),
		))
	}

	return posts, nil
}

func init() {
	sources.RegisterSource(sources.Source{
		Name:  "Kathimerini",
		Type:  "greece",
		Fetch: fetchKathimerini,
	})
}

/* ==============================
   RIZOSPASTIS
============================== */

const RIZOSPASTIS_RSS = "https://www.rizospastis.gr/rss.xml"
const RIZOSPASTIS_MAX = 50

type rizospastisRSS struct {
	Channel struct {
		Items []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
			Guid    string `xml:"guid"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchRizospastis() ([]model.Post, error) {

	resp, err := http.Get(RIZOSPASTIS_RSS)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rss rizospastisRSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	for i, item := range rss.Channel.Items {

		if i >= RIZOSPASTIS_MAX {
			break
		}

		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			t, _ = time.Parse(time.RFC1123, item.PubDate)
		}

		posts = append(posts, sources.NormalizeNews(
			item.Guid,
			item.Title,
			item.Link,
			"Rizospastis",
			t.Unix(),
		))
	}

	return posts, nil
}

func init() {
	sources.RegisterSource(sources.Source{
		Name:  "Rizospastis",
		Type:  "greece",
		Fetch: fetchRizospastis,
	})
}
