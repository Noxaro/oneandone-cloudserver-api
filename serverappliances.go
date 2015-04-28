package oneandone_cloudserver_api

import ()

type ServerAppliance struct {
	withId
	withName
	OsImageType        string   `json:"os_image_type"`
	OsFamily           string   `json:"os_family"`
	Os                 string   `json:"os"`
	OsVersion          string   `json:"os_version"`
	MinHddSize         int      `json:"min_hdd_size"`
	Architecture       int      `json:"architecture"`
	AdditionalSoftware []string `json:"additional_software"`
	Licenses           []string `json:"licenses"`
	Type               string   `json:"type"`
}

// GET /server_appliances
func (api *API) GetServerAppliances() []ServerAppliance {
}
