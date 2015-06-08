package oneandone_cloudserver_api

import (
	log "github.com/Sirupsen/logrus"
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
	session := api.prepareSession()
	result := []PublicIp{}
	response, apiError := session.Get(createUrl(api, PublicIpPathSegment), nil, &result, nil)
	if resultError := isError(response, http.StatusOK, apiError); resultError != nil {
		return nil, resultError
	}
	for index, _ := range result {
		result[index].api = api
	}
	return result, nil
}

// POST /public_ips
func (api *API) CreatePublicIp(configuration PublicIpSettings) (*PublicIp, error) {
	log.Debug("Booking a new public ip with type: '" + configuration.Type + "' and reverse dns: '" + configuration.ReverseDns + "'")
	session := api.prepareSession()
	res := new(PublicIp)
	response, apiError := session.Post(createUrl(api, PublicIpPathSegment), &configuration, &res, nil)
	if resultError := isError(response, http.StatusCreated, apiError); resultError != nil {
		return nil, resultError
	} else {
		res.api = api
		return res, nil
	}
}

// GET /public_ips/{id}
func (api *API) GetPublicIp(Id string) (*PublicIp, error) {
	log.Debug("requesting information about the public ip: '" + Id + "'")
	session := api.prepareSession()
	result := new(PublicIp)
	response, apiError := session.Get(createUrl(api, PublicIpPathSegment, Id), nil, &result, nil)
	if resultError := isError(response, http.StatusOK, apiError); resultError != nil {
		return nil, resultError
	} else {
		result.api = api
		return result, nil
	}
}

// DELETE /public_ips/{id}
func (ip *PublicIp) Delete() (*PublicIp, error) {
	log.Debug("deleting public ip address '" + ip.Id + "'")
	session := ip.api.prepareSession()
	result := new(PublicIp)
	response, apiError := session.Delete(createUrl(ip.api, PublicIpPathSegment, ip.Id), &result, nil)
	if resultError := isError(response, http.StatusOK, apiError); resultError != nil {
		return nil, resultError
	} else {
		result.api = ip.api
		return result, nil
	}
}

// PUT /public_ips/{id}
func (ip *PublicIp) UpdateReverseDns(ipAddressConfiguration PublicIpSettings) (*PublicIp, error) {
	log.Debug("updating public ip address '" + ip.Id + "'")
	session := ip.api.prepareSession()
	result := new(PublicIp)
	response, apiError := session.Put(createUrl(ip.api, PublicIpPathSegment, ip.Id), &ipAddressConfiguration, &result, nil)
	if resultError := isError(response, http.StatusOK, apiError); resultError != nil {
		return nil, resultError
	} else {
		result.api = ip.api
		return result, nil
	}
}
