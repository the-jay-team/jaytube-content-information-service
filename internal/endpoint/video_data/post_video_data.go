package video_data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/the-jay-team/jaytube-content-information-service/internal/configs"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/data"
	"net/http"
)

type VideoController struct {
	OpensearchClient *opensearch.Client
	Config           configs.OpenSearchConfig
}

func (controller *VideoController) UploadVideoData(ginContext *gin.Context) {
	var videoData data.UploadVideoData
	if ginContext.ShouldBindBodyWith(&videoData, binding.JSON) != nil {
		ginContext.JSON(http.StatusInternalServerError, "Malformed JSON in request body")
		return
	}
	videoDataBytes, _ := json.Marshal(videoData)

	request := opensearchapi.IndexRequest{
		Index:      controller.Config.VideoDataIndex,
		DocumentID: ginContext.Query("id"),
		Body:       bytes.NewReader(videoDataBytes),
	}
	insertResponse, _ := request.Do(context.Background(), controller.OpensearchClient)

	if insertResponse.StatusCode != 200 {
		ginContext.JSON(http.StatusInternalServerError,
			fmt.Sprintf("Could not index new document in OpenSearch: %s", insertResponse.Body))
	}
}
