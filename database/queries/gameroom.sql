-- name: CreateGameRoom :one
INSERT INTO game_rooms(created_by) values ($1) returning *;

-- name: GetGameRoom :one
SELECT * FROM game_rooms WHERE game_room_id = $1;

-- name: AddGameRoomPlayer :one
INSERT INTO game_room_players(game_room_id, player_id) values ($1, $2) returning *;
