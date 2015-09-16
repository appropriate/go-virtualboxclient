package virtualboxclient

import (
	"github.com/appropriate/go-virtualboxclient/vboxwebsrv"
)

type HardDisk struct {
	managedObjectId string
}

func (svc *VirtualBoxClient) CreateHardDisk(format, location string) (*HardDisk, error) {
	svc.Logon()

	request := vboxwebsrv.IVirtualBoxcreateHardDisk{This: svc.managedObjectId, Format: format, Location: location}

	response, err := svc.client.IVirtualBoxcreateHardDisk(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &HardDisk{managedObjectId: response.Returnval}, nil
}
