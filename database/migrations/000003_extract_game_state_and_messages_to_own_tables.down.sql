ALTER TABLE games DROP COLUMN IF EXISTS state_version;
ALTER TABLE games ADD COLUMN game_state jsonb NOT NULL DEFAULT '{}'::jsonb;
DROP TABLE IF EXISTS game_state;


ALTER TABLE games ADD COLUMN messages jsonb NOT NULL DEFAULT '{}'::jsonb;
DROP TABLE IF EXISTS game_messages;
