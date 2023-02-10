package endpoint

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/the-jay-team/jaytube-content-information-service/internal/endpoint"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/data"
	"github.com/the-jay-team/jaytube-content-information-service/test"
	"github.com/the-jay-team/jaytube-content-information-service/test/data_provider"
	"net/http"
	"net/url"
	"testing"
)

func TestReturnCorrectData(t *testing.T) {
	expectedResponse := data.VideoDataResponse{
		Id:          "1",
		Title:       "Test",
		UploadDate:  "01.01.2012",
		Visibility:  data.Public,
		Description: "Test",
		Tags:        []string{"Test", "Test2"},
	}
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodGet, "/video-data", nil)

	urlValues := url.Values{}
	urlValues.Add("id", "1")
	context.Request.URL.RawQuery = urlValues.Encode()

	testEndpoint := endpoint.NewGetVideoDataEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.GetVideoData(context)

	actualResponse := data.VideoDataResponse{}
	_ = json.NewDecoder(record.Body).Decode(&actualResponse)

	assert.EqualValues(t, actualResponse, expectedResponse)
}

func TestReturnBadRequestOnMissingId(t *testing.T) {
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodGet, "/video-data", nil)

	testEndpoint := endpoint.NewGetVideoDataEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.GetVideoData(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}
