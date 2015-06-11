package oneandone_cloudserver_api

import (
	log "github.com/docker/machine/log"
	"net/http"
)

type ServerAppliance struct {
	withId
	withName
	OsImageType        string   `json:"os_image_type"`
	OsFamily           string   `json:"os_family"`
	Os                 string   `json:"os"`
	OsVersion          string   `json:"os_version"`
	MinHddSize         int      `json:"min_hdd_size"`
	Architecture       int      `json:"architecture"`
	Licenses           []string `json:"licenses"`
	IsAutomaticInstall bool     `json:"automatic_installation"`
	Type               string   `json:"type"`
	withApi
}

// GET /server_appliances
func (api *API) GetServerAppliances() []ServerAppliance {
	log.Debug("requesting information about server appliances")
	res := []ServerAppliance{}
	resp, _ := api.Client.Get(createUrl(api, "server_appliances"), &res, http.StatusOK)
	logResult(resp, 200)
	for index, _ := range res {
		res[index].api = api
	}
	return res
}

// GET /server_appliances/{id}
func (api *API) GetServerAppliance(Id string) ServerAppliance {
	log.Debug("requesting information about server appliance", Id)
	res := ServerAppliance{}
	resp, _ := api.Client.Get(createUrl(api, "server_appliances", Id), &res, http.StatusOK)
	logResult(resp, 200)
	res.api = api
	return res
}
