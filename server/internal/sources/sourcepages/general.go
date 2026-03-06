package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "MetaFilter",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://www.metafilter.com/rss.xml",
				"MetaFilter",
				50,
				false,
			)
		},
	})
}
