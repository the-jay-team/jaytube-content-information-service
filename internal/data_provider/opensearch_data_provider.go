package data_provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
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
		Body:  bytes.NewReader(dataBytes),
	}

	insertResponse, _ := request.Do(context.Background(), provider.client)
	if insertResponse.StatusCode != 200 {
		return data.VideoDataResponse{}, errors.New("failed to post video data to opensearch")
	}

	response := data.VideoDataResponse{}
	if json.NewDecoder(insertResponse.Body).Decode(&response) != nil {
		return data.VideoDataResponse{}, errors.New("could not decode opensearch response")
	}

	return response, nil
}
