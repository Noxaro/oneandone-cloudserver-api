package oneandone_cloudserver_api

import (

)

type FirewallPolicy struct {
}

type FirewallPolicyCreateData struct {
}

func (api *API) GetFirewallPolicies() []FirewallPolicy {
}

func (api *API) CreateFirewallPolicy(configuration FirewallPolicyCreateData) Vm {
}

func (api *API) GetFirewallPolicy(Id string) FirewallPolicy {
}

func (fwp *FirewallPolicy) Delete() FirewallPolicy {
}
