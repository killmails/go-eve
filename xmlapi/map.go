package xmlapi

import (
	"github.com/killmails/go-eve/client"
	"github.com/killmails/go-eve/schemas/xmlapi"
)

const (
	FacWarSystemsUrl client.EndpointUrl = "/map/FacWarSystems.xml.aspx"
	JumpsUrl         client.EndpointUrl = "/map/Jumps.xml.aspx"
	KillsUrl         client.EndpointUrl = "/map/Kills.xml.aspx"
	SovereigntyUrl   client.EndpointUrl = "/map/Sovereignty.xml.aspx"
)

type (
	mapScope struct {
		*XmlApi
	}
)

func (x *XmlApi) Map() (r mapScope) {
	r = mapScope{x}
	return
}

func (x mapScope) FacWarSystems() (result xmlapi.FacWarSystems, err error) {
	err = x.Fetch(FacWarSystemsUrl, &result, nil)
	return
}

func (x mapScope) Jumps() (result xmlapi.Jumps, err error) {
	err = x.Fetch(JumpsUrl, &result, nil)
	return
}

func (x mapScope) Kills() (result xmlapi.Kills, err error) {
	err = x.Fetch(KillsUrl, &result, nil)
	return
}

func (x mapScope) Sovereignty() (result xmlapi.Sovereignty, err error) {
	err = x.Fetch(SovereigntyUrl, &result, nil)
	return
}
