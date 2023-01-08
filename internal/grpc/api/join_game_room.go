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

func (s *Server) JoinGameRoom(ctx context.Context, req *pb.JoinGameRoomRequest) (*pb.JoinGameRoomResponse, error) {
	playerId, err := uuid.Parse(req.GetRequester().GetToken())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %s", err)
	}
	gameRoomId, err := uuid.Parse(req.GetGameRoomId())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid game room", err)
	}

	// Create the player if they don't exist
	_, err = s.store.GetOrCreatePlayer(ctx, db.GetOrCreatePlayerParams{
		PlayerId:   playerId,
		Playername: req.GetRequester().GetName(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting player: %s", err)
	}
	_, err = s.store.AddPlayerToGameRoom(ctx, db.AddPlayerToGameRoomParams{
		PlayerID:   playerId,
		GameRoomID: gameRoomId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error adding player to game room: %s", err)
	}

	gameRoom, err := s.store.GetGameRoomAndPlayers(ctx, gameRoomId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting game room: %s", err)
	}

	response := &pb.JoinGameRoomResponse{
		GameRoom: &pb.GameRoom{
			Id: gameRoom.GameRoomID.String(),
			Players: func() []*pb.Player {
				players := make([]*pb.Player, len(gameRoom.Players))
				for i, player := range gameRoom.Players {
					players[i] = &pb.Player{
						Uuid: player.PlayerID.String(),
						Name: player.Name,
					}
				}
				return players
			}(),
			CreatedBy: &pb.Player{
				Uuid: gameRoom.CreatedBy.PlayerID.String(),
				Name: gameRoom.CreatedBy.Name,
			},
			CreatedAt: timestamppb.New(gameRoom.CreatedAt),
		},
	}

	if gameRoom.ClosedAt.Valid {
		response.GameRoom.ClosedAt = timestamppb.New(gameRoom.ClosedAt.Time)
	}

	return response, nil
}
