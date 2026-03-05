package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "NPR",
		Type: "daily",
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://feeds.npr.org/1001/rss.xml",
				"NPR",
				50,
			)
		},
	})
}
