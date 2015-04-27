package oneandone_cloudserver_api

import (

)

type Vm struct {
}

type VmCreateData struct {
}

func (api *API) GetVms() []Vm {
}

func (api *API) CreateVm(configuration VmCreateData) Vm {
}

func (api *API) GetVm(Id string) Vm {
}

func (vm *Vm) Delete() Vm {
}

func (vm *Vm) Start() Vm {
}

func (vm *Vm) Stop(hardwarePowerOff bool) Vm {
}

func (vm *Vm) Reinstall(applianceId string) Vm {
}

