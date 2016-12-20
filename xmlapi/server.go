package xmlapi

import (
	"github.com/killmails/go-eve/client"
	"github.com/killmails/go-eve/schemas/xmlapi"
)

const (
	serverStatusURL client.EndpointURL = "/server/ServerStatus.xml.aspx"
)

type (
	// ServerScope is the API/server interface
	ServerScope struct {
		*XMLAPI
	}
)

// Server is the scope for /server/* API requests
func (x *XMLAPI) Server() (r ServerScope) {
	r = ServerScope{x}
	return
}

// ServerStatus provides status of the server up/down and the number of users
// logged onto the server.
func (x ServerScope) ServerStatus() (result xmlapi.ServerStatus, err error) {
	err = x.Fetch(serverStatusURL, &result, nil)
	return
}
