package data_provider

import (
	"errors"
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

func (provider *MockedDataProvider) GetVideoData(id string) (data.VideoDataResponse, error) {
	if id == "1" {
		return data.VideoDataResponse{
			Id:          "1",
			Title:       "Test",
			UploadDate:  "01.01.2012",
			Visibility:  data.Public,
			Description: "Test",
			Tags:        []string{"Test", "Test2"},
		}, nil
	}
	return data.VideoDataResponse{}, errors.New("ID does not exist")
}

func (provider *MockedDataProvider) DeleteVideoData(id string) (bool, error) {
	if id == "1" {
		return true, nil
	}
	return false, nil
}
