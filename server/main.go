package main

import (
	"context"
	"log"
	"net"

	"github.com/dimiro1/grpc-experiment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type sumServer struct{}

func (sumServer) Sum(_ context.Context, r *service.SumRequest) (*service.SumReply, error) {
	return &service.SumReply{
		Sum: r.A + r.B,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterCalculatorServer(s, &sumServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
