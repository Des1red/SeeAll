package rss

import (
	"SeeAll/internal/model"
	"encoding/xml"
	"io"
	"net/http"
)

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
