package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
	"SeeAll/internal/sources/rss"
)

func init() {

	sources.RegisterSource(sources.Source{
		Name: "BBCSport",
		Type: model.AudienceSports,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://feeds.bbci.co.uk/sport/rss.xml",
				"BBCSport",
				50,
				false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "BBCFootball",
		Type: model.AudienceSports,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://feeds.bbci.co.uk/sport/football/rss.xml",
				"BBCFootball",
				50,
				false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "GuardianSport",
		Type: model.AudienceSports,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.theguardian.com/sport/rss",
				"GuardianSport",
				50,
				false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Gazzetta",
		Type: model.AudienceSports,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.gazzetta.gr/rss",
				"Gazzetta",
				50,
				false,
			)
		},
	})

}
