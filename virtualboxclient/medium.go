package virtualboxclient

import (
	"github.com/appropriate/go-virtualboxclient/vboxwebsrv"
)

type Medium struct {
	client          *vboxwebsrv.VboxPortType
	managedObjectId string
}

func (svc *VirtualBoxClient) CreateHardDisk(format, location string) (*Medium, error) {
	svc.Logon()

	request := vboxwebsrv.IVirtualBoxcreateHardDisk{This: svc.managedObjectId, Format: format, Location: location}

	response, err := svc.client.IVirtualBoxcreateHardDisk(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &Medium{client: svc.client, managedObjectId: response.Returnval}, nil
}

func (m *Medium) CreateBaseStorage(logicalSize int64, variant []*vboxwebsrv.MediumVariant) error {
	request := vboxwebsrv.IMediumcreateBaseStorage{This: m.managedObjectId, LogicalSize: logicalSize, Variant: variant}

	_, err := m.client.IMediumcreateBaseStorage(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return nil
}

func (m *Medium) DeleteStorage() error {
	request := vboxwebsrv.IMediumdeleteStorage{This: m.managedObjectId}

	_, err := m.client.IMediumdeleteStorage(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return nil
}
