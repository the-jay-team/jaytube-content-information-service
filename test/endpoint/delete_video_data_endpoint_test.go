package endpoint

import (
	"github.com/stretchr/testify/assert"
	"github.com/the-jay-team/jaytube-content-information-service/test"
	"net/http"
	"net/url"
	"testing"
)

func TestReturnBadRequestMissingId(t *testing.T) {
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodDelete, "/video-data", nil)

	// TODO call endpoint

	assert.Equal(t, http.StatusBadRequest, record.Code)
}

func TestReturnOk(t *testing.T) {
	record, context := test.GinTestSetup()
	context.Request, _ = http.NewRequest(http.MethodDelete, "/video-data", nil)

	urlValues := url.Values{}
	urlValues.Add("id", "abc45")
	context.Request.URL.RawQuery = urlValues.Encode()

	// TODO call endpoint

	assert.Equal(t, http.StatusOK, record.Code)
}
