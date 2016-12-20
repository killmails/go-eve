package xmlapi

import (
	"github.com/killmails/go-eve/client"
	"github.com/killmails/go-eve/schemas/xmlapi"
)

const (
	accountStatusURL client.EndpointURL = "/account/AccountStatus.xml.aspx"
	aPIKeyInfoURL    client.EndpointURL = "/account/APIKeyInfo.xml.aspx"
	charactersURL    client.EndpointURL = "/account/Characters.xml.aspx"
)

type (
	// AccountScope is the API/account interface
	AccountScope struct {
		*XMLAPI
	}
)

// Account is the scope for /account/* API requests
func (x *XMLAPI) Account() (r AccountScope) {
	r = AccountScope{x}
	return
}

// AccountStatus returns information about a player’s EVE account like
// creation time, minutes spent in game etc.
func (x AccountScope) AccountStatus() (result xmlapi.AccountStatus, err error) {
	err = x.Fetch(accountStatusURL, &result, nil)
	return
}

// APIKeyInfo specifies the access rights of an API key.
func (x AccountScope) APIKeyInfo() (result xmlapi.APIKeyInfo, err error) {
	err = x.Fetch(aPIKeyInfoURL, &result, nil)
	return
}

// Characters lists all characters included in this API key.
func (x AccountScope) Characters() (result xmlapi.Characters, err error) {
	err = x.Fetch(charactersURL, &result, nil)
	return
}
