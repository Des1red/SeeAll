package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Al Jazeera",
		Type: model.AudienceDaily,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://www.aljazeera.com/xml/rss/all.xml",
				"Al Jazeera",
				50,
				false,
			)
		},
	})
}

func init() {
	sources.RegisterSource(sources.Source{
		Name: "BBC",
		Type: model.AudienceDaily,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://feeds.bbci.co.uk/news/world/rss.xml",
				"BBC",
				50,
				false,
			)
		},
	})
}

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Guardian",
		Type: model.AudienceDaily,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://www.theguardian.com/world/rss",
				"Guardian",
				50,
				false,
			)
		},
	})
}

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Reuters",
		Type: model.AudienceDaily,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://www.reutersagency.com/feed/?best-topics=world&post_type=best",
				"Reuters",
				50,
				false,
			)
		},
	})
}
