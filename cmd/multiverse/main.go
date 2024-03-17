package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ciphermountain/deadenz/pkg/multiverse"
	"github.com/ciphermountain/deadenz/pkg/multiverse/service"
)

func main() {
	port := 9090
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service.RegisterMultiverseServer(grpcServer, multiverse.NewMultiverseServer())
	grpcServer.Serve(lis)
}
