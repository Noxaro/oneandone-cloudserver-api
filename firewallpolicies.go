package oneandone_cloudserver_api

import (
	"github.com/docker/machine/log"
	"net/http"
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
func (api *API) GetFirewallPolicies() ([]FirewallPolicy, error) {
	log.Debug("requesting information about firewall policies")
	session := api.prepareSession()
	result := []FirewallPolicy{}
	response, err := session.Get(createUrl(api, "firewall_policies"), nil, &result, nil)
	if err := isError(response, http.StatusOK, err); err != nil {
		return nil, err
	} else {
		for index, _ := range result {
			result[index].api = api
		}
		return result, nil
	}
}

// POST /firewall_policies
func (api *API) CreateFirewallPolicy(configuration FirewallPolicyCreateData) (*FirewallPolicy, error) {
	log.Debug("requesting to create a new firewall policy")
	s := api.prepareSession()
	result := new(FirewallPolicy)
	response, err := s.Post(createUrl(api, "firewall_policies"), configuration, &result, nil)
	if err := isError(response, http.StatusCreated, err); err != nil {
		return nil, err
	} else {
		result.api = api
		return result, nil
	}
}

// GET /firewall_policies/{id}
func (api *API) GetFirewallPolicy(Id string) (*FirewallPolicy, error) {
	log.Debug("requesting to about firewall policy ", Id)
	session := api.prepareSession()
	result := new(FirewallPolicy)
	response, err := session.Get(createUrl(api, "firewall_policies", Id), nil, &result, nil)
	if err := isError(response, http.StatusOK, err); err != nil {
		return nil, err
	} else {
		result.api = api
		return result, nil
	}
}

// DELETE /firewall_policies/{id}
func (fwp *FirewallPolicy) Delete() (*FirewallPolicy, error) {
	log.Debug("Requested to delete firewall policy ", fwp.Id)
	session := fwp.api.prepareSession()
	result := new(FirewallPolicy)
	response, err := session.Delete(createUrl(fwp.api, "firewall_policies", fwp.Id), &result, nil)
	if err := isError(response, http.StatusOK, err); err != nil {
		return nil, err
	} else {
		result.api = fwp.api
		return result, nil
	}
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
