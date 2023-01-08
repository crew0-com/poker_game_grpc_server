-- name: CreatePlayer :one
INSERT INTO players (player_id, name) VALUES ($1, $2) RETURNING *;

-- name: GetPlayer :one
SELECT * FROM players WHERE player_id = $1;