# 1&1 Cloudserver API Go Client
  
This project contains a Go implementation of a client for the 1&1 Cloudserver's public API.
    
## About 1&1 Cloudserver API

The Cloudserver API Documentation can be found here: []()
      
The Go Client Documentation can be found here: []()

## Usage

Import the library:

```go
import oaocs "github.com/jlusiardi/oneandone-cloudserver-api"
```

Create a new API instance:

```go
api := oaocs.New("YOUR_TOKEN", "API_ENDPOINT")
```
### Servers

Query all server:

```go
servers, err := api.GetServers()
if err != nil {
	// some error handling
}
```

Create a new server:

```go
server, err := api.CreateServer(oaocs.ServerCreateData{
        Name:             "Some Name",
        Description:      "Some Description",
        ApplianceId:      "ID of some Appliance",
        FirewallPolicyId: "ID of applied Firewall Policy",
        Hardware: oaocs.Hardware{
                CoresPerProcessor: 1,
                Vcores:            1,	// 1 Core
                Ram:               1,	// 1GB Ram
                Hdds: []oaocs.Hdd{
                        oaocs.Hdd{
                                IsMain: true,
                                Size:   20,	// 20GB SSD
                	},
        	},
        },
        PowerOn: true,
})
if err != nil {
	// error handling
}
```

### Firewall Policies

Create a new firewall policy that opens all TCP ports:

```go
firewall, err := api.CreateFirewallPolicy(oaocs.FirewallPolicyCreateData{
        Name:        "Name of the Policy",
        Description: "Description of the Policy",
        Rules: []oaocs.FirewallPolicyRulesCreateData{
                oaocs.FirewallPolicyRulesCreateData{
                         Protocol: "TCP",
                         PortFrom: oaocs.Int2Pointer(1),
                         PortTo:   oaocs.Int2Pointer(65535),
                         SourceIp: "0.0.0.0",
                },
        },
})
if err != nil {
        // error handling
}
```
  
## Contributing
    
We are happy to recieve your reports on any issue and also pull requests are welcome!
