-- name: CreateGameRoom :one
INSERT INTO game_rooms(created_by) values ($1) returning *;

-- name: GetGameRoomAndPlayerRows :many
SELECT game_rooms.game_room_id AS gameroom_id, game_rooms.created_at, game_rooms.created_by, game_rooms.closed_at, players.name, players.player_id
FROM game_rooms
         JOIN game_room_players ON game_room_players.game_room_id = game_rooms.game_room_id
         JOIN players ON players.player_id = game_room_players.player_id
WHERE game_rooms.game_room_id = $1;

-- name: AddGameRoomPlayer :one
INSERT INTO game_room_players(game_room_id, player_id) values ($1, $2) returning *;
