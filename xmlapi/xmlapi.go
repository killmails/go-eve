package xmlapi

import (
	"encoding/xml"
	"io"

	"github.com/killmails/go-eve/client"
)

const (
	TranquilityUrl client.ServerUrl = "https://api.eveonline.com"
	SingularityUrl client.ServerUrl = "https://api.testeveonline.com"
)

type (
	XmlApi struct {
		*client.Client
	}
)

// decoder converts XML response to schema result
func decoder(in io.Reader, out interface{}) (err error) {
	err = xml.NewDecoder(in).Decode(out)
	return
}

// New returns a XmlApi client to access EVE Online XML API
func New(opts *client.Options) (x *XmlApi, err error) {
	if opts.ServerUrl == "" {
		opts.ServerUrl = TranquilityUrl
	}

	c, err := client.New(opts, decoder)
	x = &XmlApi{c}
	return
}
