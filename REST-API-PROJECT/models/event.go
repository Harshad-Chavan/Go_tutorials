package models

import (
	"fmt"
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

	// defer db.Close()

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

func GetAllEvents() ([]Event, error) {

	var events []Event

	selectAllquery := "SELECT * FROM EVENTS"
	rows, err := db.DB.Query(selectAllquery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)

	}

	return events, nil
}

// returning a pointer to event as
// we cannot send nil for a struct
// either any empty struct or filled struct
func GetEventById(id int64) (*Event, error) {

	var event Event

	selectById := "Select * from EVENTS where ID = ?"

	//will return only one row..as we expect
	row := db.DB.QueryRow(selectById, id)

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil

}
