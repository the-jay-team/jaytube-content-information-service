package client

import "github.com/the-jay-team/jaytube-content-information-service/pkg/data"

type IrisClient interface {
	VideoExists(id string) (bool, error)

	GetVideoData(id string) (data.VideoDataResponse, error)
}
