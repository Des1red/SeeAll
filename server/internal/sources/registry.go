package sources

import "SeeAll/internal/model"

type Source struct {
	Name  string
	Type  string
	Fetch func() ([]model.Post, error)
}

var sources []Source

func RegisterSource(s Source) {
	sources = append(sources, s)
}

func GetSources() []Source {
	return sources
}
