package session

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = "HalloweenIsComingSoonAndIAmGladAboutIt"
var SessionKey = "user"
var (
	ErrBadSignMethod   = errors.New("bad sign method")
	ErrBadToken        = errors.New("bad token")
	ErrNoPayload       = errors.New("no payload")
	ErrNoSession       = errors.New("no session")
	ErrIncorrectUserID = errors.New("incorrect userID")
	ErrSessionIsDead   = errors.New("session is dead")
)

type SessionsManager struct {
	DB *sql.DB
}

func NewSessionsManager(db *sql.DB) *SessionsManager {
	return &SessionsManager{DB: db}
}

func (sm *SessionsManager) Create(id, login string) (tokenString string, err error) {
	sess, err := NewSession(id, login)
	if err != nil {
		return "", err
	}
	if id != sess.UserID {
		return "", ErrIncorrectUserID
	}
	_, err = sm.DB.Exec(
		"INSERT INTO sessions (`token`, `user`, `expires`) VALUES (?, ?,?)",
		sess.Token,
		sess.UserID,
		sess.Expires.Unix(),
	)
	if err != nil {
		return "", err
	}
	return sess.Token, nil
}

func (sm *SessionsManager) Check(inToken string) (userInfo map[string]interface{}, err error) {
	hashSecretGetter := func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != "HS256" {
			return nil, ErrBadSignMethod
		}
		return []byte(secretKey), nil
	}

	token, err := jwt.Parse(inToken, hashSecretGetter)
	if err != nil || !token.Valid {
		return nil, ErrBadToken
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrNoPayload
	}
	userInfo = payload["user"].(map[string]interface{})

	id := userInfo["id"].(string)

	var sessExpires int64
	sess := &Session{}
	err = sm.DB.
		QueryRow("SELECT user,expires FROM sessions WHERE token = ?", inToken).
		Scan(&sess.UserID, &sessExpires)
	if err != nil {
		return nil, ErrNoSession
	}
	sess.Expires = time.Unix(sessExpires, 0)

	if id != sess.UserID {
		return nil, ErrIncorrectUserID
	} else if time.Since(sess.Expires) > 0 {
		_, err = sm.DB.Exec(
			"DELETE FROM sessions WHERE user = ?",
			id,
		)
		if err != nil {
			return nil, err
		}
		return nil, ErrSessionIsDead
	}
	return userInfo, nil
}

func SessionFromContext(ctx context.Context) (string, string) {
	user := ctx.Value(SessionKey).(map[string]interface{})
	id := user["id"].(string)
	login := user["username"].(string)
	return id, login
}
