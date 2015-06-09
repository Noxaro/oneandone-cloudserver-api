package oneandone_cloudserver_api

import (
	log "github.com/docker/machine/log"
	"net/http"
)

// Struct to describe a ISO image that can be used to boot a server.
//
// Values of this type describe ISO images that can be inserted into the servers virtual DVD drive.
//
//
type DvdIso struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	OsFamily     string `json:"os_family"`
	Os           string `json:"os"`
	OsVersion    string `json:"os_version"`
	Architecture int    `json:"architecture"`
	Type         string `json:"type"`
	withApi
}

// GET /dvd_isos
func (api *API) GetDvdIsos() ([]DvdIso, error) {
	log.Debug("requesting information about dvd isos")
	session := api.prepareSession()
	result := []DvdIso{}
	response, err := session.Get(createUrl(api, "dvd_isos"), nil, &result, nil)
	if err := isError(response, http.StatusOK, err); err != nil {
		return nil, err
	} else {
		for index, _ := range result {
			result[index].api = api
		}
		return result, nil
	}
}

// GET /dvd_isos/{id}
func (api *API) GetDvdIso(Id string) (*DvdIso, error) {
	log.Debug("requesting information about dvd iso", Id)
	session := api.prepareSession()
	result := new(DvdIso)
	response, err := session.Get(createUrl(api, "dvd_isos", Id), nil, &result, nil)
	if err := isError(response, http.StatusOK, err); err != nil {
		return nil, err
	} else {
		result.api = api
		return result, nil
	}

}
