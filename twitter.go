package twitlib

import (
	"github.com/garyburd/go-oauth/oauth"
)

var oauthCredentials oauth.Credentials

// TwitterAPI defines parameters needed to construct an API client
type TwitterAPI struct {
	client      oauth.Client
	Credentials *oauth.Credentials
}

// NewTwitterAPI takes in an access token and secret and returns a TwitterAPI struct for connection
func NewTwitterAPI(token, tokenSecret string) *TwitterAPI {
	api := &TwitterAPI{
		client: oauth.Client{
			Credentials: oauthCredentials,
			TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
			ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authenticate",
			TokenRequestURI: "https://api.twitter.com/oauth/access_token",
		},
		Credentials: &oauth.Credentials{
			Token: token,
			Secret: tokenSecret,
		},
	}
	return api
}

