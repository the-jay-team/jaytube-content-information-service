package data

import "github.com/the-jay-team/jaytube-content-information-service/pkg/data"

type OpenSearchVideoData struct {
	Id     string                `json:"_id"`
	Source data.VideoDataPayload `json:"_source"`
}
