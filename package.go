package oneandone_cloudserver_api

import (
	"github.com/jmcvetta/napping"
	"net/http"
)

const Endpoint string = "https://cloudpanel-api.1and1.com/v1"

type API struct {
	AuthToken string
	Endpoint  string
}

type withId struct {
	Id string `json:"id"`
}

type withName struct {
	Name string `json:"name"`
}

type withDescription struct {
	Description string `json:"description"`
}

func (api *API) prepareSession() *napping.Session {
	s := new(napping.Session)
	h := &http.Header{}
	h.Set("X_TOKEN", api.AuthToken)
	s.Header = h
	return s
}
