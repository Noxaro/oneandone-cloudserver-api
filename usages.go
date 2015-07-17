/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package oneandone_cloudserver_api

import (
	"github.com/docker/machine/log"
	"net/http"
)

type Usages struct {
	PublicIPs string `json:"PUBLIC_IPS"`
	withApi
}

// GET /usages
func (api *API) GetUsages(period string) (*Usages, error) {
	log.Debug("requesting information about usages")
	result := Usages{}
	url := createUrl(api, "usages")
	url = appendQueryParams(url, map[string]interface{}{"period": period})
	err := api.Client.Get(url, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api
	return &result, nil
}
