package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	DB *sql.DB
}

func Init() (*Connection, error) {
	db, err := initSqlite()

	if err != nil {
		return &Connection{}, err
	}

	ConnectionInstance := &Connection{
		db,
	}

	return ConnectionInstance, nil
}

func initSqlite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}
