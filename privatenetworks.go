package oneandone_cloudserver_api

import ()

type PrivateNetwork struct {
	withId
	withName
	CloudpanelId string `json:"cloudpanel_id"`
	withDescription
	NetworkAddress string   `json:"network_address"`
	SubnetMask     string   `json:"subnet_mask"`
	State          string   `json:"state"`
	CreationDate   string   `json:"creation_date"`
	Servers        []Server `json:"servers"`
}

type Server struct {
	withId
	withName
}

type PrivateNetworkCreateData struct {
}

// GET /private_networks
func (api *API) GetPrivateNetworks() []PrivateNetwork {
}

// POST /private_networks
func (api *API) CreatePrivateNetwork(configuration PrivateNetworkCreateData) Vm {
}

// GET /private_networks/{id}
func (api *API) GetPrivateNetwork(Id string) PrivateNetwork {
}

// DELETE /private_networks/{id}
func (fwp *PrivateNetwork) Delete() PrivateNetwork {
}

// PUT /private_networks/{id}

// GET /private_networks/{id}/servers

// PUT /private_networks/{id}/servers

// GET /private_networks/{id}/servers/{id}

// DELETE /private_networks/{id}/servers/{id}

