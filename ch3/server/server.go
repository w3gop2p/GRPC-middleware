package main

import (
	"context"
	"fmt"
	shipping "github.com/w3gop2p/GRPC-middleware/ch3/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	shipping.UnimplementedShippingServiceServer
}

func (s *server) Create(ctx context.Context, in *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	time.Sleep(2 * time.Second) // simulated delay
	return &shipping.CreateShippingResponse{ShippingId: 1243}, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	shipping.RegisterShippingServiceServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}
