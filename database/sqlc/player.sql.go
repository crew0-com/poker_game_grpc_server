// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: player.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO players (player_id, name) VALUES ($1, $2) RETURNING player_id, name
`

type CreatePlayerParams struct {
	PlayerID uuid.UUID `json:"player_id"`
	Name     string    `json:"name"`
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.queryRow(ctx, q.createPlayerStmt, createPlayer, arg.PlayerID, arg.Name)
	var i Player
	err := row.Scan(&i.PlayerID, &i.Name)
	return i, err
}

const getPlayer = `-- name: GetPlayer :one
SELECT player_id, name FROM players WHERE player_id = $1
`

func (q *Queries) GetPlayer(ctx context.Context, playerID uuid.UUID) (Player, error) {
	row := q.queryRow(ctx, q.getPlayerStmt, getPlayer, playerID)
	var i Player
	err := row.Scan(&i.PlayerID, &i.Name)
	return i, err
}
