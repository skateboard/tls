package helper

import (
	"io/ioutil"
	"net/http"
)

func GetBodyBytes(resp *http.Response) []byte {
	var respCopy = resp.Body
	bodyBytes, err := ioutil.ReadAll(respCopy)
	if err != nil {
		return nil
	}
	return bodyBytes
}