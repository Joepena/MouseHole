package models

import (
	"strings"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID `bson:"_id"`
	Email        string    `bson:"email"`
	Password     string    `bson:"-"`
	PasswordHash string    `bson:"password_hash"`
	DBName       string    `bson:"db_name"`
}

// Create wraps up the pattern of encrypting the password and
// running validations.
func (u *User) Create(db *DB) error {
	u.Email = strings.ToLower(u.Email)
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(ph)
	return db.session.DB("auth").C("users").Insert(u)
}
