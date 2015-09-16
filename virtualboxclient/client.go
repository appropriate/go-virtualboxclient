package virtualboxclient

import (
	"github.com/appropriate/go-virtualboxclient/vboxwebsrv"
)

type VirtualBoxClient struct {
	username string
	password string
	url      string

	client          *vboxwebsrv.VboxPortType
	managedObjectId string
}

func New(username, password, url string) *VirtualBoxClient {
	return &VirtualBoxClient{
		username: username,
		password: password,
		url:      url,
	}
}

func (svc *VirtualBoxClient) Logon() error {
	svc.client = vboxwebsrv.NewVboxPortType(svc.url, false, nil)

	request := vboxwebsrv.IWebsessionManagerlogon{
		Username: svc.username,
		Password: svc.password,
	}

	response, err := svc.client.IWebsessionManagerlogon(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	svc.managedObjectId = response.Returnval

	return nil
}

func (svc *VirtualBoxClient) SessionId() string {
	return svc.managedObjectId
}
