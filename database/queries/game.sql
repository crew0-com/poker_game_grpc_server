-- name: CreateGame :one
INSERT INTO games(
                  game_room_id
) values (
                  $1
) returning *;

-- name: GetGame :one
SELECT * FROM games WHERE game_id = $1;

-- name: StartGame :one
UPDATE games SET has_started = true, started_at = now() WHERE game_id = $1 returning *;

-- name: FinishGame :one
UPDATE games SET has_finished = true, finished_at = now() WHERE game_id = $1 returning *;

-- name: GetGamesByRoomId :many
SELECT * FROM games WHERE game_room_id = $1;

-- name: SetActiveGame :one
INSERT INTO active_games(game_room_id, game_id)
SELECT games.game_room_id, games.game_id FROM games where games.game_id = $1 returning *;

-- name: GetActiveGame :one
SELECT games.*
FROM game_rooms
         JOIN active_games ON active_games.game_room_id = game_rooms.game_room_id
         JOIN games ON games.game_id = active_games.game_id
WHERE game_rooms.game_room_id = $1;

-- name: UnsetActiveGame :one
DELETE FROM active_games WHERE game_id = $1 returning active_games.game_id, active_games.game_room_id;

