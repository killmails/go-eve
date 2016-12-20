package xmlapi

import (
	"fmt"
	"strings"

	"github.com/killmails/go-eve/client"
	"github.com/killmails/go-eve/schemas/xmlapi"
	"github.com/killmails/go-eve/util"
)

const (
	allianceListURL           client.EndpointURL = "/eve/AllianceList.xml.aspx"
	characterAffiliationURL   client.EndpointURL = "/eve/CharacterAffiliation.xml.aspx"
	characterIDURL            client.EndpointURL = "/eve/CharacterID.xml.aspx"
	characterInfoURL          client.EndpointURL = "/eve/CharacterInfo.xml.aspx"
	characterNameURL          client.EndpointURL = "/eve/CharacterName.xml.aspx"
	conquerableStationListURL client.EndpointURL = "/eve/ConquerableStationList.xml.aspx"
	errorListURL              client.EndpointURL = "/eve/ErrorList.xml.aspx"
	refTypesURL               client.EndpointURL = "/eve/RefTypes.xml.aspx"
	typeNameURL               client.EndpointURL = "/eve/TypeName.xml.aspx"
)

type (
	// EveScope is the API/eve interface
	EveScope struct {
		*XMLAPI
	}
)

// Eve is the scope for /eve/* API requests
func (x *XMLAPI) Eve() (r EveScope) {
	r = EveScope{x}
	return
}

// AllianceList returns active Alliances in New Eden and their member Corporations.
func (x EveScope) AllianceList(ver int) (result xmlapi.AllianceList, err error) {
	params := client.Params{
		"version": fmt.Sprintf("%v", ver),
	}
	err = x.Fetch(allianceListURL, &result, params)
	return
}

// CharacterAffiliation TODO
func (x EveScope) CharacterAffiliation(ids []int) (result xmlapi.CharacterAffiliation, err error) {
	params := client.Params{
		"ids": strings.Join(util.SliceItoa(ids), ","),
	}
	err = x.Fetch(characterAffiliationURL, &result, params)
	return
}

// CharacterID TODO
func (x EveScope) CharacterID(names []string) (result xmlapi.CharacterID, err error) {
	params := client.Params{
		"names": strings.Join(names, ","),
	}
	err = x.Fetch(characterIDURL, &result, params)
	return
}

// CharacterInfo returns information about the character.
func (x EveScope) CharacterInfo(id int) (result xmlapi.CharacterInfo, err error) {
	params := client.Params{
		"characterID": fmt.Sprintf("%v", id),
	}
	err = x.Fetch(characterInfoURL, &result, params)
	return
}

// CharacterName returns the names for a comma-separated list of owner IDs
// for characters, corporations, alliances, and so on.
func (x EveScope) CharacterName(ids []int) (result xmlapi.CharacterName, err error) {
	params := client.Params{
		"ids": strings.Join(util.SliceItoa(ids), ","),
	}
	err = x.Fetch(characterNameURL, &result, params)
	return
}

// ConquerableStationList returns a list of conquerable stations in New Eden.
func (x EveScope) ConquerableStationList() (result xmlapi.ConquerableStationList, err error) {
	err = x.Fetch(conquerableStationListURL, &result, nil)
	return
}

// ErrorList returns a list of error codes and text that the API may return.
func (x EveScope) ErrorList() (result xmlapi.ErrorList, err error) {
	err = x.Fetch(errorListURL, &result, nil)
	return
}

// RefTypes returns a list of transaction types used in the
// Corporation - WalletJournal and Character - WalletJournal.
func (x EveScope) RefTypes() (result xmlapi.RefTypes, err error) {
	err = x.Fetch(refTypesURL, &result, nil)
	return
}

// TypeName returns the names associated with a sequence of typeIDs.
func (x EveScope) TypeName(ids []int) (result xmlapi.TypeName, err error) {
	params := client.Params{
		"ids": strings.Join(util.SliceItoa(ids), ","),
	}
	err = x.Fetch(typeNameURL, &result, params)
	return
}
