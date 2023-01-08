-- name: CreateGame :one
INSERT INTO games(
                  game_state,
                  messages,
                  game_room_id
) values (
                  $1,
                  $2,
                  $3
) returning *;

-- name: GetGame :one
SELECT * FROM games WHERE game_id = $1;

-- name: UpdateGame :one
UPDATE games SET game_state = $2, messages = $3 WHERE game_id = $1 returning *;

-- name: StartGame :one
UPDATE games SET has_started = true, started_at = now() WHERE game_id = $1 returning *;

-- name: FinishGame :one
UPDATE games SET has_finished = true, finished_at = now() WHERE game_id = $1 returning *;

-- name: GetGamesByRoomId :many
SELECT * FROM games WHERE game_room_id = $1;

