package models

import (
	"fmt"

	"example.com/mod/db"
)

type User struct {
	ID       int64
	Email    string `binding:required`
	Password string `binding:required`
}

func (user User) Save() error {

	query := `INSERT INTO users(email,password) VALUES (?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)

	if err != nil {
		fmt.Println(err)
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId

	return nil

}
