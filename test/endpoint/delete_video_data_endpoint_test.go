package endpoint

import (
	"github.com/stretchr/testify/assert"
	"github.com/the-jay-team/jaytube-content-information-service/internal/endpoint"
	"github.com/the-jay-team/jaytube-content-information-service/test"
	"github.com/the-jay-team/jaytube-content-information-service/test/data_provider"
	"net/http"
	"net/url"
	"testing"
)

func TestReturnBadRequestMissingId(t *testing.T) {
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodDelete, "/video-data", nil)

	testEndpoint := endpoint.NewDeleteVideoEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.DeleteVideoData(context)

	assert.Equal(t, http.StatusBadRequest, record.Code)
}

func TestAlreadyDeleted(t *testing.T) {
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodDelete, "/video-data", nil)

	urlValues := url.Values{}
	urlValues.Add("id", "2")
	context.Request.URL.RawQuery = urlValues.Encode()

	testEndpoint := endpoint.NewDeleteVideoEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.DeleteVideoData(context)

	assert.Equal(t, http.StatusGone, record.Code)
}

func TestReturnOk(t *testing.T) {
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodDelete, "/video-data", nil)

	urlValues := url.Values{}
	urlValues.Add("id", "1")
	context.Request.URL.RawQuery = urlValues.Encode()

	testEndpoint := endpoint.NewDeleteVideoEndpoint(data_provider.NewMockedDataProvider())
	testEndpoint.DeleteVideoData(context)

	assert.Equal(t, http.StatusOK, record.Code)
}
