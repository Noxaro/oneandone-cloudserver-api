/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package oneandone_cloudserver_api

import (
	"github.com/docker/machine/log"
	"net/http"
)

type ServerAppliance struct {
	withId
	withName
	OsImageType        string                   `json:"os_image_type"`
	OsFamily           string                   `json:"os_family"`
	Os                 string                   `json:"os"`
	OsVersion          string                   `json:"os_version"`
	MinHddSize         int                      `json:"min_hdd_size"`
	Architecture       int                      `json:"architecture"`
	Licenses           []ServerApplianceLicence `json:"licenses"`
	IsAutomaticInstall bool                     `json:"automatic_installation"`
	Type               string                   `json:"type"`
	withApi
}

type ServerApplianceLicence struct {
	withName
}

// GET /server_appliances
func (api *API) GetServerAppliances() ([]ServerAppliance, error) {
	log.Debug("requesting information about server appliances")
	res := []ServerAppliance{}
	err := api.Client.Get(createUrl(api, "server_appliances"), &res, http.StatusOK)
	if err != nil {
		return nil, err
	}
	for index, _ := range res {
		res[index].api = api
	}
	return res, nil
}

// GET /server_appliances/{id}
func (api *API) GetServerAppliance(Id string) (*ServerAppliance, error) {
	log.Debug("requesting information about server appliance", Id)
	res := new(ServerAppliance)
	err := api.Client.Get(createUrl(api, "server_appliances", Id), &res, http.StatusOK)
	if err != nil {
		return nil, err
	}
	res.api = api
	return res, nil
}

func (api* API) ServerApplianceListArchitectures(family string, os string, osType string) ([]int, error) {
	apps, err := api.GetServerAppliances()
	if err != nil {
		return nil, err
	}
	architectures := make(map[int]int)
	for index, _ := range apps {
		if apps[index].OsFamily == family && apps[index].Os == os && apps[index].OsImageType == osType {
			log.Debug(apps[index])
			architectures[apps[index].Architecture] = 1
		}
	}
	return GetMapKeysInt(architectures), nil
}

func (api *API) ServerApplianceListTypes(family string, os string) ([]string, error) {
	apps, err := api.GetServerAppliances()
	if err != nil {
		return nil, err
	}
	osTypes := make(map[string]int)
	for index, _ := range apps {
		if apps[index].OsFamily == family && apps[index].Os == os {
			log.Debug(apps[index])
			osTypes[apps[index].OsImageType] = 1
		}
	}
	return GetMapKeysString(osTypes), nil
}

func (api *API) ServerApplianceListOperationSystems(family string) ([]string, error) {
	apps, err := api.GetServerAppliances()
	if err != nil {
		return nil, err
	}
	os := make(map[string]int)
	for index, _ := range apps {
		if apps[index].OsFamily == family {
			log.Debug(apps[index])
			os[apps[index].OsVersion] = 1
		}
	}
	return GetMapKeysString(os), nil
}

func (api *API) ServerApplianceListFamilies() ([]string, error) {
	apps, err := api.GetServerAppliances()
	if err != nil {
		return nil, err
	}
	osFamilies := make(map[string]int)
	for index, _ := range apps {
		log.Debug(apps[index])
		osFamilies[apps[index].OsFamily] = 1
	}
	return GetMapKeysString(osFamilies), nil
}

