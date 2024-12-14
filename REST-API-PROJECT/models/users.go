package models

import (
	"errors"
	"fmt"

	"example.com/mod/db"
	"example.com/mod/utils"
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

	//encrypting the password using hash
	hashedpass, err := utils.HashPassword(user.Password)

	if err != nil {
		fmt.Println(err)
		return err
	}

	result, err := stmt.Exec(user.Email, hashedpass)

	if err != nil {
		fmt.Println(err)
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId

	return nil

}

func (user User) ValidateCredentials() error {

	selectquery := `SELECT Id,password FROM USERS Where email = ?`

	row := db.DB.QueryRow(selectquery, user.Email)

	var passwordFromDb string
	err := row.Scan(&user.ID, &passwordFromDb)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// hash password and check
	passwordIsValid := utils.CheckPassword(user.Password, passwordFromDb)

	if !passwordIsValid {
		return errors.New("Credentials Invalid")

	}

	return nil

}
