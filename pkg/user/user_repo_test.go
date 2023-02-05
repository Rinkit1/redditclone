package user

import (
	"fmt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"reflect"
	"testing"
)

func TestAuthorize(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	id := 0
	login := "lake"
	password := "123456789"
	wrongPass := "12345678"
	// all is ok
	rows := sqlmock.
		NewRows([]string{"id", "login", "password"})
	expect := []*User{
		{int64(id), login, password},
	}
	for _, user := range expect {
		rows = rows.AddRow(user.ID, user.Login, user.password)
	}
	mock.
		ExpectQuery("SELECT id, login, password FROM users WHERE").
		WithArgs(login).
		WillReturnRows(rows)
	repo := NewMemoryRepo(db)
	user, err := repo.Authorize(login, password)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err1 := mock.ExpectationsWereMet(); err1 != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], user)
		return
	}
	// no user
	mock.
		ExpectQuery("SELECT id, login, password FROM users WHERE").
		WithArgs(login).
		WillReturnError(fmt.Errorf("sql: no rows in result set"))
	_, err = repo.Authorize(login, password)
	if err1 := mock.ExpectationsWereMet(); err1 != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	// bad pass
	rows = sqlmock.NewRows([]string{"id", "login", "password"}).
		AddRow(id, login, password)
	mock.
		ExpectQuery("SELECT id, login, password FROM users WHERE").
		WithArgs(login).
		WillReturnRows(rows)
	_, err = repo.Authorize(login, wrongPass)
	if err1 := mock.ExpectationsWereMet(); err1 != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	id := 0
	login := "lake"
	password := "123456789"
	// all is ok
	mock.
		ExpectQuery("SELECT id, login, password FROM users WHERE").
		WithArgs(login).
		WillReturnError(fmt.Errorf("sql: no rows in result set"))
	mock.
		ExpectExec("INSERT INTO users \\(`login`, `password`\\) VALUES \\(\\?, \\?\\)").
		WithArgs(login, password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewMemoryRepo(db)
	err = repo.AddUser(login, password)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err1 := mock.ExpectationsWereMet(); err1 != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	// user already exist
	rows := sqlmock.
		NewRows([]string{"id", "login", "password"})
	expect := []*User{
		{int64(id), login, password},
	}
	for _, user := range expect {
		rows = rows.AddRow(user.ID, user.Login, user.password)
	}
	mock.
		ExpectQuery("SELECT id, login, password FROM users WHERE").
		WithArgs(login).
		WillReturnRows(rows)
	err = repo.AddUser(login, password)
	if err1 := mock.ExpectationsWereMet(); err1 != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	// user already exist
	mock.
		ExpectQuery("SELECT id, login, password FROM users WHERE").
		WithArgs(login).
		WillReturnError(fmt.Errorf("sql: no rows in result set"))
	mock.
		ExpectExec("INSERT INTO users \\(`login`, `password`\\) VALUES \\(\\?, \\?\\)").
		WithArgs(login, password).
		WillReturnError(fmt.Errorf("sql: db_error"))
	err = repo.AddUser(login, password)
	if err1 := mock.ExpectationsWereMet(); err1 != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
