package xmlapi

import (
	"github.com/killmails/go-eve/client"
	"github.com/killmails/go-eve/schemas/xmlapi"
)

const (
	ServerStatusUrl client.EndpointUrl = "/server/ServerStatus.xml.aspx"
)

type (
	serverScope struct {
		*XmlApi
	}
)

func (x *XmlApi) Server() (r serverScope) {
	r = serverScope{x}
	return
}

func (x serverScope) ServerStatus() (result xmlapi.ServerStatus, err error) {
	err = x.Fetch(ServerStatusUrl, &result, nil)
	return
}
