package virtualboxclient

import (
	"github.com/appropriate/go-virtualboxclient/vboxwebsrv"
)

type Machine struct {
	virtualbox      *VirtualBox
	managedObjectId string
}

func (m *Machine) GetChipsetType() (*vboxwebsrv.ChipsetType, error) {
	request := vboxwebsrv.IMachinegetChipsetType{This: m.managedObjectId}

	response, err := m.virtualbox.IMachinegetChipsetType(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (m *Machine) GetNetworkAdapter(slot uint32) (*NetworkAdapter, error) {
	request := vboxwebsrv.IMachinegetNetworkAdapter{This: m.managedObjectId, Slot: slot}

	response, err := m.virtualbox.IMachinegetNetworkAdapter(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &NetworkAdapter{m.virtualbox, response.Returnval}, nil
}
