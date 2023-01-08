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

-- name: UpdateGame :exec
UPDATE games SET game_state = $2, messages = $3 WHERE game_id = $1;

-- name: StartGame :exec
UPDATE games SET has_started = true, started_at = now() WHERE game_id = $1;

-- name: Finish :exec
UPDATE games SET has_finished = true, finished_at = now() WHERE game_id = $1;

-- name: GetGameByRoomId :many
SELECT * FROM games WHERE game_room_id = $1;

-- name: GetActiveGameByRoomId :many
SELECT * FROM games WHERE game_room_id = $1 AND has_finished = false;

