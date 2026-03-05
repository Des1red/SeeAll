package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Al Jazeera",
		Type: "daily",
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://www.aljazeera.com/xml/rss/all.xml",
				"Al Jazeera",
				50,
			)
		},
	})
}
