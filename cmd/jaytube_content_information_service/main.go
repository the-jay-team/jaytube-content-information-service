package main

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/opensearch-project/opensearch-go"
	"github.com/the-jay-team/jaytube-content-information-service/internal/configs"
	"github.com/the-jay-team/jaytube-content-information-service/internal/endpoint/video_data"
	"log"
	"net/http"
)

func main() {
	openSearchConfig := configs.GetEnvironmentConfig().Opensearch
	client, opensearchClientError := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{openSearchConfig.Host},
		Username:  openSearchConfig.Username,
		Password:  openSearchConfig.Password,
	})
	if opensearchClientError != nil {
		log.Fatalf("Could not initialize OpenSearch Client: %s", opensearchClientError)
	}

	server := gin.Default()
	videoController := video_data.VideoController{
		OpensearchClient: client,
		Config:           openSearchConfig,
	}

	server.POST("/video", videoController.UploadVideoData)

	ginStartError := server.Run(":8080")
	if ginStartError != nil {
		log.Fatal("Could not Startup gin: ")
	}
}
