package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Reuters",
		Type: "daily",
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
