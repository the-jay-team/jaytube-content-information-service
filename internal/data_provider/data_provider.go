package data_provider

import "github.com/the-jay-team/jaytube-content-information-service/pkg/data"

type DataProvider interface {
	PostVideoData(payload data.VideoDataPayload) (data.VideoDataResponse, error)

	GetVideoData(id string) (data.VideoDataResponse, bool, error)

	DeleteVideoData(id string) (bool, error)

	PatchVideoData(id string, payload data.VideoDataPayload) (data.VideoDataResponse, error)
}
