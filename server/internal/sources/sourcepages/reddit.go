package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

func init() {

	sources.RegisterSource(sources.Source{
		Name: "Reddit",
		Type: model.AudienceLive,
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
		Type: model.AudienceLive,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/worldnews/new/.rss",
				"Reddit",
				50,
			)
		},
	})

	// TECH FEEDS
	sources.RegisterSource(sources.Source{
		Name: "RedditTech",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/technology/.rss",
				"Reddit",
				50,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "RedditProgramming",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/programming/.rss",
				"Reddit",
				50,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "RedditLinux",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/linux/.rss",
				"Reddit",
				50,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "RedditNetsec",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/netsec/.rss",
				"Reddit",
				50,
			)
		},
	})

	// SPORTS
	sources.RegisterSource(sources.Source{
		Name: "RedditSports",
		Type: model.AudienceSports,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/sports/.rss",
				"Reddit",
				50,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "RedditSoccer",
		Type: model.AudienceSports,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/soccer/.rss",
				"Reddit",
				50,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "RedditFormula1",
		Type: model.AudienceSports,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://www.reddit.com/r/formula1/.rss",
				"Reddit",
				50,
			)
		},
	})
}
