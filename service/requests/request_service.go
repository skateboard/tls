package requests

import (
	"context"
	"github.com/skateboard/tls-client/client"
	"github.com/skateboard/tls-client/helper"
	requestProto "github.com/skateboard/tls-client/service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceServer struct {
	requestProto.UnimplementedRequestServiceServer
}


func (s ServiceServer) Request(ctx context.Context, request *requestProto.NewRequest) (*requestProto.Response, error) {
	tlsClient := client.CreateClient(request.GetProxy())
	if tlsClient == nil {
		return nil, status.Error(codes.Internal, "failed to create tls client")
	}
	
	requester := client.Requester{
		Client: tlsClient,
	}

	switch request.Method {
	case "GET":
		response, err := requester.GET(request.Url)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		bodyBytes := helper.GetBodyBytes(response)

		var headers map[string]string
		for name, values := range response.Header {
			for _, value := range values {
				headers[name] = value
			}
		}

		return &requestProto.Response{
			Status:  response.Status,
			Body:    string(bodyBytes),
			Headers: headers,
			Cookies: nil,
		}, nil
	}
	
	
	return nil, status.Errorf(codes.Unimplemented, "failed to find method %s", request.Method)
}
