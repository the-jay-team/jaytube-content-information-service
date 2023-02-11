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
	"net/url"
	"testing"
)

func TestMalformedJsonReturnsBadRequest(t *testing.T) {
	var testPayload, _ = json.Marshal(struct {
		Title string `json:"title" binding:"required"`
	}{Title: "Test"})
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodPatch, "/video-data", bytes.NewReader(testPayload))

	urlValues := url.Values{}
	urlValues.Add("id", "1")
	context.Request.URL.RawQuery = urlValues.Encode()

	testEndpoint := endpoint.NewPatchVideoDataEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.PatchVideoData(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}

func TestMissingIdReturnsBadRequest(t *testing.T) {
	var testPayload, _ = json.Marshal(data.VideoDataPayload{
		Title:       "test",
		UploadDate:  "01.01.2012",
		Visibility:  data.Public,
		Description: "Test",
		Tags:        []string{"Test", "Test2"},
	})
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodPatch, "/video-data", bytes.NewReader(testPayload))

	testEndpoint := endpoint.NewPatchVideoDataEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.PatchVideoData(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}

func TestNotExistsReturnsNotFound(t *testing.T) {
	var testPayload, _ = json.Marshal(data.VideoDataPayload{
		Title:       "test",
		UploadDate:  "01.01.2012",
		Visibility:  data.Public,
		Description: "Test",
		Tags:        []string{"Test", "Test2"},
	})
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodPatch, "/video-data", bytes.NewReader(testPayload))

	urlValues := url.Values{}
	urlValues.Add("id", "aadawd")
	context.Request.URL.RawQuery = urlValues.Encode()

	testEndpoint := endpoint.NewPatchVideoDataEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.PatchVideoData(context)

	assert.Equal(t, http.StatusNotFound, record.Code)
}

func TestReturnCorrectPatchResponse(t *testing.T) {
	expectedResponse := data.VideoDataResponse{
		Id:          "1",
		Title:       "test",
		UploadDate:  "01.01.2012",
		Visibility:  data.Public,
		Description: "Test",
		Tags:        []string{"Test", "Test2"},
	}
	var testPayload, _ = json.Marshal(data.VideoDataPayload{
		Title:       "test",
		UploadDate:  "01.01.2012",
		Visibility:  data.Public,
		Description: "Test",
		Tags:        []string{"Test", "Test2"},
	})
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodPatch, "/video-data", bytes.NewReader(testPayload))

	urlValues := url.Values{}
	urlValues.Add("id", "1")
	context.Request.URL.RawQuery = urlValues.Encode()

	testEndpoint := endpoint.NewPatchVideoDataEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.PatchVideoData(context)

	actualResponse := data.VideoDataResponse{}
	_ = json.NewDecoder(record.Body).Decode(&actualResponse)

	assert.EqualValues(t, expectedResponse, actualResponse)
	assert.Equal(t, http.StatusOK, record.Code)
}
