package virtualboxclient

import (
	"github.com/appropriate/go-virtualboxclient/vboxwebsrv"
)

type Medium struct {
	managedObjectId string
}

func (svc *VirtualBoxClient) CreateHardDisk(format, location string) (*Medium, error) {
	svc.Logon()

	request := vboxwebsrv.IVirtualBoxcreateHardDisk{This: svc.managedObjectId, Format: format, Location: location}

	response, err := svc.client.IVirtualBoxcreateHardDisk(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &Medium{managedObjectId: response.Returnval}, nil
}
