package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DBCon *sql.DB
)

type Person struct {
	Id       string
	Name     string
	LastName string
}

func InitDb(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}
	return db
}

func CreateTable(db *sql.DB) {
	// Create table if not exists
	personTable := `
	CREATE TABLE IF NOT EXISTS person(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT,
		LastName Text,
		created_at Date
	)
	`
	_, err := db.Exec(personTable)
	if err != nil {
		panic(err)
	}
}
