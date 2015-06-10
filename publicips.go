package oneandone_cloudserver_api

import (
	"github.com/docker/machine/log"
	"net/http"
)

type PublicIp struct {
	withId
	IpAddress    string      `json:"ip"`
	withType
	AssignedTo   ServerForIp `json:"assigned_to"`
	ReverseDns   string      `json:"reverse_dns"`
	IsDhcp       bool        `json:"dhcp"`
	State        string      `json:"state"`
	CreationDate string      `json:"creation_date"`
	withApi
}

type ServerForIp struct {
	withId
	withName
	withType
}

type PublicIpSettings struct {
	ReverseDns 	string 		`json:"reverse_dns"`
	Type 		string 		`json:"type"`
}

const  (
	IpTypeV4 = "IPV4"
	IpTypeV6 = "IPV6"
)

// GET /public_ips
func (api *API) GetPublicIps() ([]PublicIp, error) {
	log.Debug("Requesting information about public ips")
	result := []PublicIp{}
	apiError := api.RestClient.Get(api.RestClient.CreateUrl(PublicIpPathSegment), &result, http.StatusOK)
	if apiError != nil {
		return nil, apiError
	}
	for index, _ := range result {
		result[index].api = api
	}
	return result, nil
}

// POST /public_ips
func (api *API) CreatePublicIp(configuration PublicIpSettings) (*PublicIp, error) {
	log.Debug("Booking a new public ip with type: '%s' and reverse dns: '%s'", configuration.Type, configuration.ReverseDns)
	res := new(PublicIp)
	apiError := api.RestClient.Post(api.RestClient.CreateUrl(PublicIpPathSegment), &configuration, &res, http.StatusCreated)
	if apiError != nil {
		return nil, apiError
	}
	res.api = api
	return res, nil
}

// GET /public_ips/{id}
func (api *API) GetPublicIp(Id string) (*PublicIp, error) {
	log.Debugf("requesting information about the public ip: '%s'", Id)
	result := new(PublicIp)
	apiError := api.RestClient.Get(api.RestClient.CreateUrl(PublicIpPathSegment, Id), &result, http.StatusOK)
	if apiError != nil {
		return nil, apiError
	}
	result.api = api
	return result, nil
}

// DELETE /public_ips/{id}
func (ip *PublicIp) Delete() (*PublicIp, error) {
	log.Debugf("deleting public ip address '%s'", ip.Id)
	result := new(PublicIp)
	apiError := ip.api.RestClient.Delete(ip.api.RestClient.CreateUrl(PublicIpPathSegment, ip.Id), &result, http.StatusOK)
	if apiError != nil {
		return nil, apiError
	}
	result.api = ip.api
	return result, nil
}

// PUT /public_ips/{id}
func (ip *PublicIp) UpdateReverseDns(ipAddressConfiguration PublicIpSettings) (*PublicIp, error) {
	log.Debug("updating public ip address '%s'", ip.Id)
	result := new(PublicIp)
	apiError := ip.api.RestClient.Put(ip.api.RestClient.CreateUrl(PublicIpPathSegment, ip.Id), &ipAddressConfiguration, &result, http.StatusOK)
	if apiError != nil {
		return nil, apiError
	}
	result.api = ip.api
	return result, nil
}
