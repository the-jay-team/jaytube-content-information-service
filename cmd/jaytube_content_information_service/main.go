package main

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/opensearch-project/opensearch-go"
	"github.com/the-jay-team/jaytube-content-information-service/internal/configs"
	"github.com/the-jay-team/jaytube-content-information-service/internal/data_provider"
	"github.com/the-jay-team/jaytube-content-information-service/internal/endpoint"
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
	dataProvider := data_provider.NewDataProvider(client, openSearchConfig.VideoDataIndex)
	postVideoDataEndpoint := endpoint.NewPostVideoDataEndpoint(dataProvider)
	getVideoDataEndpoint := endpoint.NewGetVideoDataEndpoint(dataProvider)
	deleteVideoDataEndpoint := endpoint.NewDeleteVideoEndpoint(dataProvider)

	server.POST("/video-data", postVideoDataEndpoint.PostVideoData)
	server.GET("/video-data", getVideoDataEndpoint.GetVideoData)
	server.DELETE("/video-data", deleteVideoDataEndpoint.DeleteVideoData)

	ginStartError := server.Run(":8080")
	if ginStartError != nil {
		log.Fatalf("Could not Startup gin: %s", ginStartError)
	}
}
