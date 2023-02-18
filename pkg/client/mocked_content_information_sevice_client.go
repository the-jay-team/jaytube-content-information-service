package client

import (
	"github.com/the-jay-team/jaytube-content-information-service/pkg/data"
)

type MockedContentInformationServiceClient struct {
}

func NewMockedContentInformationServiceClient() *MockedContentInformationServiceClient {
	return &MockedContentInformationServiceClient{}
}

func (client *MockedContentInformationServiceClient) VideoExists(id string) (bool, error) {
	return id == "1", nil
}

func (client *MockedContentInformationServiceClient) GetVideoData(id string) (data.VideoDataResponse, error) {
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
	return data.VideoDataResponse{}, nil
}

func (client *MockedContentInformationServiceClient) DeleteVideoData(id string) (bool, error) {
	if id == "1" {
		return true, nil
	}

	return false, nil
}
