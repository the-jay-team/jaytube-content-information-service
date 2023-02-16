package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-content-information-service/internal/endpoint"
	"github.com/the-jay-team/jaytube-content-information-service/test/data_provider"
	"log"
)

func main() {
	server := gin.Default()
	dataProvider := data_provider.NewMockedDataProvider()
	postVideoDataEndpoint := endpoint.NewPostVideoDataEndpoint(dataProvider)
	getVideoDataEndpoint := endpoint.NewGetVideoDataEndpoint(dataProvider)
	deleteVideoDataEndpoint := endpoint.NewDeleteVideoEndpoint(dataProvider)
	patchVideoDataEndpoint := endpoint.NewPatchVideoDataEndpoint(dataProvider)

	server.POST("/video-data", postVideoDataEndpoint.PostVideoData)
	server.GET("/video-data", getVideoDataEndpoint.GetVideoData)
	server.DELETE("/video-data", deleteVideoDataEndpoint.DeleteVideoData)
	server.PATCH("/video-data", patchVideoDataEndpoint.PatchVideoData)

	ginStartError := server.Run(":8080")
	if ginStartError != nil {
		log.Fatalf("Could not Startup gin: %s", ginStartError)
	}
}
