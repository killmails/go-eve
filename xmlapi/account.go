package xmlapi

import (
	"github.com/killmails/go-eve/client"
	"github.com/killmails/go-eve/schemas/xmlapi"
)

const (
	AccountStatusUrl client.EndpointUrl = "/account/AccountStatus.xml.aspx"
	APIKeyInfoUrl    client.EndpointUrl = "/account/APIKeyInfo.xml.aspx"
	CharactersUrl    client.EndpointUrl = "/account/Characters.xml.aspx"
)

type (
	accountScope struct {
		*XmlApi
	}
)

func (x *XmlApi) Account() (r accountScope) {
	r = accountScope{x}
	return
}

func (x accountScope) AccountStatus() (result xmlapi.AccountStatus, err error) {
	err = x.Fetch(AccountStatusUrl, &result, nil)
	return
}

func (x accountScope) APIKeyInfo() (result xmlapi.APIKeyInfo, err error) {
	err = x.Fetch(APIKeyInfoUrl, &result, nil)
	return
}

func (x accountScope) Characters() (result xmlapi.Characters, err error) {
	err = x.Fetch(CharactersUrl, &result, nil)
	return
}
