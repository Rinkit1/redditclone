package user

import (
	"database/sql"
	"errors"
)

var (
	ErrNoUser           = errors.New("no user found")
	ErrBadPass          = errors.New("invalid password")
	ErrUserAlreadyExist = errors.New("user already exist")
)

type UserMysqlRepository struct {
	DB *sql.DB
}

func NewMemoryRepo(db *sql.DB) *UserMysqlRepository {
	return &UserMysqlRepository{DB: db}
}
func (repo *UserMysqlRepository) Authorize(login, pass string) (*User, error) {
	user := &User{}
	err := repo.DB.
		QueryRow("SELECT id, login, password FROM users WHERE login = ?", login).
		Scan(&user.ID, &user.Login, &user.password)
	if err != nil {
		return nil, ErrNoUser
	}
	if user.password != pass {
		return nil, ErrBadPass
	}

	return user, nil
}
func (repo *UserMysqlRepository) AddUser(login, pass string) error {
	user := &User{}
	err := repo.DB.
		QueryRow("SELECT id, login, password FROM users WHERE login = ?", login).
		Scan(&user.ID, &user.Login, &user.password)
	if err == nil {
		return ErrUserAlreadyExist
	}

	_, err = repo.DB.Exec(
		"INSERT INTO users (`login`, `password`) VALUES (?, ?)",
		login,
		pass,
	)
	if err != nil {
		return err
	}
	return nil
}
