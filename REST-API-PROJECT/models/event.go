package models

import (
	"time"

	"example.com/mod/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	// store it in database
	// events = append(events, e)

	insertquery := `INSERT INTO EVENTS(name, description, location, dateTime, user_id )
					Values (?,?,?,?,?)`

	// prepares the statment to be executed
	stmt, err := db.DB.Prepare(insertquery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// executes the statement
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserID)
	if err != nil {
		return err
	}

	// returns the last inseeted id
	id, err := result.LastInsertId()

	e.ID = id

	return err

}

func GetAllEvents() []Event {
	return events
}
