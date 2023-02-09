package data

type VideoDataResponse struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UploadDate  string     `json:"uploadDate"`
	Tags        []string   `json:"tags"`
	Visibility  Visibility `json:"visibility"`
}
