package api

import (
	"context"
	"github.com/crew_0/poker/internal/grpc/pb"
)

func (s *Server) Hello(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: "Hello " + req.Name,
	}, nil
}