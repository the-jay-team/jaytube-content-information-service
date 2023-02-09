package data_provider

import (
	"github.com/the-jay-team/jaytube-content-information-service/pkg/data"
)

type MockedDataProvider struct {
}

func NewMockedDataProvider() *MockedDataProvider {
	return &MockedDataProvider{}
}

func (provider *MockedDataProvider) PostVideoData(payload data.VideoDataPayload) (data.VideoDataResponse, error) {
	return data.VideoDataResponse{
		Id:          "1",
		Tags:        payload.Tags,
		Title:       payload.Title,
		Description: payload.Description,
		Visibility:  payload.Visibility,
		UploadDate:  payload.UploadDate,
	}, nil
}
