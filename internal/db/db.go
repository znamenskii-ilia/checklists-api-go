package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func New(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS checklists (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL UNIQUE,
		tasks TEXT NOT NULL DEFAULT '[]',
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		return nil, err
	}

	return db, nil
}
