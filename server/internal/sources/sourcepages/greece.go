package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
)

/* ==============================
   ERT NEWS
============================== */

func init() {
	sources.RegisterSource(sources.Source{
		Name: "ERT News",
		Type: "greece",
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://www.ertnews.gr/feed/",
				"ERT News",
				50,
			)
		},
	})
}

/* ==============================
   RIZOSPASTIS
============================== */

func init() {
	sources.RegisterSource(sources.Source{
		Name: "Rizospastis",
		Type: "greece",
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://www.rizospastis.gr/rss.xml",
				"Rizospastis",
				50,
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
		Type: "greece",
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://www.naftemporiki.gr/rss",
				"Naftemporiki",
				50,
			)
		},
	})
}
