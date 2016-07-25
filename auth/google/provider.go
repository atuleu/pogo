package google

import "net/http"

type Provider struct {
	username string
	password string
	ticket   string
	http     *http.Client
}

const providerString = "ptc"

func NewProvider(username, password string) *Provider {
	return &Provider{
		username: username,
		password: password,
	}
}

func (p *Provider) GetProviderString() string {
	return providerString
}

func (p *Provider) GetAccessToken() string {
	return p.ticket
}

func raiseLoginError(message string) (string, error) {
	return "", &LoginError{message}
}

func (p *Provider) Login() (string, error) {
	return raiseLoginError("Not yet implemented")
}
