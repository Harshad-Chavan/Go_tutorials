package db

import (
	"database/sql"
	"fmt"

	// special syntax we will not directly use it we will intereact with database/sql pacakge that is standard library
	// _ is required as we need it...ohterwise saving will remove it
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		fmt.Print(err)
		panic("Could not connect to database.")
	}

	// connection pooling
	DB.SetMaxOpenConns(10)

	// connection kept open when no one using open at all time (max - 10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createaEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER	
	)`

	_, err := DB.Exec(createaEventsTable)

	if err != nil {
		panic("Could not creted event table")
	}

}
