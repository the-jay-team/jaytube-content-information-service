package main

import (
	"github.com/gin-gonic/gin"
	"github.com/the-jay-team/jaytube-content-information-service/interneal/clients"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/configs"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/controllers"
)

func main() {
	server := gin.Default()
	openSearchConfig := configs.GetEnvironmentConfig().OpenSearch
	videoController := controllers.NewVideoController(
		*clients.NewOpenSearchClient(
			openSearchConfig.Host,
			openSearchConfig.VideoDataIndex,
			openSearchConfig.Username,
			openSearchConfig.Password))

	server.POST("/video", videoController.PostVideoData)
	server.GET("/video", videoController.GetVideoById)
	server.DELETE("/video", videoController.DeleteVideoById)

	server.Run(":8080")
}
