package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	UserAgentName    string  = "go-eve"
	UserAgentVersion float32 = 1.0
	UserAgentUrl     string  = "https://github.com/killmails/go-eve"
)

type (
	DecoderFunc func(in io.Reader, out interface{}) (err error)

	ServerUrl   string
	EndpointUrl string

	UserAgent struct {
		Name    string
		Version float32
		Url     string
	}

	ApiKey struct {
		KeyId            int
		VerificationCode string
	}

	Client struct {
		client  *http.Client
		options *Options
		decoder DecoderFunc
	}

	Options struct {
		ServerUrl ServerUrl
		UserAgent UserAgent
		ApiKey    ApiKey
	}

	Params map[string]string
)

var (
	ErrApiKeyInvalidKeyId            = errors.New("You must provide a valid API Key ID!")
	ErrApiKeyInvalidVerificationCode = errors.New("You must provide a valid API Verification Code!")
)

// New returns a Client to handle requests
func New(opts *Options, dec DecoderFunc) (c *Client, err error) {
	if opts.UserAgent.Name == "" {
		opts.UserAgent.Name = UserAgentName
	}
	if opts.UserAgent.Version == 0 {
		opts.UserAgent.Version = UserAgentVersion
	}
	if opts.UserAgent.Url == "" {
		opts.UserAgent.Url = UserAgentUrl
	}

	c = &Client{
		client:  &http.Client{},
		options: opts,
		decoder: dec,
	}
	return
}

func (c *Client) Fetch(endpoint EndpointUrl, result interface{}, params Params) (err error) {
	var req *http.Request
	var res *http.Response
	var query url.Values
	var url string = fmt.Sprintf("%s%s", c.options.ServerUrl, endpoint)

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return
	}

	req.Header.Add("User-Agent", c.options.UserAgent.String())

	query = req.URL.Query()
	if c.options.ApiKey.KeyId != 0 {
		query.Set("keyID", fmt.Sprintf("%d", c.options.ApiKey.KeyId))
		query.Set("vCode", c.options.ApiKey.VerificationCode)
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
	u = fmt.Sprintf("%s/%.1f (+%s)", ua.Name, ua.Version, ua.Url)
	return
}
