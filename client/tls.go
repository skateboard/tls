package client

import (
	cclient "github.com/Humphryyy/cclientwtf"
	utls "github.com/refraction-networking/utls"
	"net/http"
)

func CreateClient(proxy string) *http.Client {
	client, err := cclient.NewClient(utls.HelloChrome_83, proxy)
	if err != nil {
		return nil
	}

	return &client
}