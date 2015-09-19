package virtualboxclient

import (
	"github.com/appropriate/go-virtualboxclient/vboxwebsrv"
)

type VirtualBox struct {
	*vboxwebsrv.VboxPortType

	username string
	password string

	managedObjectId string
}

func New(username, password, url string) *VirtualBox {
	return &VirtualBox{
		VboxPortType: vboxwebsrv.NewVboxPortType(url, false, nil),

		username: username,
		password: password,
	}
}

func (vb *VirtualBox) Logon() error {
	if vb.managedObjectId != "" {
		// Already logged in
		return nil
	}

	request := vboxwebsrv.IWebsessionManagerlogon{
		Username: vb.username,
		Password: vb.password,
	}

	response, err := vb.IWebsessionManagerlogon(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	vb.managedObjectId = response.Returnval

	return nil
}

func (vb *VirtualBox) CreateHardDisk(format, location string) (*Medium, error) {
	vb.Logon()

	request := vboxwebsrv.IVirtualBoxcreateHardDisk{This: vb.managedObjectId, Format: format, Location: location}

	response, err := vb.IVirtualBoxcreateHardDisk(&request)
	if err != nil {
		return nil, err // TODO: Wrap the error
	}

	return &Medium{virtualbox: vb, managedObjectId: response.Returnval}, nil
}
