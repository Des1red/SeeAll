package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
	"SeeAll/internal/sources/rss"
)

func init() {

	sources.RegisterSource(sources.Source{
		Name: "Reddit",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/worldnews/.rss",
				"Reddit",
				50,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "RedditNew",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/worldnews/new/.rss",
				"Reddit",
				50,
			)
		},
	})
}
