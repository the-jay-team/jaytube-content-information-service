package clients

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/the-jay-team/jaytube-content-information-service/pkg/models"
	"io"
	"log"
	"net/http"
	"strings"
)

type OpenSearchClient struct {
	hostUrl    string
	username   string
	password   string
	httpClient *http.Client
}

func NewOpenSearchClient(hostUrl string, index string, username string, password string) *OpenSearchClient {
	if !strings.HasSuffix(hostUrl, "/") {
		hostUrl += "/"
	}
	hostUrl += index
	log.Println(hostUrl)
	client := &OpenSearchClient{hostUrl, username, password, createHttpClient()}
	return client
}

func createHttpClient() *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient := &http.Client{Transport: transport}
	return httpClient
}

func (client OpenSearchClient) UpdateVideoData(id string, data models.OpenSearchVideoData) error {
	body := bytes.Buffer{}
	bodyEncoderError := json.NewEncoder(&body).Encode(data)
	if bodyEncoderError != nil {
		return bodyEncoderError
	}

	request := client.createRequest(http.MethodPut, id, &body)
	response, responseError := client.httpClient.Do(request)
	if responseError != nil {
		return responseError
	}

	buffer, _ := io.ReadAll(response.Body)
	log.Println(string(buffer))
	log.Println(response.StatusCode)

	return nil
}

func (client OpenSearchClient) GetVideoDataById(id string) (models.OpenSearchVideoData, error) {
	request := client.createRequest(http.MethodGet, id, nil)
	response, responseError := client.httpClient.Do(request)

	if response.StatusCode == http.StatusNotFound {
		return models.OpenSearchVideoData{}, errors.New("no VideoData found")
	}
	if responseError != nil {
		return models.OpenSearchVideoData{}, responseError
	}

	responseBytes, _ := io.ReadAll(response.Body)
	responseMap := map[string]any{}
	responseStruct := models.OpenSearchVideoData{}

	unmarshalError := json.Unmarshal(responseBytes, &responseMap)
	if unmarshalError != nil {
		return models.OpenSearchVideoData{}, unmarshalError
	}
	mappingError := mapstructure.Decode(responseMap["_source"], &responseStruct)
	if mappingError != nil {
		return models.OpenSearchVideoData{}, mappingError
	}

	return responseStruct, nil
}

func (client OpenSearchClient) DeleteVideoDataById(id string) error {
	request := client.createRequest(http.MethodDelete, id, nil)
	_, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}
	return nil
}

func (client OpenSearchClient) DoesVideoDataExistsForId(id string) bool {
	request := client.createRequest(http.MethodHead, id, nil)
	response, err := client.httpClient.Do(request)
	if err != nil {
		return false
	}
	return response.StatusCode != http.StatusOK
}

func (client OpenSearchClient) createRequest(methode string, documentId string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(methode,
		fmt.Sprintf("%s/_doc/%s/", client.hostUrl, documentId), body)
	req.SetBasicAuth(client.username, client.password)
	req.Header.Add("Content-Type", "application/json")
	return req
}
