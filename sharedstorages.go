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
	Id string `json:"id"`
	Name string `json:"name"`
	Rights	string	`json:"rights"`
}

type SharedStorageSettings struct {
	Name 		string `json:"name"`
	Description string `json:"description"`
	Size 		int    `json:"size"`
}


// GET /shared_storages
func (api *API) GetSharedStorages() ([]SharedStorage, error) {
	log.Debug("Requesting information about shared storages")
	result := []SharedStorage{}
	err := api.Client.Get(createUrl(api, SharedStoragesPathSegment), &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	for index, _ := range result {
		result[index].api = api
	}
	return result, nil
}

// POST /shared_storages
func (api *API) CreateSharedStorage(configuration SharedStorageSettings) (*SharedStorage, error){
	log.Debugf("Creating a new shared storage with name '%s'", configuration.Name)
	result := new(SharedStorage)
	err := api.Client.Post(createUrl(api, SharedStoragesPathSegment), configuration, &result, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}

// GET /shared_storages/{id}
func (api *API) GetSharedStorage(sharedStorageId string) (*SharedStorage, error) {
	log.Debugf("Requesting information about the shared storage with the id: '%s'", sharedStorageId)
	result := new(SharedStorage)
	err := api.Client.Get(createUrl(api, SharedStoragesPathSegment, sharedStorageId), &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = api
	return result, nil
}

// DELETE /shared_storages/{id}
func (st *SharedStorage) DeleteSharedStorage() (*SharedStorage, error) {
	log.Debugf("Deleteing shared storage with id: '%s'", st.Id)
	result := new(SharedStorage)
	err := st.api.Client.Delete(createUrl(st.api, SharedStoragesPathSegment, st.Id), &result, http.StatusAccepted)
	if err != nil {
		return nil, err
	}
	result.api = st.api
	return result, nil
}

// PUT /shared_storages/{id}
func (st *SharedStorage) UpdateSharedStorage(configuration SharedStorageSettings) (*SharedStorage, error) {
	log.Debugf("Updateing the shared storage with the id: '%s'", st.Id)
	result := new(SharedStorage)
	err := st.api.Client.Put(createUrl(st.api, SharedStoragesPathSegment, st.Id), configuration, &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	result.api = st.api
	return result, nil
}

// GET /shared_storages/{id}/servers
func (st *SharedStorage) GetSharedStorageServersPermissions() ([]SharedStorageServer, error){
	log.Debugf("Requesting servers with permissions the the shared storage with the id: '%s'", st.Id)
	result := []SharedStorageServer{}
	err := st.api.Client.Get(createUrl(st.api, SharedStoragesPathSegment, st.Id, "servers"), &result, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PUT /shared_storages/{id}/servers
func (st *SharedStorage) UpdateSharedStorageServerPermissions(sharedStorageServer SharedStorageServer) (*SharedStorage, error) {
	log.Debugf("Updateing server permissions for the shared storage with the id: '%s'", st.Id)
	result := new(SharedStorage)
	resultError := st.api.Client.Put(createUrl(st.api, SharedStoragesPathSegment, st.Id, "Servers"), &sharedStorageServer, &result, http.StatusOK)
	if resultError != nil {
		return nil, resultError
	}
	return result, nil
}

// GET /shared_storages/{id}/servers/{id}

// DELETE /shared_storages/{id}/servers/{id}

// GET /shared_storages/access

// PUT /shared_storages/access

