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
