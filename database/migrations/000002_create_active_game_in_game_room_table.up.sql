CREATE TABLE IF NOT EXISTS active_games(
    game_room_id uuid UNIQUE NOT NULL,
    game_id uuid UNIQUE NOT NULL,
    PRIMARY KEY (game_room_id, game_id),
    FOREIGN KEY (game_room_id) REFERENCES game_rooms(game_room_id),
    FOREIGN KEY (game_id) REFERENCES games(game_id)
);