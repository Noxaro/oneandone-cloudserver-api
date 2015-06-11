package oneandone_cloudserver_api

import (
	"fmt"
	//log "github.com/docker/machine/log"
	//"net/http"
)

// Struct to hold the required information for accessing the API.
//
// Instances of this type contain the URL of the endpoint to access the API as well as the API access token to be used.
// They offer also all methods that allow to access the various objects that are returned by top level resources of
// the API.
type API struct {
	Endpoint string
	Client   *RestClient
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

// Struct to hold the status of an API object.
//
// Values of this type are used to represent the status of API objects like servers, firewall policies and the like.
//
// The value of the "State" field can represent fixed states like "ACTIVE" or "POWERED_ON" but also transitional
// states like "POWERING_ON" or "CONFIGURING".
//
// For fixed states the "Percent" field is empty where as for transitional states it contains the progress of the
// transition in percent.
type Status struct {
	State   string `json:"state"`
	Percent string `json:"percent"`
}

type errorResponse struct {
	Type    string `json:"Type"`
	Message string `json:"Message"`
	// TODO Errors is missing by intention due to unclear meaning
}

// Creates a new API instance.
//
// Explanations about given token and url information can be found online under the following url TODO add url!
func New(token string, url string) *API {
	api := new(API)
	api.Endpoint = url
	api.Client = NewRestClient(token)
	return api
}

// Function to convert a given integer value into a pointer to the same value.
//
// This function is used to be able to define ports with the CreateFirewallPolicy function and the definition of ports
// in the FirewallPolicyCreateData struct.
func Int2Pointer(input int) *int {
	result := new(int)
	*result = input
	return result
}

/*
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
		errorResponse := errorResponse{}
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
*/
func createUrl(api *API, sections ...interface{}) string {
	url := api.Endpoint
	for _, section := range sections {
		url += "/" + fmt.Sprint(section)
	}
	return url
}
