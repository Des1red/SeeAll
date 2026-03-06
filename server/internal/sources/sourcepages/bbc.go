package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "BBC",
		Type: "daily",
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
