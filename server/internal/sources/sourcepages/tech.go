package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
	"SeeAll/internal/sources/rss"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "ArsTechnica",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://feeds.arstechnica.com/arstechnica/index",
				"ArsTechnica",
				50,
				false,
			)
		},
	})
}

func init() {
	sources.RegisterSource(sources.Source{
		Name: "TechCrunch",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://techcrunch.com/feed/",
				"TechCrunch",
				50,
				false,
			)
		},
	})
}

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Krebs",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://krebsonsecurity.com/feed/",
				"Krebs",
				50,
				false,
			)
		},
	})
}

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Lobsters",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://lobste.rs/rss",
				"Lobsters",
				50,
			)
		},
	})
}

func init() {
	sources.RegisterSource(sources.Source{
		Name: "DevTo",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://dev.to/feed",
				"DevTo",
				50,
				false,
			)
		},
	})
}
