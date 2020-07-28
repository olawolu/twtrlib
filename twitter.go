package twtrlib

import (
	"net/http"

	"github.com/garyburd/go-oauth/oauth"
)

const (
	_Get    = iota
	_Post   = iota
	_Delete = iota
	_PUT    = iota
	url     = "https://api.twitter.com/1.1"
)

// var (
// 	authCredentials oauth.Credentials
// 	// authClient       *oauth.Client
// 	// httpClient       *http.Client
// )

// TwitterClient defines parameters needed to construct an API client
type TwitterClient struct {
	AuthClient      *oauth.Client
	AuthCredentials *oauth.Credentials
	BaseURL         string
	HTTPClient      *http.Client
}

// NewTwitterClient takes in the consumer_key, consumer_secret along with the access_token and secret
//and returns a TwitterAPI struct for connection
func NewTwitterClient(key, secret, token, tokenSecret string) *TwitterClient {
	client := &TwitterClient{
		AuthClient: &oauth.Client{
			Credentials: oauth.Credentials{
				Token:  key,
				Secret: secret,
			},
			TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
			ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authenticate",
			TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		},
		AuthCredentials: &oauth.Credentials{
			Token:  token,
			Secret: tokenSecret,
		},
		BaseURL: url,
		HTTPClient: &http.Client{
			Transport: &http.Transport{},
		},
	}
	return client
}

// NewTwitterClient
