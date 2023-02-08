package data

type Visibility string

const (
	Public    Visibility = "PUBLIC"
	NotListed            = "NOT_LISTED"
	Private              = "PRIVATE"
)

type UploadVideoData struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UploadDate  string     `json:"uploadDate"`
	Tags        []string   `json:"tags"`
	Visibility  Visibility `json:"visibility"`
}
