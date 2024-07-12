package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatalf("could not connect to DB: %v", err)
	}
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(20)
	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("could not create users table: %v", err)
	}

	createGroupTable := `CREATE TABLE IF NOT EXISTS groups(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = DB.Exec(createGroupTable)
	if err != nil {
		log.Fatalf("could not create groups table %v", err)
	}
}