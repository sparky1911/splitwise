package models

import (
	"errors"
	"log"
	"splitwise/db"
	"splitwise/utils"
	"time"
)

type User struct {
	ID        int64
	Username  string `binding:"required"`
	Email     string
	Password  string `binding:"required"`
	CreatedAt time.Time
}

func (u *User) Save() error {
	query := `INSERT INTO users(username,email,password_hash,created_at)
	VALUES($1,$2,$3,$4) RETURNING id`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	result, err := stmt.Exec(u.Username, u.Email, hashedPassword, u.CreatedAt)
	if err != nil {
		log.Println(err)
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err

}

func (u *User) ValidateUser() error {
	var hashedPassword string
	query := `SELECT password_hash FROM users WHERE username = ?`
	err := db.DB.QueryRow(query, u.Username).Scan(&hashedPassword)
	if err != nil {

		return errors.New("user not found")
	}
	if !utils.CheckPasswordHash(u.Password, hashedPassword) {
		return errors.New("invalid password")
	}

	return nil
}
