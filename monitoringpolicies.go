/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package oneandone_cloudserver_api

import (
	"github.com/docker/machine/log"
	"net/http"
)

type MonitoringPolicy struct {
	withId
	withDescription
	withApi
	Status
	Name         string              `json:"name"`
	Default      int                 `json:"default"`
	CreationDate string              `json:"creation_date"`
	Email        string              `json:"email"`
	Agent        bool                `json:"agent"`
	Servers      []MonitoringServer  `json:"servers"`
	Thresholds   MonitoringThreshold `json:"thresholds"`
	Ports        []MonitoringPort    `json:"ports"`
	Processes    []MonitoringProcess `json:"processes"`
	CloudPanelId string              `json:"cloudpanel_id"`
}

type MonitoringServer struct {
	withId
	withName
}

type MonitoringThreshold struct {
	Cpu          MonitoringLevel `json:"cpu"`
	Ram          MonitoringLevel `json:"ram"`
	Disk         MonitoringLevel `json:"disk"`
	Transfer     MonitoringLevel `json:"transfer"`
	InternalPing MonitoringLevel `json:"internal_ping"`
}

type MonitoringLevel struct {
	Warning  MonitoringValue `json:"warning"`
	Critical MonitoringValue `json:"critical"`
}

type MonitoringValue struct {
	Value int  `json:"value"`
	Alert bool `json:"alert"`
}

type MonitoringPort struct {
	Protocol          string `json:"protocol"`
	Port              int    `json:"port"`
	AlertIf           string `json:"alert_if"`
	EmailNotification bool   `json:"email_notification"`
}

type MonitoringProcess struct {
	Process           string `json:"process"`
	AlertIf           string `json:"alert_if"`
	EmailNotification bool   `json:"email_notification"`
}

type MonitoringStatus struct {
	Resource string          `json:"resource"`
	Date     string          `json:"date"`
	State    string          `json:"state"`
	Value    MonitoringValue `json:"value"`
}
type MonitoringLimits struct {
	Upper int `json:"upper"`
	Lower int `json:"lower"`
}

type MonitoringPolicyCreateData struct {
}

// GET /monitoring_policies
func (api *API) GetMonitoringPolicies() ([]MonitoringPolicy, error) {
	log.Debug("requesting information about monitoring policies")
	result := []MonitoringPolicy{}
	err := api.Client.Get(createUrl(api, "monitoring_policies"), &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	for index, _ := range result {
		result[index].api = api
	}
	return result, nil
}

// POST /monitoring_policies
func (api *API) CreateMonitoringPolicy(configuration MonitoringPolicy) (*MonitoringPolicy, error) {
	log.Debug("requesting to create a new monitoring policy")
	result := new(MonitoringPolicy)
	err := api.Client.Post(createUrl(api, "monitoring_policies"), &configuration, &result, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}

// GET /monitoring_policies/{id}
func (api *API) GetMonitoringPolicy(Id string) (*MonitoringPolicy, error) {
	log.Debug("requesting information about monitoring policy ", Id)
	result := new(MonitoringPolicy)
	err := api.Client.Get(createUrl(api, "monitoring_policies", Id), &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}

// DELETE /monitoring_policies/{id}
func (mp *MonitoringPolicy) Delete() MonitoringPolicy {
	log.Debugf("Requested to delete monitoring policy '%v' ", mp.Id) (*MonitoringPolicy, error)
	result := new(MonitoringPolicy)
	err := mp.api.Client.Delete(createUrl(mp.api, "monitoring_policies", mp.Id), &result, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	result.api = mp.api
	return result, nil
}

// PUT /monitoring_policies/{id}

// GET /monitoring_policies/{id}/ports

// PUT /monitoring_policies/{id}/ports

// GET /monitoring_policies/{id}/ports/{id}

// DELETE /monitoring_policies/{id}/ports/{id}

// PUT /monitoring_policies/{id}/ports/{id}

// GET /monitoring_policies/{id}/processes

// PUT /monitoring_policies/{id}/processes

// GET /monitoring_policies/{id}/processes/{id}

// DELETE /monitoring_policies/{id}/processes/{id}

// PUT /monitoring_policies/{id}/processes/{id}

// GET /monitoring_policies/{id}/servers

// PUT /monitoring_policies/{id}/servers

// GET /monitoring_policies/{id}/servers/{id}

// DELETE /monitoring_policies/{id}/servers/{id}
