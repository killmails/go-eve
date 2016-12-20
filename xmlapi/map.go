package xmlapi

import (
	"github.com/killmails/go-eve/client"
	"github.com/killmails/go-eve/schemas/xmlapi"
)

const (
	facWarSystemsURL client.EndpointURL = "/map/FacWarSystems.xml.aspx"
	jumpsURL         client.EndpointURL = "/map/Jumps.xml.aspx"
	killsURL         client.EndpointURL = "/map/Kills.xml.aspx"
	sovereigntyURL   client.EndpointURL = "/map/Sovereignty.xml.aspx"
)

type (
	// MapScope is the API/map interface
	MapScope struct {
		*XMLAPI
	}
)

// Map is the scope for /map/* API requests
func (x *XMLAPI) Map() (r MapScope) {
	r = MapScope{x}
	return
}

// FacWarSystems returns a list of contestable solarsystems and the NPC faction
// currently occupying them
func (x MapScope) FacWarSystems() (result xmlapi.FacWarSystems, err error) {
	err = x.Fetch(facWarSystemsURL, &result, nil)
	return
}

// Jumps returns the number of jumps in solarsystems within the last hour,
// doesn’t include wormhole space.
func (x MapScope) Jumps() (result xmlapi.Jumps, err error) {
	err = x.Fetch(jumpsURL, &result, nil)
	return
}

// Kills provides number of ship, pod and NPC kills per solar system
// within the last hour, doesn’t include wormhole space.
func (x MapScope) Kills() (result xmlapi.Kills, err error) {
	err = x.Fetch(killsURL, &result, nil)
	return
}

// Sovereignty provides Sovereignty / Ownership information for solar systems
// in New Eden excluding wormhole space.
func (x MapScope) Sovereignty() (result xmlapi.Sovereignty, err error) {
	err = x.Fetch(sovereigntyURL, &result, nil)
	return
}
