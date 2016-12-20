package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	userAgentName    string  = "go-eve"
	userAgentVersion float32 = 1.0
	userAgentURL     string  = "https://github.com/killmails/go-eve"
)

type (
	// DecoderFunc is used to process the http.Response into a result
	DecoderFunc func(in io.Reader, out interface{}) (err error)

	// ServerURL represents a API server URL
	ServerURL string
	// EndpointURL represents a API endpoint URL
	EndpointURL string

	// UserAgent is the "User-Agent" header requests will be send with
	UserAgent struct {
		Name    string
		Version float32
		URL     string
	}

	// APIKey holds the API Key authentication
	APIKey struct {
		KeyID            int
		VerificationCode string
	}

	// Client is a tiny wrapper over http.Client
	Client struct {
		client  *http.Client
		options *Options
		decoder DecoderFunc
	}

	// Options holds various options for Client
	Options struct {
		ServerURL ServerURL
		UserAgent UserAgent
		APIKey    APIKey
	}

	// Params is used to build the query for http.Request
	Params map[string]string
)

// New returns a Client to handle requests
func New(opts *Options, dec DecoderFunc) (c *Client, err error) {
	if opts.UserAgent.Name == "" {
		opts.UserAgent.Name = userAgentName
	}
	if opts.UserAgent.Version == 0 {
		opts.UserAgent.Version = userAgentVersion
	}
	if opts.UserAgent.URL == "" {
		opts.UserAgent.URL = userAgentURL
	}

	c = &Client{
		client:  &http.Client{},
		options: opts,
		decoder: dec,
	}
	return
}

// Fetch makes a http.Request to the API endpoint and processes the http.Response
func (c *Client) Fetch(endpoint EndpointURL, result interface{}, params Params) (err error) {
	var req *http.Request
	var res *http.Response
	var query url.Values
	var url = fmt.Sprintf("%s%s", c.options.ServerURL, endpoint)

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}

	req.Header.Add("User-Agent", c.options.UserAgent.String())

	query = req.URL.Query()
	if c.options.APIKey.KeyID != 0 {
		query.Set("keyID", fmt.Sprintf("%d", c.options.APIKey.KeyID))
		query.Set("vCode", c.options.APIKey.VerificationCode)
	}
	if params != nil {
		for key, value := range params {
			query.Set(key, value)
		}
	}
	req.URL.RawQuery = query.Encode()

	if res, err = c.client.Do(req); err != nil {
		return
	}
	defer res.Body.Close()

	if 200 > res.StatusCode && res.StatusCode < 400 {
		err = errors.New(res.Status)
		return
	}

	if err = c.decoder(res.Body, &result); err != nil {
		return
	}
	return
}

func (ua UserAgent) String() (u string) {
	u = fmt.Sprintf("%s/%.1f (+%s)", ua.Name, ua.Version, ua.URL)
	return
}
