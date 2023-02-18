package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/data"
	"io"
	"net/http"
)

type ContentInformationServiceClient struct {
	target string
}

func NewContentInformationServiceClient(target string) *ContentInformationServiceClient {
	return &ContentInformationServiceClient{target}
}

func (client *ContentInformationServiceClient) VideoExists(id string) (bool, error) {
	response, err := http.Get(fmt.Sprintf("%s/video-data?id=%s", client.target, id))
	if err != nil {
		return false, err
	}

	body, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return false, errors.New(fmt.Sprintf("request failed with code %d: %s",
			response.StatusCode, string(body)))
	}

	return true, nil
}

func (client *ContentInformationServiceClient) GetVideoData(id string) (data.VideoDataResponse, error) {
	response, err := http.Get(fmt.Sprintf("%s/video-data?id=%s", client.target, id))
	if err != nil {
		return data.VideoDataResponse{}, err
	}

	body, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return data.VideoDataResponse{}, errors.New(fmt.Sprintf("request failed with code %d: %s",
			response.StatusCode, string(body)))
	}

	videoDate := data.VideoDataResponse{}
	_ = json.Unmarshal(body, &videoDate)

	return videoDate, nil
}

func (client *ContentInformationServiceClient) DeleteVideoData(id string) error {
	request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/video-data?id=%s", client.target, id), nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("could not delete video data via iris [%d]: %s",
			response.StatusCode, string(body)))
	}
	return nil
}
