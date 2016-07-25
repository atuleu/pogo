package google

import (
	"fmt"
	"net/http"
)

type Provider struct {
	username string
	password string
	ticket   string
	http     *http.Client
}

const providerString = "ptc"

func NewProvider(username, password string) *Provider {

}

func (p *Provider) GetProviderString() {
	return providerString
}

func (p *Provider) GetAccessToken() string {
	return p.ticket
}

func (p *Provider) Login() (string, error) {
	return "", fmt.Errorf("Not yet implemented")
}
