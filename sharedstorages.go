package oneandone_cloudserver_api

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

type SharedStorage struct {
	withId
	Size	 		int 	  `json:"size"`
	State 			string	  `json:"state"`
	withDescription
	CloudPanelId	string	  `json:"cloudpanel_id"`
	SizeUsed		string    `json:"size_used"`
    CifsPath		string    `json:"cifs_path"`
	NfsPath			string    `json:"nfs_path"`
	withName
	CreationDate	string    `json:"creation_date"`
	SharedStorage	[]SharedStorageServer	`json:"servers"`
	withApi
}

type SharedStorageServer struct {
	withId
	withName
	rights	string	`json:"rights"`
}

type SharedStorageCreate struct {
	Name 		string `json:"name"`
	Description string `json:"description"`
	Size 		int    `json:"size"`
}


// GET /shared_storages
func (api *API) GetSharedStorages() ([]SharedStorage, error) {
	log.Debug("requesting information about shared storages")
	session := api.prepareSession()
	result := []SharedStorage{}
	response, apiError := session.Get(createUrl(api, SharedStoragesPathSegment), nil, &result, nil)
	if resultError := isError(response, http.StatusOK, apiError); resultError != nil {
		return nil, resultError
	} else {
		for index, _ := range result {
			result[index].api = api
		}
		return result, nil
	}
}

// POST /shared_storages
func (api *API) CreateSharedStorage(configuration SharedStorageCreate) (*SharedStorage, error){
	log.Debug("Creating a new shared storage with name: '" + configuration.Name + "'")
	session := api.prepareSession()
	result := new(SharedStorage)
	response, apiError := session.Post(createUrl(api, SharedStoragesPathSegment), &configuration, &result, nil)
	if resultError := isError(response, http.StatusAccepted, apiError); resultError != nil {
		return nil, resultError
	} else {
		return result, nil
	}
}

// GET /shared_storages/{id}

// DELETE /shared_storages/{id}

// PUT /shared_storages/{id}

// GET /shared_storages/{id}/servers

// PUT /shared_storages/{id}/servers

// GET /shared_storages/{id}/servers/{id}

// DELETE /shared_storages/{id}/servers/{id}

// GET /shared_storages/access

// PUT /shared_storages/access

