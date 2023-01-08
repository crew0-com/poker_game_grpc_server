-- name: GetGameMessages :many
-- desc: Returns all game messages
SELECT * FROM game_messages WHERE game_id = $1 ORDER BY message_id DESC LIMIT 100;

-- name: PaginatedGameMessages :many
-- desc: Returns a page of game messages
SELECT * FROM game_messages WHERE game_id = $1 ORDER BY message_id DESC LIMIT $2 OFFSET $3;

-- name: PaginatedGameMessageByGameRoom :many
SELECT * FROM game_messages WHERE game_room_id = $1 ORDER BY message_id DESC LIMIT $2 OFFSET $3;

-- name: AddGameMessage :one
-- desc: Adds a game message
INSERT INTO game_messages (game_id, game_room_id,  message)
VALUES ($1, $2, $3) returning *;
