package models

import (
	"context"
	"database/sql"
	"errors"
	"os"

	internalDB "github.com/taiidani/middara/internal/db"
)

var db *sql.DB

func InitDB(ctx context.Context) error {
	switch os.Getenv("DB_TYPE") {
	case "postgres":
		client, err := internalDB.New(ctx, os.Getenv("DATABASE_URL"))
		db = client
		return err
	default:
		return errors.New("unknown DB_TYPE database version specified")
	}
}
