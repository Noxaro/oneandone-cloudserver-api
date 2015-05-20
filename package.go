package oneandone_cloudserver_api

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jmcvetta/napping"
	"net/http"
)

type API struct {
	AuthToken string
	Endpoint  string
}

type withApi struct {
	api *API
}

type withId struct {
	Id string `json:"id"`
}

type withType struct {
	Type string `json:"type"`
}

type withName struct {
	Name string `json:"name"`
}

type withDescription struct {
	Description string `json:"description"`
}

type Status struct {
	State   string `json:"state"`
	Percent string `json:"percent"`
}

type ErrorResponse struct {
	Type    string `json:"Type"`
	Message string `json:"Message"`
	// Errors is missing by intention due to unclear meaning
}

func New(token string, url string) *API {
	api := new(API)
	api.AuthToken = token
	api.Endpoint = url
	return api
}

func Int2Pointer(input int) *int {
	result := new(int)
	*result = input
	return result
}

func (api *API) prepareSession() *napping.Session {
	s := new(napping.Session)
	h := &http.Header{}
	h.Set("X_TOKEN", api.AuthToken)
	s.Header = h
	return s
}

func logResult(response *napping.Response, expectedStatus int) {
	if response != nil {
		log.Debug("response Status:", response.Status())
		if response.Status() != expectedStatus {
			log.Debug("response:", response.RawText())
			log.Debug("response Headers:", response.HttpResponse().Header)
		}
	}
}

func isError(response *napping.Response, expectedStatus int, err error) error {
	if err != nil {
		return err
	}
	logResult(response, expectedStatus)
	if response != nil {
		if response.Status() == expectedStatus {
			// we got a response with the expected HTTP status code, hence no error
			return nil
		}

		// extract the API's error message to be returned later
		errorResponse := ErrorResponse{}
		err := response.Unmarshal(&errorResponse)
		if err != nil {
			log.Debug(err)
		}

		switch response.Status() {
		case http.StatusUnauthorized:
			// TODO: hopefully all error will have a proper error object in the body later
			return &ApiError{response.Status(), "Authentication is failed, please check your X-TOKEN"}
		default:
			return &ApiError{response.Status(), errorResponse.Message}
		}

	} else {
		// no response from API means generic error
	}
	return nil
}

func createUrl(api *API, sections ...interface{}) string {
	url := api.Endpoint
	for _, section := range sections {
		url += "/" + fmt.Sprint(section)
	}
	return url
}
