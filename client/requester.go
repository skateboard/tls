package client

import (
	cclient "github.com/Humphryyy/cclientwtf"
	"net/http"
)

type Requester struct {
	Browser Browser
	Proxy 	string

	Client http.Client
}

func (r *Requester) Initialize()  {
	client, err := cclient.NewClient(r.Browser.ClientID, r.Proxy)
	if err != nil {
		return
	}

	r.Client = client
}

func (r Requester) request(method string, url string, body []byte) ([]byte, error) {
	return nil, nil
}

func (r *Requester) GET(url string) (*http.Response, error) {
	response, responseError := r.Client.Get(url)
	if responseError != nil {
		return response, responseError
	}

	return response, nil
}