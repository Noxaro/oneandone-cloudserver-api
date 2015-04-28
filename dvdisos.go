package oneandone_cloudserver_api

import ()

type DvdIso struct {
	withId
	withName
	OsFamily     string `json:"os_family"`
	Os           string `json:"os"`
	OsVersion    string `json:"os_version"`
	Architecture int    `json:"architecture"`
	Type         string `json:"type"`
}

// GET /dvd_isos
func (api *API) GetDvdIsos() []DvdIso {
}
