package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

func init() {

	sources.RegisterSource(sources.Source{
		Name: "Reddit",
		Type: "live",
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/worldnews/.rss",
				"Reddit",
				50,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "RedditNew",
		Type: "live",
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/worldnews/new/.rss",
				"Reddit",
				50,
			)
		},
	})
}
