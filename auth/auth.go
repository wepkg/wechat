package auth

import "net/url"

// AccessToken ..
type AccessToken struct{}

const authURL = "/cgi-bin/token"

// Auth ..
type Auth interface {
	GetToken() string
}

type auth struct {
	Appid  string
	Secret string
}

func (a *auth) ClientCredential() string {
	u := url.Values{}
	u.Add("grant_type", "client_credential")
	u.Add("appid", a.Appid)
	u.Add("secret", a.Secret)
	return u.Encode()
}
func (a *auth) GetToken() AccessToken {
	a.
	return
}
