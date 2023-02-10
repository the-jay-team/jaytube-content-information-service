package endpoint

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/the-jay-team/jaytube-content-information-service/internal/data_provider"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/data"
	"net/http"
)

type PostVideoDataEndpoint struct {
	dataProvider data_provider.DataProvider
}

func NewPostVideoDataEndpoint(dataProvider data_provider.DataProvider) *PostVideoDataEndpoint {
	return &PostVideoDataEndpoint{dataProvider}
}

func (endpoint *PostVideoDataEndpoint) PostVideoData(ginContext *gin.Context) {
	var videoData data.VideoDataPayload
	if ginContext.ShouldBindBodyWith(&videoData, binding.JSON) != nil {
		ginContext.JSON(http.StatusBadRequest, "Malformed JSON in request body")
		return
	}

	response, providerErr := endpoint.dataProvider.PostVideoData(videoData)
	if providerErr != nil {
		ginContext.JSON(http.StatusInternalServerError,
			fmt.Sprintf("error happened in data provider: %s", providerErr))
		return
	}

	ginContext.JSON(http.StatusOK, response)
}
