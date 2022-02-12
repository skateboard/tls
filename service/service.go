package service

import (
	"fmt"
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

	fmt.Println("Serving grpc service:", listen.Addr().String())
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}