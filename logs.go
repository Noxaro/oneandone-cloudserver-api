/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package oneandone_cloudserver_api

import (
	"github.com/docker/machine/log"
	"net/http"
)

type Log struct {
	withId
	StartDate string `json:"start_date"`
	EndDate  string `json:"end_date"`
	Duration int    `json:"duration"`
	Status   Status
	Action   string `json:"action"`
	withType
	Resource     LogDetails
	User         LogDetails
	CloudPanelId string `json:"cloudpanel_id"`
}

type LogDetails struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// GET /logs
func (api *API) GetLogs(period string) ([]Log, error) {
	log.Debug("requesting information about logs")
	result := []Log{}
	err := api.Client.Get(createUrl(api, "logs")+"?period="+period, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GET /logs/{id}
func (api *API) GetLog(id string) (*Log, error) {
	log.Debugf("requesting information about log: '%s'", id)
	result := new(Log)
	err := api.Client.Get(createUrl(api, "logs", id), &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return result, nil
}
