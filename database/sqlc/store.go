package database

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type GetOrCreatePlayerParams struct {
	Playername string    `json:"playername"`
	PlayerId   uuid.UUID `json:"player_id"`
}

type Store interface {
	Querier
	GetOrCreatePlayer(ctx context.Context, arg GetOrCreatePlayerParams) (Player, error)
	GetGameRoomWithPlayers(ctx context.Context, gameRoomID uuid.UUID) (GameRoomWithPlayers, error)
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) GetOrCreatePlayer(ctx context.Context, arg GetOrCreatePlayerParams) (player Player, err error) {
	player, err = store.GetPlayer(ctx, arg.PlayerId)
	if err != nil {
		if err != sql.ErrNoRows {
			return
		}

		player, err = store.CreatePlayer(ctx, CreatePlayerParams{
			PlayerID: arg.PlayerId,
			Name:     arg.Playername,
		})

		if err != nil {
			return
		}
	}

	return
}

type GameRoomWithPlayers struct {
	GameRoom
	CreatedBy Player
	Players   []Player
}

func (store *SQLStore) GetGameRoomWithPlayers(ctx context.Context, gameRoomID uuid.UUID) (gameRoom GameRoomWithPlayers, err error) {
	gameRoomRows, err := store.Queries.GetGameRoomAndPlayerRows(ctx, gameRoomID)
	if err != nil {
		return
	}

	createdBy := gameRoomRows[0].CreatedBy
	createdByPlayer, err := store.GetPlayer(ctx, createdBy)
	if err != nil {
		return
	}

	gameRoom.GameRoomID = gameRoomRows[0].GameRoomID
	gameRoom.CreatedAt = gameRoomRows[0].CreatedAt
	gameRoom.ClosedAt = gameRoomRows[0].ClosedAt
	gameRoom.CreatedBy = createdByPlayer

	for _, row := range gameRoomRows {
		gameRoom.Players = append(gameRoom.Players, Player{
			PlayerID: row.PlayerID,
			Name:     row.Name,
		})
	}

	return gameRoom, nil
}
