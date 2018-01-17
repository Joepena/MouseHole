package actions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	log "github.com/sirupsen/logrus"
	"github.com/joepena/mouse_hole/models"
	"github.com/pkg/errors"
)


type MMClaims struct {
	UserName string `json:"user_name"`
	ApplicationName string `json:"application_name"`
	jwt.StandardClaims
}

func getAuthToken(c buffalo.Context) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["user_name"] = c.Request().Header.Get("user_name")
	claims["application_name"] = c.Request().Header.Get("application_name")

	tokenString, _ := token.SignedString(SERVER_SECRET)

	// logic to attach token to user obj
	log.Infof("token generated %v", tokenString)

	return tokenString, nil
}

func createUserHandler(c buffalo.Context) error {
	token, err := getAuthToken(c)
	if err != nil {
		log.Error(err)
	}
	h := c.Request().Header
	u := models.User{
		ID:           token,
		Email:        h.Get("email"),
		Password:     h.Get("password"),
		PasswordHash: "",
		DBName:       h.Get("application_name"),
	}

	err = u.Create(models.GetDBInstance())
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func loginHandler(c buffalo.Context) error {
	email := c.Value("email").(string)
	password := c.Value("password").(string)

	u := models.User{
		Email:        email,
		Password:     password,
	}

	validUser := u.Authenticate()
	if !validUser {
		err := errors.New("invalid credentials")
		log.Error(err)
		c.Flash().Add("error", err.Error())
		return err
	} else {
		// create valid session
		s := c.Session()
		s.Set("email", u.Email)
		s.Set("DbName", u.DBName)
		err := s.Save()
		if err != nil {
			return err
		}
	}

	return nil
}

