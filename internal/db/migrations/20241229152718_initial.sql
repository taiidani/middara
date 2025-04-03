-- +goose Up
-- +goose StatementBegin
CREATE TABLE game (
  id		      SERIAL PRIMARY KEY,
  slug		    CHAR(36) NOT NULL UNIQUE,
  gold		    INTEGER NOT NULL DEFAULT 0,
  page		    INTEGER NOT NULL DEFAULT 1,
  notes       TEXT NOT NULL DEFAULT '',
  created_at  TIMESTAMP NOT NULL,
  updated_at  TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE TABLE character (
  id		       SERIAL PRIMARY KEY,
  game_id 	   INTEGER NOT NULL REFERENCES game,
  name		     VARCHAR(32) NOT NULL,
  xp		       INTEGER NOT NULL DEFAULT 0,
  injured	     BOOLEAN NOT NULL DEFAULT false,
  unselectable BOOLEAN NOT NULL DEFAULT false,
  damage       INTEGER NOT NULL DEFAULT 0,
  created_at   TIMESTAMP NOT NULL,
  updated_at   TIMESTAMP NOT NULL DEFAULT current_timestamp,
  UNIQUE(game_id, name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS character;
DROP TABLE IF EXISTS game;
-- +goose StatementEnd
