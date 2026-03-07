package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
	"SeeAll/internal/sources/rss"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "MetaFilter",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.metafilter.com/rss.xml",
				"MetaFilter",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "BoingBoing",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://boingboing.net/feed",
				"BoingBoing",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Kottke",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://feeds.kottke.org/main",
				"Kottke",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "The Verge",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.theverge.com/rss/index.xml",
				"The Verge",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Ars Technica",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://feeds.arstechnica.com/arstechnica/index",
				"Ars Technica",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Wired",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.wired.com/feed/rss",
				"Wired",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Longreads",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://longreads.com/feed/",
				"Longreads",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "The Atlantic",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.theatlantic.com/feed/all/",
				"The Atlantic",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "The New Yorker",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.newyorker.com/feed/everything",
				"The New Yorker",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Slate",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://slate.com/feeds/all.rss",
				"Slate",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Vox",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.vox.com/rss/index.xml",
				"Vox",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Hakai Magazine",
		Type: model.AudienceGeneral,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://hakaimagazine.com/feed/",
				"Hakai Magazine",
				50, false,
			)
		},
	})
}
