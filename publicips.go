package oneandone_cloudserver_api

import ()

type PublicIp struct {
	withId
	IpAddress    string      `json:"ip"`
	Type         string      `json:"type"`
	AssignedTo   ServerForIp `json:"assigned_to"`
	ReverseDns   string      `json:"reverse_dns"`
	MacAddress   string      `json:"mac_address"`
	Dhcp         bool        `json:"dhcp"`
	State        string      `json:"state"`
	CreationDate string      `json:"creation_date"`
}

type ServerForIp struct {
	withId
	withName
	Type string `json:"type"`
}

type PublicIpCreateData struct {
	ReverseDns string `json:"reverse_dns"`
}

// GET /public_ips
func (api *API) GetPublicIps() []PublicIp {
}

// POST /public_ips
func (api *API) CreatePublicIp(configuration PublicIpCreateData) PublicIp {
}

// GET /public_ips/{id}
func (api *API) GetPublicIp(Id string) PublicIp {
}

// DELETE /public_ips/{id}
func (ip *PublicIp) Delete() PublicIp {
}

// PUT /public_ips/{id}
func (ip *PublicIp) UpdateReverseDns(NewReverseDns string) PublicIp {
}
