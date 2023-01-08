// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package database

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Game struct {
	GameID      uuid.UUID       `json:"game_id"`
	GameRoomID  uuid.UUID       `json:"game_room_id"`
	GameState   json.RawMessage `json:"game_state"`
	Messages    json.RawMessage `json:"messages"`
	StartedAt   sql.NullTime    `json:"started_at"`
	FinishedAt  sql.NullTime    `json:"finished_at"`
	HasFinished bool            `json:"has_finished"`
	HasStarted  bool            `json:"has_started"`
}

type GameRoom struct {
	GameRoomID uuid.UUID    `json:"game_room_id"`
	CreatedBy  uuid.UUID    `json:"created_by"`
	CreatedAt  time.Time    `json:"created_at"`
	ClosedAt   sql.NullTime `json:"closed_at"`
}

type GameRoomPlayer struct {
	GameRoomID uuid.UUID `json:"game_room_id"`
	PlayerID   uuid.UUID `json:"player_id"`
}

type Player struct {
	PlayerID uuid.UUID `json:"player_id"`
	Name     string    `json:"name"`
}
