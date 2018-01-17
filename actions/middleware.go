package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/mouse_hole/models"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/pkg/errors"
	"fmt"
)

const READ_REQUEST = "readRequest"

func ReadRequestAssigner(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		readReq := models.NewReadRequest()
		u := c.Data()["User"].(models.User)

		readReq.SetDB(u.DBName) // this is a test
		c.Set(READ_REQUEST, readReq)

		err := next(c)

		return err
	}
}

func authenticateRequest(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &MMClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
				log.Error(err)
				return nil, err
			}
			return []byte(SERVER_SECRET), nil
		})

		if claims, ok := token.Claims.(*MMClaims); ok && token.Valid {
			log.Infof("user: %v, appName: %v, expiration: %v", claims.UserName, claims.ApplicationName, claims.StandardClaims.ExpiresAt)
			user, err := models.GetDBInstance().GetUserById(tokenString)
			if err != nil {
				return err
			}
			c.Data()["User"] = user
			log.Infof("User model was atttached: %v", c.Data()["User"].(models.User))
		} else {
			return errors.New("Bad auth token!")
		}

		err = next(c)

		return err
	}
}
