package model

type Post struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Source string `json:"source"`
	Time   int64  `json:"time"`
	Score  *int   `json:"score"`
}
