ALTER TABLE games DROP COLUMN IF EXISTS game_state;
ALTER TABLE games ADD COLUMN state_version INT NOT NULL DEFAULT 0;

CREATE TABLE game_state
(
    game_id UUID  NOT NULL,
    state   JSONB NOT NULL,
    version INT   NOT NULL,
    primary key (game_id),
    foreign key (game_id) references games (game_id)
);

ALTER TABLE games DROP COLUMN IF EXISTS messages;

CREATE TABLE game_messages
(
    game_id UUID  NOT NULL,
    game_room_id UUID NOT NULL,
    message_id BIGSERIAL NOT NULL,
    message JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    primary key (game_id, message_id),
    foreign key (game_id) references games (game_id),
    foreign key (game_room_id) references game_rooms (game_room_id)
);



