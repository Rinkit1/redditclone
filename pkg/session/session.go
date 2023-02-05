package session

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Session struct {
	Token   string
	UserID  string
	Expires time.Time
}

func NewSession(id string, login string) (*Session, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]interface{}{
			"username": login,
			"id":       id,
		},
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &Session{
		Token:   tokenString,
		UserID:  id,
		Expires: time.Now().Add(time.Hour),
	}, nil
}
