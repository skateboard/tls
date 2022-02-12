package service

import (
	"fmt"
	requestProto "github.com/skateboard/tls-client/service/proto"
	"github.com/skateboard/tls-client/service/requests"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartService() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	requestProto.RegisterRequestServiceServer(grpcServer, &requests.ServiceServer{})

	fmt.Println("Serving grpc service:", listen.Addr().String())
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}