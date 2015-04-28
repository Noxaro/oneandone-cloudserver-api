package oneandone_cloudserver_api

import ()

type Server struct {
}

type Size struct {
	Vcores            int              `json:"vcore"`
	CoresPerProcessor int              `json:"cores_per_processor"`
	Ram               int              `json:"ram"`
	Hdds              []Hdd            `json:"hdds"`
	Dvd               Dvd              `json:"dvd"`
	Image             Image            `json:"image"`
	PrivateNetworks   []PrivateNetwork `json:"private_networks"`
	Alerts            []Alert          `json:"alerts"`
	MonitoringPolicy  MonitoringPolicy `json:"monitoring_policy"`
	Ips               []Ip             `json:"ips"`
}

type Hdd struct {
	withId
	Size   int  `json:"size"`
	IsMain bool `json:"is_main"`
}

type Dvd struct {
	withId
	withName
}

type Image struct {
	withId
	withName
}

type Alert struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type VmCreateData struct {
}

// GET /servers
func (api *API) GetServers() []Server {
}

// POST /servers
func (api *API) CreateServer(configuration ServerCreateData) Server {
}

// GET /servers/{id}
func (api *API) GetServer(Id string) Server {
}

// DELETE /servers/{id}
func (server *Server) Delete() Server {
}

// PUT /servers/{id}

// GET /servers/{id}/size

// PUT /servers/{id}/size

// GET /servers/{id}/size/hdds

// POST /servers/{id}/size/hdds

// GET /servers/{id}/size/hdds/{id}

// DELETE /servers/{id}/size/hdds/{id}

// PUT /servers/{id}/size/hdds/{id}

// GET /servers/{id}/image

// PUT /servers/{id}/image

// GET /servers/{id}/ips

// POST /servers/{id}/ips

// GET /servers/{id}/ips/{id}

// DELETE /servers/{id}/ips/{id}

// PUT /servers/{id}/ips/{id}

// GET /servers/{id}/status

// PUT /servers/{id}/status

// GET /servers/{id}/dvd

// DELETE /servers/{id}/dvd

// PUT /servers/{id}/dvd

// GET /servers/{id}/private_networks

// PUT /servers/{id}/private_networks

// GET /servers/{id}/private_networks/{id}

// DELETE /servers/{id}/private_networks/{id}

// GET /servers/{id}/snapshots

// POST /servers/{id}/snapshots

// GET /servers/{id}/snapshots/{id}

// DELETE /servers/{id}/snapshots/{id}

// PUT /servers/{id}/snapshots/{id}

// POST /servers/{server_id}/clone
func (server *Server) Clone(NewName string) Server {
}
