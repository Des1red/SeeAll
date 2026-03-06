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
				false,
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
		Type: "greece",
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
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
		Type: "greece",
		Fetch: func() ([]model.Post, error) {
			return sources.FetchRSS(
				"https://www.902.gr/rss",
				"902",
				50,
				true, // requires browser UA
			)
		},
	})
}
