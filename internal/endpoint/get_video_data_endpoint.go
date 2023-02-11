package endpoint

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-content-information-service/internal/data_provider"
	"net/http"
)

type GetVideoDataEndpoint struct {
	dataProvider data_provider.DataProvider
}

func NewGetVideoDataEndpoint(dataProvider data_provider.DataProvider) *GetVideoDataEndpoint {
	return &GetVideoDataEndpoint{dataProvider}
}

func (endpoint *GetVideoDataEndpoint) GetVideoData(ginContext *gin.Context) {
	id, exists := ginContext.GetQuery("id")
	if !exists {
		ginContext.JSON(http.StatusBadRequest, "Missing Query: id")
		return
	}

	response, exists, providerErr := endpoint.dataProvider.GetVideoData(id)
	if providerErr != nil {
		ginContext.JSON(http.StatusInternalServerError,
			fmt.Sprintf("error happened in data provider: %s", providerErr))
		return
	}
	if !exists {
		ginContext.JSON(http.StatusNotFound, "")
		return
	}
	ginContext.JSON(http.StatusOK, response)
}
