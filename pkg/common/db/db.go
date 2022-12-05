package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	DB *sql.DB
}

func Init(dbPath string) (*Connection, error) {
	db, err := initSqlite(dbPath)

	if err != nil {
		return &Connection{}, err
	}

	ConnectionInstance := &Connection{
		db,
	}

	return ConnectionInstance, nil
}

func initSqlite(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	return db, nil
}
