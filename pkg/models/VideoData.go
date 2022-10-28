package models

type Visibility string

const (
	Public    Visibility = "PUBLIC"
	NotListed            = "NOT_LISTED"
	Private              = "PRIVATE"
)

type VideoData struct {
	Id          string     `json:"id" binding:"required"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Tags        []string   `json:"tags"`
	Creator     string     `json:"creator"`
	Visibility  Visibility `json:"visibility"`
}
