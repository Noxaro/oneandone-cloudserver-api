package oneandone_cloudserver_api

import (
	log "github.com/Sirupsen/logrus"
)

type FirewallPolicy struct {
	withId
	withName
	withDescription
	// should be fixed: Status        Status                   `json:"status"`
	Status        string                   `json:"state"`
	DefaultPolicy int                      `json:"default"`
	Rules         []FirewallPolicyRules    `json:"rules"`
	ServerIps     []FirewallPolicyServerIp `json:"server_ips"`
	withApi
}

type FirewallPolicyRules struct {
	withId
	Protocol string `json:"protocol"`
	PortFrom *int   `json:"port_from"`
	PortTo   *int   `json:"port_to"`
	SourceIp string `json:"source"`
}

type FirewallPolicyServerIp struct {
	withId
	Ip         string `json:"ip"`
	ServerName string `json:"server_name"`
}

type FirewallPolicyCreateData struct {
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	Rules       []FirewallPolicyRulesCreateData `json:"rules"`
}

type FirewallPolicyRulesCreateData struct {
	Protocol string `json:"protocol"`
	PortFrom *int   `json:"port_from"`
	PortTo   *int   `json:"port_to"`
	SourceIp string `json:"source"`
}

// GET /firewall_policies
func (api *API) GetFirewallPolicies() []FirewallPolicy {
	log.Debug("requesting information about firewall policies")
	session := api.prepareSession()
	res := []FirewallPolicy{}
	resp, _ := session.Get(createUrl(api, "firewall_policies"), nil, &res, nil)
	logResult(resp, 200)
	for index, _ := range res {
		res[index].api = api
	}
	return res
}

// POST /firewall_policies
func (api *API) CreateFirewallPolicy(configuration FirewallPolicyCreateData) FirewallPolicy {
	log.Debug("requesting to create a new firewall policy")
	s := api.prepareSession()
	res := FirewallPolicy{}
	resp, _ := s.Post(createUrl(api, "firewall_policies"), configuration, &res, nil)
	logResult(resp, 201)
	res.api = api
	return res
}

// GET /firewall_policies/{id}
func (api *API) GetFirewallPolicy(Id string) FirewallPolicy {
	log.Debug("requesting to about firewall policy ", Id)
	session := api.prepareSession()
	res := FirewallPolicy{}
	resp, _ := session.Get(createUrl(api, "firewall_policies", Id), nil, &res, nil)
	logResult(resp, 200)
	res.api = api
	return res
}

// DELETE /firewall_policies/{id}
func (fwp *FirewallPolicy) Delete() FirewallPolicy {
	log.Debug("Requested to delete firewall policy ", fwp.Id)
	session := fwp.api.prepareSession()
	res := FirewallPolicy{}
	resp, _ := session.Delete(createUrl(fwp.api, "firewall_policies", fwp.Id), &res, nil)
	logResult(resp, 200)
	res.api = fwp.api
	return res

}

// PUT /firewall_policies/{id}

// GET /firewall_policies/{id}/server_ips

// PUT /firewall_policies/{id}/server_ips

// GET /firewall_policies/{id}/server_ips/{id}

// DELETE /firewall_policies/{id}/server_ips/{id}

// GET /firewall_policies/{id}/rules

// PUT /firewall_policies/{id}/rules

// GET /firewall_policies/{id}/rules/{id}

// DELETE /firewall_policies/{id}/rules/{id}
