package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"git.aprentout.com/touslesmemes/api/models"
	"github.com/badoux/checkmail"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest represents a login form.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UsersLogin perform a login with the given credentials.
func UsersLogin(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	var req LoginRequest
	err := c.Bind(&req)

	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}

	pwd := req.Password
	if len(pwd) == 0 {
		return c.Error(http.StatusBadRequest, errors.New("Invalid password"))
	}

	email := req.Email
	if checkmail.ValidateFormat(email) != nil {
		return c.Error(http.StatusBadRequest, errors.New("Invalid email"))
	}

	u := &models.User{Email: email}

	if err := tx.Where("email = ?", u.Email).First(u); err != nil {
		return c.Error(404, err)
	}

	if err != nil {
		return c.Error(http.StatusBadRequest, errors.New("Login failed"))
	}

	pwdCompare := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if pwdCompare != nil {
		return c.Error(http.StatusBadRequest, errors.New("Login failed"))
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(oneWeek()).Unix(),
		Issuer:    fmt.Sprintf("%s.%s", envy.Get("GO_ENV", "development"), u.Pseudo),
		Id:        u.ID.String(),
	}

	signingKey, err := ioutil.ReadFile(envy.Get("JWT_KEY_PATH", ""))

	if err != nil {
		return fmt.Errorf("could not open jwt key, %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return fmt.Errorf("could not sign token, %v", err)
	}

	return c.Render(200, r.JSON(map[string]string{"token": tokenString}))
}

func oneWeek() time.Duration {
	return 7 * 24 * time.Hour
}

// RestrictedHandlerMiddleware searches and parses the jwt token in order to authenticate
// the request and populate the Context with the user contained in the claims.
func RestrictedHandlerMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		tx, ok := c.Value("tx").(*pop.Connection)
		if !ok {
			return errors.WithStack(errors.New("no transaction found"))
		}

		tokenString := c.Request().Header.Get("Authorization")

		if len(tokenString) == 0 {
			return c.Error(http.StatusUnauthorized, fmt.Errorf("No token set in headers"))
		}

		// parsing token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// key
			mySignedKey, err := ioutil.ReadFile(envy.Get("JWT_KEY_PATH", ""))

			if err != nil {
				return nil, fmt.Errorf("could not open jwt key, %v", err)
			}

			return mySignedKey, nil
		})

		if err != nil {
			return c.Error(http.StatusUnauthorized, fmt.Errorf("Could not parse the token, %v", err))
		}

		// getting claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			logrus.Errorf("claims: %v", claims)

			// retrieving user from db
			u := &models.User{ID: uuid.FromStringOrNil(claims["jti"].(string))}

			// To find the User the parameter user_id is used.
			if err := tx.Where("id = ?", u.ID).First(u); err != nil {
				return c.Error(404, err)
			}

			if err != nil {
				return c.Error(http.StatusUnauthorized, fmt.Errorf("Could not identify the user"))
			}

			c.Set("user", u)

		} else {
			return c.Error(http.StatusUnauthorized, fmt.Errorf("Failed to validate token: %v", claims))
		}

		return next(c)
	}
}
