package endpoint

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/the-jay-team/jaytube-content-information-service/internal/data_provider"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/data"
	"net/http"
)

type PatchVideoDataEndpoint struct {
	dataProvider data_provider.DataProvider
}

func NewPatchVideoDataEndpoint(datProvider data_provider.DataProvider) *PatchVideoDataEndpoint {
	return &PatchVideoDataEndpoint{datProvider}
}

func (endpoint *PatchVideoDataEndpoint) PatchVideoData(ginContext *gin.Context) {
	id, exists := ginContext.GetQuery("id")
	var videoData data.VideoDataPayload

	if ginContext.ShouldBindBodyWith(&videoData, binding.JSON) != nil {
		ginContext.JSON(http.StatusBadRequest, "Malformed JSON in request body")
		return
	}
	if !exists {
		ginContext.JSON(http.StatusBadRequest, "Missing Query: id")
		return
	}

	_, videoExists, _ := endpoint.dataProvider.GetVideoData(id)
	if !videoExists {
		ginContext.JSON(http.StatusNotFound, "")
		return
	}

	response, providerErr := endpoint.dataProvider.PatchVideoData(id, videoData)
	if providerErr != nil {
		ginContext.JSON(http.StatusInternalServerError,
			fmt.Sprintf("error happend in data provider %s", providerErr))
		return
	}

	ginContext.JSON(http.StatusOK, response)
}
