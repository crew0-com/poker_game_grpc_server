package api

import (
	"context"
	db "github.com/crew_0/poker/database/sqlc"
	"github.com/crew_0/poker/internal/grpc/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateGameRoom(ctx context.Context, req *pb.CreateGameRoomRequest) (*pb.CreateGameRoomResponse, error) {
	playerId, err := uuid.Parse(req.GetRequester().GetToken())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %s", err)
	}

	// Create the player if they don't exist
	_, err = s.store.GetOrCreatePlayer(ctx, db.GetOrCreatePlayerParams{
		PlayerId:   playerId,
		Playername: req.GetRequester().GetName(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting player: %s", err)
	}

	gr, err := s.store.CreateGameRoom(ctx, playerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating game room: %s", err)
	}

	return &pb.CreateGameRoomResponse{
		GameRoom: &pb.GameRoom{
			Id:        gr.GameRoomID.String(),
			CreatedAt: timestamppb.New(gr.CreatedAt),
			CreatedBy: gr.CreatedBy.String(),
			Players:   []*pb.Player{},
		},
	}, nil
}
