package db

import (
	"database/sql"
	"fmt"

	// special syntax we will not directly use it we will intereact with database/sql pacakge that is standard library
	// _ is required as we need it...ohterwise saving will remove it
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "harshad"
)

var DB *sql.DB

func InitDB(db_name string) {
	var err error

	if db_name == "sqlite3" {
		DB, err = sql.Open("sqlite3", "api.db")
	}

	if db_name == "postgres" {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		DB, err = sql.Open("postgres", psqlInfo)
	}

	if err != nil {
		fmt.Print(err)
		panic("Could not connect to database.")

	}
	// defer DB.Close()

	// err = DB.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// connection pooling
	DB.SetMaxOpenConns(10)

	// connection kept open when no one using open at all time (max - 10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	)`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not creted Users table")
	}

	createaEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
		FOREIGN KEY (user_id) REFERENCES users(id)	
	)`

	_, err = DB.Exec(createaEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not creted event table")
	}

}
