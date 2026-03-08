package sourcepages

import (
	"SeeAll/internal/model"
	"SeeAll/internal/sources"
	"SeeAll/internal/sources/rss"
)

func init() {
	sources.RegisterSource(sources.Source{
		Name: "ArsTechnica",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS("https://feeds.arstechnica.com/arstechnica/index", "ArsTechnica", 50, false)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "TechCrunch",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS("https://techcrunch.com/feed/", "TechCrunch", 50, false)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Krebs",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS("https://krebsonsecurity.com/feed/", "Krebs", 50, false)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "DevTo",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS("https://dev.to/feed", "DevTo", 50, false)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "HackersNews",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS("https://feeds.feedburner.com/TheHackersNews", "HackersNews", 50, false)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "TheVerge",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS("https://www.theverge.com/rss/index.xml", "TheVerge", 50, false)
		},
	})

	sources.RegisterSource(sources.Source{
		Name: "Thenewstack",
		Type: model.AudienceTech,
		Fetch: func() ([]model.Post, error) {
			return rss.FetchRSS("https://thenewstack.io/feed/", "Thenewstack", 50, false)
		},
	})
}
