// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: game.sql

package database

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

const createGame = `-- name: CreateGame :one
INSERT INTO games(
                  id,
                  game_room_id,
                  game_state,
                  messages
) values (
                  $1,
                  $2,
                  $3,
                  $4
) returning id, game_room_id, game_state, messages, started_at, finished_at, is_finished, has_started
`

type CreateGameParams struct {
	ID         uuid.UUID       `json:"id"`
	GameRoomID uuid.UUID       `json:"game_room_id"`
	GameState  json.RawMessage `json:"game_state"`
	Messages   json.RawMessage `json:"messages"`
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.queryRow(ctx, q.createGameStmt, createGame,
		arg.ID,
		arg.GameRoomID,
		arg.GameState,
		arg.Messages,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.GameRoomID,
		&i.GameState,
		&i.Messages,
		&i.StartedAt,
		&i.FinishedAt,
		&i.IsFinished,
		&i.HasStarted,
	)
	return i, err
}

const finish = `-- name: Finish :exec
UPDATE games SET is_finished = true, finished_at = now() WHERE id = $1
`

func (q *Queries) Finish(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.finishStmt, finish, id)
	return err
}

const getActiveGameByRoomId = `-- name: GetActiveGameByRoomId :many
SELECT id, game_room_id, game_state, messages, started_at, finished_at, is_finished, has_started FROM games WHERE game_room_id = $1 AND is_finished = false
`

func (q *Queries) GetActiveGameByRoomId(ctx context.Context, gameRoomID uuid.UUID) ([]Game, error) {
	rows, err := q.query(ctx, q.getActiveGameByRoomIdStmt, getActiveGameByRoomId, gameRoomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Game{}
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.GameRoomID,
			&i.GameState,
			&i.Messages,
			&i.StartedAt,
			&i.FinishedAt,
			&i.IsFinished,
			&i.HasStarted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGame = `-- name: GetGame :one
SELECT id, game_room_id, game_state, messages, started_at, finished_at, is_finished, has_started FROM games WHERE id = $1
`

func (q *Queries) GetGame(ctx context.Context, id uuid.UUID) (Game, error) {
	row := q.queryRow(ctx, q.getGameStmt, getGame, id)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.GameRoomID,
		&i.GameState,
		&i.Messages,
		&i.StartedAt,
		&i.FinishedAt,
		&i.IsFinished,
		&i.HasStarted,
	)
	return i, err
}

const getGameByRoomId = `-- name: GetGameByRoomId :many
SELECT id, game_room_id, game_state, messages, started_at, finished_at, is_finished, has_started FROM games WHERE game_room_id = $1
`

func (q *Queries) GetGameByRoomId(ctx context.Context, gameRoomID uuid.UUID) ([]Game, error) {
	rows, err := q.query(ctx, q.getGameByRoomIdStmt, getGameByRoomId, gameRoomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Game{}
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.GameRoomID,
			&i.GameState,
			&i.Messages,
			&i.StartedAt,
			&i.FinishedAt,
			&i.IsFinished,
			&i.HasStarted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const startGame = `-- name: StartGame :exec
UPDATE games SET has_started = true, started_at = now() WHERE id = $1
`

func (q *Queries) StartGame(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.startGameStmt, startGame, id)
	return err
}

const updateGame = `-- name: UpdateGame :exec
UPDATE games SET game_state = $2, messages = $3 WHERE id = $1
`

type UpdateGameParams struct {
	ID        uuid.UUID       `json:"id"`
	GameState json.RawMessage `json:"game_state"`
	Messages  json.RawMessage `json:"messages"`
}

func (q *Queries) UpdateGame(ctx context.Context, arg UpdateGameParams) error {
	_, err := q.exec(ctx, q.updateGameStmt, updateGame, arg.ID, arg.GameState, arg.Messages)
	return err
}
