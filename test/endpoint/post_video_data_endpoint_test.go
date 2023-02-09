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

var expected = data.VideoDataResponse{
	Id:          "1",
	Title:       "Test",
	UploadDate:  "01.01.2012",
	Visibility:  data.Public,
	Description: "Test",
	Tags:        []string{"Test", "Test2"},
}

func TestPostVideoDataReturnsCorrectResponse(t *testing.T) {
	var testPayload, _ = json.Marshal(data.VideoDataPayload{
		Title:       "Test",
		UploadDate:  "01.01.2012",
		Visibility:  data.Public,
		Description: "Test",
		Tags:        []string{"Test", "Test2"},
	})

	record, testContext := test.GinTestSetup()
	testContext.Request, _ = http.NewRequest(http.MethodPost, "/video-data", bytes.NewReader(testPayload))

	testEndpoint := endpoint.NewPostVideoData(data_provider.NewMockedDataProvider())
	testEndpoint.PostVideoData(testContext)

	actualResponse := data.VideoDataResponse{}
	_ = json.NewDecoder(record.Body).Decode(&actualResponse)

	assert.EqualValues(t, expected, actualResponse)
}

func TestMalformedJsonPayload(t *testing.T) {
	var testPayload = []byte(`{"Title": "test"}`)

	record, testContext := test.GinTestSetup()
	testContext.Request, _ = http.NewRequest(http.MethodPost, "/video-data", bytes.NewReader(testPayload))

	testEndpoint := endpoint.NewPostVideoData(data_provider.NewMockedDataProvider())
	testEndpoint.PostVideoData(testContext)

	assert.Equal(t, http.StatusInternalServerError, record.Code)
}
