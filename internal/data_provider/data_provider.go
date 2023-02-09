package data_provider

import "github.com/the-jay-team/jaytube-content-information-service/pkg/data"

type DataProvider interface {
	PostVideoData(payload data.VideoDataPayload) (data.VideoDataResponse, error)
}
