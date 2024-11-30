package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func New() (*sql.DB, error) {
	// TODO: GORM
	// TODO: DB path from config/.env
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, fmt.Errorf("cant open database: %w", err)
	}
	return db, err
}
