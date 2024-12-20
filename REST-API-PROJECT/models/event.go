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
	UserID      int64
}

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

func (event Event) Update() error {

	updateeventquery := `UPDATE EVENTS 
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE ID = ?	
	`
	stmt, err := db.DB.Prepare(updateeventquery)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Datetime, event.ID)

	if err != nil {
		return err
	}

	return err

}

func (event Event) DeleteEvent() error {

	deleteById := "Delete from EVENTS where ID = ?"

	stmt, err := db.DB.Prepare(deleteById)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e Event) Register(Userid int64) error {
	// store it in database
	// events = append(events, e)

	insertquery := `INSERT INTO Registrations(event_id, user_id )
					Values (?,?)`

	// prepares the statment to be executed
	stmt, err := db.DB.Prepare(insertquery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// executes the statement
	_, err = stmt.Exec(e.ID, Userid)
	if err != nil {
		return err
	}

	return err

}

func (e Event) CancelRegisteration() error {

	query := `DELETE from Registrations where event_id = ? and user_id = ?`

	// prepares the statment to be executed
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// executes the statement
	_, err = stmt.Exec(e.ID, e.UserID)
	if err != nil {
		return err
	}

	return err

}
