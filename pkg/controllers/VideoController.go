package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/the-jay-team/jaytube-content-information-service/interneal/clients"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/models"
	"net/http"
)

type VideoController struct {
	openSearch clients.OpenSearchClient
}

func NewVideoController(openSearch clients.OpenSearchClient) *VideoController {
	controller := &VideoController{openSearch}
	return controller
}

func (controller *VideoController) PostVideoData(context *gin.Context) {
	var data models.VideoData

	if context.ShouldBindBodyWith(&data, binding.JSON) != nil {
		context.JSON(http.StatusInternalServerError, "Malformed JSON in request body")
		return
	}
	openSearchData := models.OpenSearchVideoData{
		Title:       data.Title,
		Description: data.Description,
		Tags:        data.Tags,
		Creator:     data.Creator,
		Visibility:  data.Visibility}

	openSearchError := controller.openSearch.UpdateVideoData(data.Id, openSearchData)
	if openSearchError != nil {
		context.JSON(http.StatusInternalServerError,
			fmt.Sprintf("error while updating video data: %s", openSearchError))
	}
}

func (controller *VideoController) GetVideoById(context *gin.Context) {
	id := context.Query("id")

	if controller.openSearch.DoesVideoDataExistsForId(id) {
		context.JSON(http.StatusBadRequest, "No VideoData found in OpenSearch")
		return
	}

	openSearchData, openSearchError := controller.openSearch.GetVideoDataById(id)
	if openSearchError != nil {
		context.JSON(http.StatusInternalServerError, openSearchError.Error())
		return
	}

	videoData := models.VideoData{
		Id:          id,
		Title:       openSearchData.Title,
		Description: openSearchData.Description,
		UploadDate:  openSearchData.UploadDate,
		Tags:        openSearchData.Tags,
		Creator:     openSearchData.Creator,
		Visibility:  openSearchData.Visibility,
	}

	context.JSON(http.StatusOK, videoData)
}

func (controller *VideoController) DeleteVideoById(context *gin.Context) {
	id := context.Query("id")
	if controller.openSearch.DoesVideoDataExistsForId(id) {
		context.JSON(http.StatusBadRequest, "No VideoData found in OpenSearch")
		return
	}

	err := controller.openSearch.DeleteVideoDataById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}
}
