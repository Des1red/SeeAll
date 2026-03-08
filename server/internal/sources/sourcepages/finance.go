package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
	"SeeAll/internal/sources/rss"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "New Money",
		Type: model.AudienceFinance,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.newmoney.gr/feed/",
				"New Money",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Wall Street Journal Markets",
		Type: model.AudienceFinance,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://feeds.a.dj.com/rss/RSSMarketsMain.xml",
				"Wall Street Journal Markets",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Financial Times",
		Type: model.AudienceFinance,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.ft.com/rss/home/international",
				"Financial Times",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "CNBC Finance",
		Type: model.AudienceFinance,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://search.cnbc.com/rs/search/combinedcms/view.xml?partnerId=wrss01&id=10000664",
				"CNBC Finance",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "The Economist",
		Type: model.AudienceFinance,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://www.economist.com/finance-and-economics/rss.xml",
				"The Economist",
				50, false,
			)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "MarketWatch",
		Type: model.AudienceFinance,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS(
				"https://feeds.content.dowjones.io/public/rss/mw_topstories",
				"MarketWatch",
				50, false,
			)
		},
	})
}
