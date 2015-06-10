package restclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"github.com/docker/machine/log"
	"bytes"
)

const (
	Get    string = "GET"
	Post   string = "POST"
	Put    string = "PUT"
	Delete string = "DELETE"
)

type RestClient struct {
	http     http.Client
	xToken   string
	endpoint string
}

func New(endpoint string, token string) *RestClient {
	restClient := new(RestClient)
	restClient.xToken = token
	restClient.http = http.Client{}
	restClient.endpoint = endpoint
	return restClient
}

func (c RestClient) CreateUrl(sections ...interface{}) string {
	url := c.endpoint
	for _, section := range sections {
		url += "/" + fmt.Sprint(section)
	}
	return url
}

func (c RestClient) Get(url string, result interface{}, expectedHttpStatus int) error {
	return doRequest(&c, url, Get, nil, result, expectedHttpStatus)
}

func (c *RestClient) Delete(url string, result interface{}, expectedHttpStatus int) error {
	return doRequest(c, url, Delete, nil, result, expectedHttpStatus)
}

func (c RestClient) Post(url string, content interface{}, result interface{}, expectedHttpStatus int) error {
	return doRequest(&c, url, Post, content, result, expectedHttpStatus)
}

func (c RestClient) Put(url string, content interface{}, result interface{}, expectedHttpStatus int) error {
	return doRequest(&c, url, Put, content, result, expectedHttpStatus)
}

func doRequest(c *RestClient, url string, requestType string, content interface{}, result interface{}, expectedHttpStatus int) error {
	//Request part
	var requestError error
	var request *http.Request

	if content != nil {
		jsonContent, jsonEncodeError := json.Marshal(content)
		if jsonEncodeError != nil {
			return jsonEncodeError
		}
		request, requestError = http.NewRequest(requestType, url, bytes.NewBuffer(jsonContent))
	} else {
		request, requestError = http.NewRequest(requestType, url, nil)
	}

	if requestError != nil {
		return requestError
	}
	request.Header.Add("X-TOKEN", c.xToken)
	request.Header.Add("Content-Type", "application/json")

	requestResult, httpError := c.http.Do(request)
	if httpError != nil {
		return httpError
	}

	// Response part
	if httpError = isError(requestResult, expectedHttpStatus, httpError); httpError != nil {
		return httpError
	}

	body, ioError := ioutil.ReadAll(requestResult.Body)
	requestResult.Body.Close()
	if ioError != nil {
		return ioError
	}
	jsonDecodeError := json.Unmarshal(body, result)
	if jsonDecodeError != nil {
		return jsonDecodeError
	}

	return nil
}
