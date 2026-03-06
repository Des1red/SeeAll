package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Lobsters",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchJSONRSS(
				"https://api.rss2json.com/v1/api.json?rss_url=https://lobste.rs/rss",
				"Lobsters",
				50,
			)
		},
	})
}
