package models

type OpenSearchVideoData struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Tags        []string   `json:"tags"`
	Creator     string     `json:"creator"`
	Visibility  Visibility `json:"visibility"`
}
