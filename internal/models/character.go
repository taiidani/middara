package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Character struct {
	ID           int
	GameID       int
	Name         string    `json:"name"`
	XP           int       `json:"xp"`
	Damage       int       `json:"damage"`
	Injured      bool      `json:"injured"`
	Unselectable bool      `json:"unselectable"`
	CreatedAt    time.Time `json:"created_at"`
}

func GetCharactersForGame(ctx context.Context, id int) ([]Character, error) {
	ret := []Character{}
	rows, err := db.QueryContext(ctx, `
SELECT
	id,
	game_id,
	name,
	xp,
	damage,
	injured,
	unselectable,
	created_at
FROM character
WHERE game_id = $1
`, id)
	if err != nil {
		return ret, fmt.Errorf("could not search for game characters: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		add := Character{}
		err = rows.Scan(
			&add.ID,
			&add.GameID,
			&add.Name,
			&add.XP,
			&add.Damage,
			&add.Injured,
			&add.Unselectable,
			&add.CreatedAt,
		)
		if err != nil {
			return ret, fmt.Errorf("could not get character: %w", err)
		}

		ret = append(ret, add)
	}
	if err := rows.Err(); err != nil {
		return ret, err
	}

	return ret, nil
}

func (c *Character) Validate() error {
	if len(c.Name) == 0 {
		return errors.New("name must not be empty")
	}

	return nil
}

func (c *Character) save(ctx context.Context, tx *sql.Tx) error {
	if c.ID == 0 {
		err := c.insert(ctx, tx)
		if err != nil {
			return fmt.Errorf("failed to insert character %q: %w", c.Name, err)
		}
	} else {
		err := c.update(ctx, tx)
		if err != nil {
			return fmt.Errorf("failed to update character %q: %w", c.Name, err)
		}
	}

	return nil
}

func (c *Character) insert(ctx context.Context, tx *sql.Tx) error {
	stmt, err := tx.PrepareContext(ctx, `
INSERT INTO character (game_id, name, xp, damage, injured, unselectable, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
RETURNING id
`)

	if err != nil {
		return err
	}

	c.CreatedAt = time.Now().UTC()
	return stmt.QueryRowContext(ctx,
		&c.GameID,
		&c.Name,
		&c.XP,
		&c.Damage,
		&c.Injured,
		&c.Unselectable,
		&c.CreatedAt,
	).Scan(&c.ID)
}

func (c *Character) update(ctx context.Context, tx *sql.Tx) error {
	stmt, err := tx.PrepareContext(ctx, `
UPDATE character SET
	name = $2,
	xp = $3,
	damage = $4,
	injured = $5,
	unselectable = $6,
	updated_at = NOW()
WHERE id = $1
RETURNING id
		`)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		&c.ID,
		&c.Name,
		&c.XP,
		&c.Damage,
		&c.Injured,
		&c.Unselectable,
	)
	return err
}
