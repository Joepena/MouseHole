package models

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID           string    `bson:"_id"`
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
	return nil
}
func (u *User) Authenticate() bool {
	collection := GetDBInstance().session.DB("auth").C("users")

	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	u.PasswordHash = string(ph)

	err = collection.Find(bson.M{"email": u.Email, "password_hash": u.PasswordHash}).One(&u)
	if err != nil {
		return false
	}

	return true
}
