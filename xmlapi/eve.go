package xmlapi

import (
	"fmt"
	"strings"

	"github.com/killmails/go-eve/client"
	"github.com/killmails/go-eve/schemas/xmlapi"
	"github.com/killmails/go-eve/util"
)

const (
	AllianceListUrl           client.EndpointUrl = "/eve/AllianceList.xml.aspx"
	CharacterAffiliationUrl   client.EndpointUrl = "/eve/CharacterAffiliation.xml.aspx"
	CharacterIDUrl            client.EndpointUrl = "/eve/CharacterID.xml.aspx"
	CharacterInfoUrl          client.EndpointUrl = "/eve/CharacterInfo.xml.aspx"
	CharacterNameUrl          client.EndpointUrl = "/eve/CharacterName.xml.aspx"
	ConquerableStationListUrl client.EndpointUrl = "/eve/ConquerableStationList.xml.aspx"
	ErrorListUrl              client.EndpointUrl = "/eve/ErrorList.xml.aspx"
	RefTypesUrl               client.EndpointUrl = "/eve/RefTypes.xml.aspx"
	TypeNameUrl               client.EndpointUrl = "/eve/TypeName.xml.aspx"
)

type (
	eveScope struct {
		*XmlApi
	}
)

func (x *XmlApi) Eve() (r eveScope) {
	r = eveScope{x}
	return
}

func (x eveScope) AllianceList(ver int) (result xmlapi.AllianceList, err error) {
	params := client.Params{
		"version": fmt.Sprintf("%v", ver),
	}
	err = x.Fetch(AllianceListUrl, &result, params)
	return
}

func (x eveScope) CharacterAffiliation(ids []int) (result xmlapi.CharacterAffiliation, err error) {
	params := client.Params{
		"ids": strings.Join(util.SliceItoa(ids), ","),
	}
	err = x.Fetch(CharacterAffiliationUrl, &result, params)
	return
}

func (x eveScope) CharacterID(names []string) (result xmlapi.CharacterID, err error) {
	params := client.Params{
		"names": strings.Join(names, ","),
	}
	err = x.Fetch(CharacterIDUrl, &result, params)
	return
}

func (x eveScope) CharacterInfo(id int) (result xmlapi.CharacterInfo, err error) {
	params := client.Params{
		"characterID": fmt.Sprintf("%v", id),
	}
	err = x.Fetch(CharacterInfoUrl, &result, params)
	return
}

func (x eveScope) CharacterName(ids []int) (result xmlapi.CharacterName, err error) {
	params := client.Params{
		"ids": strings.Join(util.SliceItoa(ids), ","),
	}
	err = x.Fetch(CharacterNameUrl, &result, params)
	return
}

func (x eveScope) ConquerableStationList() (result xmlapi.ConquerableStationList, err error) {
	err = x.Fetch(ConquerableStationListUrl, &result, nil)
	return
}

func (x eveScope) ErrorList() (result xmlapi.ErrorList, err error) {
	err = x.Fetch(ErrorListUrl, &result, nil)
	return
}

func (x eveScope) RefTypes() (result xmlapi.RefTypes, err error) {
	err = x.Fetch(RefTypesUrl, &result, nil)
	return
}

func (x eveScope) TypeName(ids []int) (result xmlapi.TypeName, err error) {
	params := client.Params{
		"ids": strings.Join(util.SliceItoa(ids), ","),
	}
	err = x.Fetch(TypeNameUrl, &result, params)
	return
}
