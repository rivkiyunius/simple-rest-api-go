package models

import (
	"errors"
	"event-booking/db"
	"event-booking/utils"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) SignUp() error {
	query := `
		INSERT INTO users (email, password) VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Email, hashedPassword)
	return err
}

func (u *User) ValidateCredentials() error {
	query := `
		SELECT id, password FROM users WHERE email = ?
	`
	rows := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := rows.Scan(&u.Id, &retrievedPassword)
	if err != nil {
		return err
	}
	if !utils.CheckPasswordHash(u.Password, retrievedPassword) {
		return errors.New("Invalid credentials")
	}
	return nil
}
