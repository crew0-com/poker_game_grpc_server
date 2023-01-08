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
                  game_state,
                  messages,
                  game_room_id
) values (
                  $1,
                  $2,
                  $3
) returning game_id, game_room_id, game_state, messages, started_at, finished_at, has_finished, has_started
`

type CreateGameParams struct {
	GameState  json.RawMessage `json:"game_state"`
	Messages   json.RawMessage `json:"messages"`
	GameRoomID uuid.UUID       `json:"game_room_id"`
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.queryRow(ctx, q.createGameStmt, createGame, arg.GameState, arg.Messages, arg.GameRoomID)
	var i Game
	err := row.Scan(
		&i.GameID,
		&i.GameRoomID,
		&i.GameState,
		&i.Messages,
		&i.StartedAt,
		&i.FinishedAt,
		&i.HasFinished,
		&i.HasStarted,
	)
	return i, err
}

const finishGame = `-- name: FinishGame :one
UPDATE games SET has_finished = true, finished_at = now() WHERE game_id = $1 returning game_id, game_room_id, game_state, messages, started_at, finished_at, has_finished, has_started
`

func (q *Queries) FinishGame(ctx context.Context, gameID uuid.UUID) (Game, error) {
	row := q.queryRow(ctx, q.finishGameStmt, finishGame, gameID)
	var i Game
	err := row.Scan(
		&i.GameID,
		&i.GameRoomID,
		&i.GameState,
		&i.Messages,
		&i.StartedAt,
		&i.FinishedAt,
		&i.HasFinished,
		&i.HasStarted,
	)
	return i, err
}

const getActiveGame = `-- name: GetActiveGame :one
SELECT games.game_id, games.game_room_id, games.game_state, games.messages, games.started_at, games.finished_at, games.has_finished, games.has_started
FROM game_rooms
         JOIN active_games ON active_games.game_room_id = game_rooms.game_room_id
         JOIN games ON games.game_id = active_games.game_id
WHERE game_rooms.game_room_id = $1
`

func (q *Queries) GetActiveGame(ctx context.Context, gameRoomID uuid.UUID) (Game, error) {
	row := q.queryRow(ctx, q.getActiveGameStmt, getActiveGame, gameRoomID)
	var i Game
	err := row.Scan(
		&i.GameID,
		&i.GameRoomID,
		&i.GameState,
		&i.Messages,
		&i.StartedAt,
		&i.FinishedAt,
		&i.HasFinished,
		&i.HasStarted,
	)
	return i, err
}

const getGame = `-- name: GetGame :one
SELECT game_id, game_room_id, game_state, messages, started_at, finished_at, has_finished, has_started FROM games WHERE game_id = $1
`

func (q *Queries) GetGame(ctx context.Context, gameID uuid.UUID) (Game, error) {
	row := q.queryRow(ctx, q.getGameStmt, getGame, gameID)
	var i Game
	err := row.Scan(
		&i.GameID,
		&i.GameRoomID,
		&i.GameState,
		&i.Messages,
		&i.StartedAt,
		&i.FinishedAt,
		&i.HasFinished,
		&i.HasStarted,
	)
	return i, err
}

const getGamesByRoomId = `-- name: GetGamesByRoomId :many
SELECT game_id, game_room_id, game_state, messages, started_at, finished_at, has_finished, has_started FROM games WHERE game_room_id = $1
`

func (q *Queries) GetGamesByRoomId(ctx context.Context, gameRoomID uuid.UUID) ([]Game, error) {
	rows, err := q.query(ctx, q.getGamesByRoomIdStmt, getGamesByRoomId, gameRoomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Game{}
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.GameID,
			&i.GameRoomID,
			&i.GameState,
			&i.Messages,
			&i.StartedAt,
			&i.FinishedAt,
			&i.HasFinished,
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

const setActiveGame = `-- name: SetActiveGame :one
INSERT INTO active_games(game_room_id, game_id)
SELECT games.game_room_id, games.game_id FROM games where games.game_id = $1 returning game_room_id, game_id
`

func (q *Queries) SetActiveGame(ctx context.Context, gameID uuid.UUID) (ActiveGame, error) {
	row := q.queryRow(ctx, q.setActiveGameStmt, setActiveGame, gameID)
	var i ActiveGame
	err := row.Scan(&i.GameRoomID, &i.GameID)
	return i, err
}

const startGame = `-- name: StartGame :one
UPDATE games SET has_started = true, started_at = now() WHERE game_id = $1 returning game_id, game_room_id, game_state, messages, started_at, finished_at, has_finished, has_started
`

func (q *Queries) StartGame(ctx context.Context, gameID uuid.UUID) (Game, error) {
	row := q.queryRow(ctx, q.startGameStmt, startGame, gameID)
	var i Game
	err := row.Scan(
		&i.GameID,
		&i.GameRoomID,
		&i.GameState,
		&i.Messages,
		&i.StartedAt,
		&i.FinishedAt,
		&i.HasFinished,
		&i.HasStarted,
	)
	return i, err
}

const unsetActiveGame = `-- name: UnsetActiveGame :one
DELETE FROM active_games WHERE game_id = $1 returning active_games.game_id, active_games.game_room_id
`

type UnsetActiveGameRow struct {
	GameID     uuid.UUID `json:"game_id"`
	GameRoomID uuid.UUID `json:"game_room_id"`
}

func (q *Queries) UnsetActiveGame(ctx context.Context, gameID uuid.UUID) (UnsetActiveGameRow, error) {
	row := q.queryRow(ctx, q.unsetActiveGameStmt, unsetActiveGame, gameID)
	var i UnsetActiveGameRow
	err := row.Scan(&i.GameID, &i.GameRoomID)
	return i, err
}

const updateGame = `-- name: UpdateGame :one
UPDATE games SET game_state = $2, messages = $3 WHERE game_id = $1 returning game_id, game_room_id, game_state, messages, started_at, finished_at, has_finished, has_started
`

type UpdateGameParams struct {
	GameID    uuid.UUID       `json:"game_id"`
	GameState json.RawMessage `json:"game_state"`
	Messages  json.RawMessage `json:"messages"`
}

func (q *Queries) UpdateGame(ctx context.Context, arg UpdateGameParams) (Game, error) {
	row := q.queryRow(ctx, q.updateGameStmt, updateGame, arg.GameID, arg.GameState, arg.Messages)
	var i Game
	err := row.Scan(
		&i.GameID,
		&i.GameRoomID,
		&i.GameState,
		&i.Messages,
		&i.StartedAt,
		&i.FinishedAt,
		&i.HasFinished,
		&i.HasStarted,
	)
	return i, err
}
