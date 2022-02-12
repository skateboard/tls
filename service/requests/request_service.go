package requests

import (
	"context"
	"errors"
	"github.com/skateboard/tls-client/client"
	"github.com/skateboard/tls-client/helper"
	requestProto "github.com/skateboard/tls-client/service/proto"
)

type ServiceServer struct {
	requestProto.UnimplementedRequestServiceServer
}

func (s ServiceServer) Request(ctx context.Context, request *requestProto.NewRequest) (*requestProto.Response, error) {
	browser := client.Browsers[request.GetBrowser()]
	if browser.UserAgent == "" {
		return nil, errors.New("browser fingerprint not found")
	}

	requester := client.Requester{
		Browser: browser,
		Proxy:   request.GetProxy(),
	}
	requester.Initialize()

	switch request.Method {
	case "GET":
		response, err := requester.GET(request.Url)
		if err != nil {
			return nil, err
		}
		bodyBytes := helper.GetBodyBytes(response)

		headers := make(map[string]string)
		for name, values := range response.Header {
			for _, value := range values {
				headers[name] = value
			}
		}

		//#todo: add cookie parsing


		return &requestProto.Response{
			Status:  int32(response.StatusCode),
			Body:    string(bodyBytes),
			Headers: headers,
			Cookies: nil,
		}, nil
	}

	return nil, errors.New("failed to find method!")

}
