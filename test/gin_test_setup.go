package test

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

func GinTestSetup() (*httptest.ResponseRecorder, *gin.Context) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(recorder)

	return recorder, testContext
}
