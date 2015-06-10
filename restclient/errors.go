package restclient

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/docker/machine/log"
)

type ApiError struct {
	httpStatusCode int
	message        string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%d - %s", e.httpStatusCode, e.message)
}

func (e *ApiError) HttpStatusCode() int {
	return e.httpStatusCode
}

func (e *ApiError) Message() string {
	return e.message
}

type ErrorResponse struct {
	Type    string `json:"Type"`
	Message string `json:"Message"`
	// TODO Errors is missing by intention due to unclear meaning
}

func isError(response *http.Response, expectedStatus int, err error) error {
	if err != nil {
		return err
	}
	if response != nil {
		if response.StatusCode == expectedStatus {
			// we got a response with the expected HTTP status code, hence no error
			return nil
		}

		// extract the API's error message to be returned later
		errorResponse := ErrorResponse{}
		errorBody, ioError := ioutil.ReadAll(response.Body)
		if ioError != nil {
			return ioError
		}
		err := json.Unmarshal(errorBody, &errorResponse)
		if err != nil {
			log.Debug(err)
		}
		return &ApiError{response.StatusCode, errorResponse.Message}
	}
	return nil
}