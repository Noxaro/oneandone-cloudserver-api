package oneandone_cloudserver_api

import (
	"github.com/docker/machine/log"
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
	result := []SharedStorage{}
	resultError := api.RestClient.Get(api.RestClient.CreateUrl(SharedStoragesPathSegment), &result, http.StatusOK)
	if resultError != nil {
		return nil, resultError
	}
	for index, _ := range result {
		result[index].api = api
	}
	return result, nil
}

// POST /shared_storages
func (api *API) CreateSharedStorage(configuration SharedStorageCreate) (*SharedStorage, error){
	log.Debugf("Creating a new shared storage with name '%s'", configuration.Name)
	result := new(SharedStorage)
	resultError := api.RestClient.Post(api.RestClient.CreateUrl(SharedStoragesPathSegment), configuration, &result, http.StatusAccepted)
	if resultError != nil {
		return nil, resultError
	}
	result.api = api
	return result, nil
}

// GET /shared_storages/{id}
func (api *API) GetSharedStorage(sharedStorageId string) (*SharedStorage, error) {
	log.Debugf("Requesting information about the shared storage with the id: '%s'", sharedStorageId)
	result := new(SharedStorage)
	resultError := api.RestClient.Get(api.RestClient.CreateUrl(SharedStoragesPathSegment, sharedStorageId), &result, http.StatusOK)
	if resultError != nil {
		return nil, resultError
	}
	result.api = api
	return result, nil
}

// DELETE /shared_storages/{id}
func (st *SharedStorage) DeleteSharedStorage() (*SharedStorage, error) {
	log.Debugf("Trying to delete shared storage with id: '%s'", st.Id)
	result := new(SharedStorage)
	resultError := st.api.RestClient.Delete(st.api.RestClient.CreateUrl(SharedStoragesPathSegment, st.Id), &result, http.StatusOK)
	if resultError != nil {
		return nil, resultError
	}
	return result, nil
}

// PUT /shared_storages/{id}

// GET /shared_storages/{id}/servers

// PUT /shared_storages/{id}/servers

// GET /shared_storages/{id}/servers/{id}

// DELETE /shared_storages/{id}/servers/{id}

// GET /shared_storages/access

// PUT /shared_storages/access

