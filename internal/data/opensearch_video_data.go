package data

type OpenSearchVideoData struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UploadDate  string     `json:"uploadDate"`
	Tags        []string   `json:"tags"`
	Creator     string     `json:"creator"`
	Visibility  Visibility `json:"visibility"`
}
