package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"log"
	"my/redditclone/pkg/session"
	"my/redditclone/pkg/user"
	"strconv"

	"net/http/httptest"
	"os"

	"testing"
	"text/template"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := NewMockUserRepo(ctrl)
	s := NewMockSessManager(ctrl)
	repo := &UserHandler{
		Tmpl:     template.Must(template.ParseGlob("../../static/html/index.html")),
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
		UserRepo: u,
		Sessions: s,
	}

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	repo.Index(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("expected resp status 200, got %d", resp.StatusCode)
		return
	}

	templateName = ""
	req = httptest.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()
	repo.Index(w, req)
	resp = w.Result()
	if resp.StatusCode == 200 {
		t.Errorf("expected resp status 200, got %d", resp.StatusCode)
		return
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := NewMockUserRepo(ctrl)
	s := NewMockSessManager(ctrl)
	repo := &UserHandler{
		Tmpl:     template.Must(template.ParseGlob("../../static/html/index.html")),
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
		UserRepo: u,
		Sessions: s,
	}

	password := "123456789"
	userInfo := &user.User{
		Login: "lake",
		ID:    1,
	}
	sessInfo := &session.Session{
		UserID: "1",
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiMSIsInVzZXJuYW1lIjoibGFrZSJ9fQ.MxDHms3U-y9AbkvYq_uweddyT33J8_Ph2xeoPcdaA3k",
	}

	s.EXPECT().Create(strconv.Itoa(int(userInfo.ID)), userInfo.Login).Return(sessInfo.Token, nil)
	u.EXPECT().Authorize(userInfo.Login, password).Return(userInfo, nil)

	bodyMap := map[string]string{"username": userInfo.Login, "password": password}
	bodyBytes, err := json.Marshal(bodyMap)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	body := bytes.NewBuffer(bodyBytes)
	req := httptest.NewRequest("GET", "/", body)
	w := httptest.NewRecorder()
	repo.Login(w, req)
	resp := w.Result()
	if resp.StatusCode != 201 {
		t.Errorf("expected resp status 201, got %d", resp.StatusCode)
		return
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read all error")
		return
	}
	resp.Body.Close()
	if !bytes.Contains(bodyResp, []byte(sessInfo.Token)) {
		t.Errorf("no text found")
		return
	}

	bodyBytes = []byte{'0'}
	body = bytes.NewBuffer(bodyBytes)
	req = httptest.NewRequest("GET", "/", body)
	w = httptest.NewRecorder()
	repo.Login(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}

	u.EXPECT().Authorize(userInfo.Login, password).Return(nil, errors.New(""))
	bodyMap = map[string]string{"username": userInfo.Login, "password": password}
	bodyBytes, err = json.Marshal(bodyMap)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	body = bytes.NewBuffer(bodyBytes)
	req = httptest.NewRequest("GET", "/", body)
	w = httptest.NewRecorder()
	repo.Login(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}

	s.EXPECT().Create(strconv.Itoa(int(userInfo.ID)), userInfo.Login).Return("", errors.New(""))
	u.EXPECT().Authorize(userInfo.Login, password).Return(userInfo, nil)
	bodyMap = map[string]string{"username": userInfo.Login, "password": password}
	bodyBytes, err = json.Marshal(bodyMap)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	body = bytes.NewBuffer(bodyBytes)
	req = httptest.NewRequest("GET", "/", body)
	w = httptest.NewRecorder()
	repo.Login(w, req)
	resp = w.Result()
	if resp.StatusCode != 500 {
		t.Errorf("expected resp status 500, got %d", resp.StatusCode)
		return
	}
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := NewMockUserRepo(ctrl)
	s := NewMockSessManager(ctrl)
	repo := &UserHandler{
		Tmpl:     template.Must(template.ParseGlob("../../static/html/index.html")),
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
		UserRepo: u,
		Sessions: s,
	}

	password := "123456789"
	userInfo := &user.User{
		Login: "lake",
		ID:    1,
	}
	sessInfo := &session.Session{
		UserID: "1",
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiMSIsInVzZXJuYW1lIjoibGFrZSJ9fQ.MxDHms3U-y9AbkvYq_uweddyT33J8_Ph2xeoPcdaA3k",
	}

	u.EXPECT().AddUser(userInfo.Login, password).Return(nil)
	s.EXPECT().Create(strconv.Itoa(int(userInfo.ID)), userInfo.Login).Return(sessInfo.Token, nil)
	u.EXPECT().Authorize(userInfo.Login, password).Return(userInfo, nil)
	bodyMap := map[string]string{"username": userInfo.Login, "password": password}
	bodyBytes, err := json.Marshal(bodyMap)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	body := bytes.NewBuffer(bodyBytes)
	req := httptest.NewRequest("GET", "/", body)
	w := httptest.NewRecorder()
	repo.Register(w, req)
	resp := w.Result()
	if resp.StatusCode != 201 {
		t.Errorf("expected resp status 201, got %d", resp.StatusCode)
		return
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read all error")
		return
	}
	resp.Body.Close()
	if !bytes.Contains(bodyResp, []byte(sessInfo.Token)) {
		t.Errorf("no text found")
		return
	}

	bodyBytes = []byte{'0'}
	body = bytes.NewBuffer(bodyBytes)
	req = httptest.NewRequest("GET", "/", body)
	w = httptest.NewRecorder()
	repo.Register(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}

	u.EXPECT().AddUser(userInfo.Login, password).Return(errors.New(""))
	bodyMap = map[string]string{"username": userInfo.Login, "password": password}
	bodyBytes, err = json.Marshal(bodyMap)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	body = bytes.NewBuffer(bodyBytes)
	req = httptest.NewRequest("GET", "/", body)
	w = httptest.NewRecorder()
	repo.Register(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}

	u.EXPECT().AddUser(userInfo.Login, password).Return(nil)
	u.EXPECT().Authorize(userInfo.Login, password).Return(nil, errors.New(""))
	bodyMap = map[string]string{"username": userInfo.Login, "password": password}
	bodyBytes, err = json.Marshal(bodyMap)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	body = bytes.NewBuffer(bodyBytes)
	req = httptest.NewRequest("GET", "/", body)
	w = httptest.NewRecorder()
	repo.Register(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}

	u.EXPECT().AddUser(userInfo.Login, password).Return(nil)
	u.EXPECT().Authorize(userInfo.Login, password).Return(userInfo, nil)
	s.EXPECT().Create(strconv.Itoa(int(userInfo.ID)), userInfo.Login).Return("", errors.New(""))
	bodyMap = map[string]string{"username": userInfo.Login, "password": password}
	bodyBytes, err = json.Marshal(bodyMap)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	body = bytes.NewBuffer(bodyBytes)
	req = httptest.NewRequest("GET", "/", body)
	w = httptest.NewRecorder()
	repo.Register(w, req)
	resp = w.Result()
	if resp.StatusCode != 500 {
		t.Errorf("expected resp status 500, got %d", resp.StatusCode)
		return
	}
}
