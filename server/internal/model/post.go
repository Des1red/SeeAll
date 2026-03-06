package model

type Post struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Image  string `json:"image,omitempty"`
	Source string `json:"source"`
	Time   int64  `json:"time"`
	Score  *int   `json:"score"`
}

const (
	AudienceDaily   = "daily"
	AudienceGeneral = "general"
	AudienceGreece  = "greece"
	AudienceTech    = "tech"
	AudienceSports  = "sports"
)
