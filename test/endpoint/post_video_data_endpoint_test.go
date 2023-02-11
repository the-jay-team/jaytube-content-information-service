package endpoint

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/the-jay-team/jaytube-content-information-service/internal/endpoint"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/data"
	"github.com/the-jay-team/jaytube-content-information-service/test"
	"github.com/the-jay-team/jaytube-content-information-service/test/data_provider"
	"net/http"
	"testing"
)

func TestPostVideoDataReturnsCorrectResponse(t *testing.T) {
	expectedResponse := data.VideoDataResponse{
		Id:          "1",
		Title:       "Test",
		UploadDate:  "01.01.2012",
		Visibility:  data.Public,
		Description: "Test",
		Tags:        []string{"Test", "Test2"},
	}
	var testPayload, _ = json.Marshal(data.VideoDataPayload{
		Title:       "Test",
		UploadDate:  "01.01.2012",
		Visibility:  data.Public,
		Description: "Test",
		Tags:        []string{"Test", "Test2"},
	})
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodPost, "/video-data", bytes.NewReader(testPayload))

	testEndpoint := endpoint.NewPostVideoDataEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.PostVideoData(context)

	actualResponse := data.VideoDataResponse{}
	_ = json.NewDecoder(record.Body).Decode(&actualResponse)

	assert.EqualValues(t, expectedResponse, actualResponse)
}

func TestMalformedJsonPayload(t *testing.T) {
	var testPayload = []byte(`{"Title": "test"}`)

	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodPost, "/video-data", bytes.NewReader(testPayload))

	testEndpoint := endpoint.NewPostVideoDataEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.PostVideoData(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}
