package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
	"SeeAll/internal/sources/rss"
)

/* ==============================
   ERT NEWS
============================== */

func init() {
	sources.RegisterSource(sources.Source{
		Name: "ERT News",
		Type: model.AudienceGreece,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.ertnews.gr/feed/",
				"ERT News",
				50,
				false,
			)
		},
	})
}

/* ==============================
   NAFTEBORIKI
============================== */

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Naftemporiki",
		Type: model.AudienceGreece,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.naftemporiki.gr/rss",
				"Naftemporiki",
				50,
				false,
			)
		},
	})
}

/* ==============================
   902
============================== */

func init() {
	sources.RegisterSource(sources.Source{
		Name: "902",
		Type: model.AudienceGreece,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.902.gr/feed/recent",
				"902",
				50,
				true,
			)
		},
	})
}
