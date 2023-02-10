package data_provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	data2 "github.com/the-jay-team/jaytube-content-information-service/internal/data"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/data"
)

type OpensearchDataProvider struct {
	client *opensearch.Client
	index  string
}

func NewDataProvider(client *opensearch.Client, index string) *OpensearchDataProvider {
	provider := &OpensearchDataProvider{
		client,
		index,
	}
	return provider
}

func (provider *OpensearchDataProvider) PostVideoData(payload data.VideoDataPayload) (data.VideoDataResponse, error) {
	dataBytes, _ := json.Marshal(payload)

	request := opensearchapi.IndexRequest{
		Index: provider.index,
		Body:  bytes.NewReader(dataBytes)}
	opensearchResponse, _ := request.Do(context.Background(), provider.client)
	if opensearchResponse.StatusCode != 201 {
		return data.VideoDataResponse{}, errors.New(fmt.Sprintf("failed to post video data to opensearch: %d",
			opensearchResponse.StatusCode))
	}

	response := struct {
		Id string `json:"_id"`
	}{}
	if json.NewDecoder(opensearchResponse.Body).Decode(&response) != nil {
		return data.VideoDataResponse{}, errors.New("could not decode opensearch response")
	}

	return data.VideoDataResponse{
		Id:          response.Id,
		Title:       payload.Title,
		Description: payload.Description,
		UploadDate:  payload.UploadDate,
		Tags:        payload.Tags,
		Visibility:  payload.Visibility}, nil
}

func (provider *OpensearchDataProvider) GetVideoData(id string) (data.VideoDataResponse, error) {
	request := opensearchapi.GetRequest{
		Index:      provider.index,
		DocumentID: id}
	opensearchResponse, _ := request.Do(context.Background(), provider.client)
	if opensearchResponse.StatusCode != 200 {
		return data.VideoDataResponse{}, errors.New("failed to get video data to opensearch")
	}

	response := data2.OpenSearchVideoData{}
	if json.NewDecoder(opensearchResponse.Body).Decode(&response) != nil {
		return data.VideoDataResponse{}, errors.New("could not decode opensearch response")
	}
	return data.VideoDataResponse{
		Id:          response.Id,
		Title:       response.Source.Title,
		Description: response.Source.Title,
		UploadDate:  response.Source.UploadDate,
		Tags:        response.Source.Tags,
		Visibility:  response.Source.Visibility}, nil
}
