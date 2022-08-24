package main

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Name 		  string `json:"name" db:"name"`
}

// get user from id
func (u *User) get(db *sqlx.DB) error {
	return db.Get(u, "SELECT name, email FROM users WHERE id=?", u.ID)
}

// Update user
funct (u *User) update(db *sqlx.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = db.Exect("UPDATE users SET name=?, email=?, password=? WHERE id=?", u.Name, u.Email, string(hashedPassword), u.ID)

	return err // return error because it is a void method, so if is successfully then error will be null
}

// Delete User
func (u *User) delete(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=?", u.ID)

	return err
}

// Create User
func (u *User) create(db *sqlx.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?) RETURNING id", u.Name, u.Email, string(hashedPassword))

	return err
}

// List users
func list(db *sqlx.DB) ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT id, name, email FROM users")

	if err != nil {
		return nil, err
	}

	return users, err
}