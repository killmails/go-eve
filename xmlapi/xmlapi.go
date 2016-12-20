package xmlapi

import (
	"encoding/xml"
	"io"

	"github.com/killmails/go-eve/client"
)

const (
	// TranquilityURL is the address of Tranquility API server
	TranquilityURL client.ServerURL = "https://api.eveonline.com"
	// SingularityURL is the address of Singularity API server
	SingularityURL client.ServerURL = "https://api.testeveonline.com"
)

type (
	// XMLAPI is the API interface
	XMLAPI struct {
		*client.Client
	}
)

// decoder converts XML response to schema result
func decoder(in io.Reader, out interface{}) (err error) {
	err = xml.NewDecoder(in).Decode(out)
	return
}

// New returns a XMLAPI client to access EVE Online XML API
func New(opts *client.Options) (x *XMLAPI, err error) {
	if opts.ServerURL == "" {
		opts.ServerURL = TranquilityURL
	}

	c, err := client.New(opts, decoder)
	x = &XMLAPI{c}
	return
}
