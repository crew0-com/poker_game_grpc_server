package api

import (
	"context"
	"github.com/crew_0/poker/internal/grpc/pb"
)

func (s *Server) CreateGameRoom(ctx context.Context, req *pb.CreateGameRoomRequest) (*pb.CreateGameRoomResponse, error) {
	return &pb.CreateGameRoomResponse{
		GameRoom: &pb.GameRoom{},
	}, nil
}
