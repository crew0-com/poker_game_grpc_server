// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package database

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddPlayerToGameRoom(ctx context.Context, arg AddPlayerToGameRoomParams) (GameRoomPlayer, error)
	CreateGame(ctx context.Context, arg CreateGameParams) (Game, error)
	CreateGameRoom(ctx context.Context, createdBy uuid.UUID) (GameRoom, error)
	CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error)
	FinishGame(ctx context.Context, gameID uuid.UUID) (Game, error)
	GetGame(ctx context.Context, gameID uuid.UUID) (Game, error)
	GetGameRoomAndPlayerRows(ctx context.Context, gameRoomID uuid.UUID) ([]GetGameRoomAndPlayerRowsRow, error)
	GetGamesByRoomId(ctx context.Context, gameRoomID uuid.UUID) ([]Game, error)
	GetPlayer(ctx context.Context, playerID uuid.UUID) (Player, error)
	StartGame(ctx context.Context, gameID uuid.UUID) (Game, error)
	UpdateGame(ctx context.Context, arg UpdateGameParams) (Game, error)
}

var _ Querier = (*Queries)(nil)
