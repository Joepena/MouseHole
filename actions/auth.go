package actions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	log "github.com/sirupsen/logrus"
)

// TODO: remove this from source later
var SERVER_SECRET = []byte("057E3CE6B941756FD9CAB17D93C522F7C3745A78066A278E83999FFF547C5A8F")

type MMClaims struct {
	UserName string `json:"user_name"`
	ApplicationName string `json:"application_name"`
	jwt.StandardClaims
}

func getTokenHandler(c buffalo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["user_name"] = c.Request().Header.Get("user_name")
	claims["application_name"] = c.Request().Header.Get("application_name")

	tokenString, _ := token.SignedString(SERVER_SECRET)

	// logic to attach token to user obj
	log.Infof("token generated %v", tokenString)

	return nil
}
