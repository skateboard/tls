package client

import "net/http"

type Requester struct {
	Client *http.Client
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