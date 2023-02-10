package endpoint

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-content-information-service/internal/data_provider"
	"net/http"
)

type DeleteVideoDataEndpoint struct {
	dataProvider data_provider.DataProvider
}

func NewDeleteVideoEndpoint(dataProvider data_provider.DataProvider) *DeleteVideoDataEndpoint {
	return &DeleteVideoDataEndpoint{dataProvider}
}

func (endpoint *DeleteVideoDataEndpoint) DeleteVideoData(ginContext *gin.Context) {
	id, exists := ginContext.GetQuery("id")
	if !exists {
		ginContext.JSON(http.StatusBadRequest, "Missing Querry: id")
		return
	}

	existed, providerErr := endpoint.dataProvider.DeleteVideoData(id)
	if providerErr != nil {
		ginContext.JSON(http.StatusInternalServerError,
			fmt.Sprintf("error happened in data provider: %s", providerErr))
		return
	}

	if !existed {
		ginContext.JSON(http.StatusGone, "")
		return
	}
	ginContext.JSON(http.StatusOK, "")
}
