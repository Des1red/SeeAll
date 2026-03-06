package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

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
