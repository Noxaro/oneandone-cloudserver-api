package oneandone_cloudserver_api

import (
	"bytes"
	"encoding/json"
	log "github.com/docker/machine/log"
	"io"
	"io/ioutil"
	"net/http"
)

type RestClient struct {
	token string
}

func NewRestClient(token string) *RestClient {
	restClient := new(RestClient)
	restClient.token = token
	return restClient
}

func (c *RestClient) Get(url string, result interface{}, expectedStatus int) (*http.Response, error) {
	return c.doRequest(url, "GET", nil, result, expectedStatus)
}

func (c *RestClient) Post(url string, requestBody interface{}, result interface{}, expectedStatus int) (*http.Response, error) {
	return c.doRequest(url, "POST", requestBody, result, expectedStatus)
}

func (c *RestClient) doRequest(url string, method string, requestBody interface{}, result interface{}, expectedStatus int) (*http.Response, error) {
	var bodyData io.Reader
	if requestBody != nil {
		data, _ := json.Marshal(requestBody)
		bodyData = bytes.NewBuffer(data)
	}

	request, err := http.NewRequest(method, url, bodyData)
	if err != nil {
		return nil, err
	}

	request.Header.Add("X-Token", c.token)
	request.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(response.Body)
	log.Debug(string(body[:]))
	response.Body.Close()
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return response, err
	}

	return response, nil
}

func logResult(response *http.Response, expectedStatus int) {
	if response != nil {
		log.Debug("response Status:", response.StatusCode)
		if response.StatusCode != expectedStatus {
			body, _ := ioutil.ReadAll(response.Body)
			log.Debug("response:", body)
			//log.Debug("response Headers:", response.HttpResponse().Header)
		}
	}
}

func isError(response *http.Response, expectedStatus int, err error) error {
	if err != nil {
		return err
	}
	//logResult(response, expectedStatus)
	if response != nil {
		if response.StatusCode == expectedStatus {
			// we got a response with the expected HTTP status code, hence no error
			return nil
		}
		println(response.StatusCode)
		// extract the API's error message to be returned later
		errorResponse := errorResponse{}
		body, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(body, errorResponse)
		//err := response.Unmarshal(&errorResponse)
		if err != nil {
			log.Debug(err)
		}

		switch response.StatusCode {
		case http.StatusUnauthorized:
			// TODO: hopefully all error will have a proper error object in the body later
			return &ApiError{response.StatusCode, "Authentication is failed, please check your X-TOKEN"}
		default:
			return &ApiError{response.StatusCode, errorResponse.Message}
		}

	} else {
		// no response from API means generic error
	}
	return nil
}
