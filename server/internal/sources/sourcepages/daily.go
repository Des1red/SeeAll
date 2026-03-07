package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
	"SeeAll/internal/sources/rss"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Al Jazeera",
		Type: model.AudienceDaily,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.aljazeera.com/xml/rss/all.xml",
				"Al Jazeera",
				50,
				false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "BBC",
		Type: model.AudienceDaily,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://feeds.bbci.co.uk/news/world/rss.xml",
				"BBC",
				50,
				false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Guardian",
		Type: model.AudienceDaily,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.theguardian.com/world/rss",
				"Guardian",
				50,
				false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Reuters",
		Type: model.AudienceDaily,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.reutersagency.com/feed/?best-topics=world&post_type=best",
				"Reuters",
				50,
				false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "NPR",
		Type: model.AudienceDaily,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://feeds.npr.org/1001/rss.xml",
				"NPR",
				50,
				false,
			)
		},
	})
}
