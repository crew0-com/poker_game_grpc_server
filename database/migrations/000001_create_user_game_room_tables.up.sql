CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS players(
    player_id uuid NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (player_id)
);

CREATE TABLE IF NOT EXISTS game_rooms(
    game_room_id uuid NOT NULL  DEFAULT uuid_generate_v4(),
    created_by uuid NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    closed_at TIMESTAMP,
    PRIMARY KEY (game_room_id),
    FOREIGN KEY (created_by) REFERENCES players(player_id)
);

CREATE TABLE IF NOT EXISTS game_room_players(
    game_room_id uuid NOT NULL,
    player_id uuid NOT NULL,
    PRIMARY KEY (game_room_id, player_id),
    FOREIGN KEY (game_room_id) REFERENCES game_rooms(game_room_id),
    FOREIGN KEY (player_id) REFERENCES players(player_id)
);

CREATE TABLE IF NOT EXISTS games(
    game_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    game_room_id uuid NOT NULL,
    game_state jsonb NOT NULL,
    messages jsonb  NOT NULL,
    started_at timestamp,
    finished_at timestamp,
    is_finished boolean NOT NULL default false,
    has_started boolean NOT NULL default false,
    PRIMARY KEY (game_id),
    FOREIGN KEY (game_room_id) REFERENCES game_rooms(game_room_id)
);



