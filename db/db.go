package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DBCon *sql.DB
)
