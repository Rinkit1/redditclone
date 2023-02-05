package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"my/redditclone/pkg/post"
	"my/redditclone/pkg/session"
	"net/http/httptest"
	"os"
	"testing"
)

var pt = &post.Post{
	ID:       "1",
	Type:     "text",
	Category: "programming",
	Author: post.Author{
		Username: "lake",
		ID:       "1",
	},
	Vote: []*post.Vote{{
		UserID: "1",
		Votes:  1,
	}},
	Comment: []*post.Comment{{
		Author: post.Author{
			Username: "lake",
			ID:       "1",
		},
		Body: "hi",
		ID:   "1",
	}},
	Score: 1,
}

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &PostHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	p.EXPECT().GetAll().Return([]*post.Post{pt}, nil)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	repo.List(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("expected resp status 200, got %d", resp.StatusCode)
		return
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read all error")
		return
	}
	resp.Body.Close()
	ptMarshal, err := json.Marshal(pt)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	if !bytes.Contains(bodyResp, ptMarshal) {
		t.Errorf("no text found")
		return
	}

	p.EXPECT().GetAll().Return(nil, errors.New(""))
	req = httptest.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()
	repo.List(w, req)
	resp = w.Result()
	if resp.StatusCode != 500 {
		t.Errorf("expected resp status 500, got %d", resp.StatusCode)
		return
	}

}

func TestCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &PostHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	p.EXPECT().Category("programming").Return([]*post.Post{pt}, nil)
	req := httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"category": "programming",
	})
	w := httptest.NewRecorder()
	repo.Category(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("expected resp status 200, got %d", resp.StatusCode)
		return
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read all error")
		return
	}
	resp.Body.Close()
	ptMarshal, err := json.Marshal(pt)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	if !bytes.Contains(bodyResp, ptMarshal) {
		t.Errorf("no text found")
		return
	}

	p.EXPECT().Category("programming").Return(nil, errors.New(""))
	req = httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"category": "programming",
	})
	w = httptest.NewRecorder()
	repo.Category(w, req)
	resp = w.Result()
	if resp.StatusCode != 500 {
		t.Errorf("expected resp status 500, got %d", resp.StatusCode)
		return
	}
}

func TestUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &PostHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	p.EXPECT().User("lake").Return([]*post.Post{pt}, nil)
	req := httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"user": "lake",
	})
	w := httptest.NewRecorder()
	repo.User(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("expected resp status 200, got %d", resp.StatusCode)
		return
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read all error")
		return
	}
	resp.Body.Close()
	ptMarshal, err := json.Marshal(pt)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	if !bytes.Contains(bodyResp, ptMarshal) {
		t.Errorf("no text found")
		return
	}

	p.EXPECT().User("lake").Return(nil, errors.New(""))
	req = httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"user": "lake",
	})
	w = httptest.NewRecorder()
	repo.User(w, req)
	resp = w.Result()
	if resp.StatusCode != 500 {
		t.Errorf("expected resp status 500, got %d", resp.StatusCode)
		return
	}
}

func TestOpen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &PostHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	p.EXPECT().OpenPost("1").Return(pt, nil)
	req := httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	w := httptest.NewRecorder()
	repo.Open(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("expected resp status 200, got %d", resp.StatusCode)
		return
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read all error")
		return
	}
	resp.Body.Close()
	ptMarshal, err := json.Marshal(pt)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	if !bytes.Contains(bodyResp, ptMarshal) {
		t.Errorf("no text found")
		return
	}

	p.EXPECT().OpenPost("1").Return(nil, errors.New(""))
	req = httptest.NewRequest("GET", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	w = httptest.NewRecorder()
	repo.Open(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 500, got %d", resp.StatusCode)
		return
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &PostHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	p.EXPECT().Delete("1", "1").Return(nil)
	req := httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w := httptest.NewRecorder()
	repo.Delete(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("expected resp status 200, got %d", resp.StatusCode)
		return
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read all error")
		return
	}
	resp.Body.Close()
	if !bytes.Contains(bodyResp, []byte(`{"message":"success"}`)) {
		t.Errorf("no text found")
		return
	}

	p.EXPECT().Delete("1", "1").Return(errors.New(""))
	req = httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w = httptest.NewRecorder()
	repo.Delete(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &PostHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	p.EXPECT().AddPost(pt, "1", "lake").Return(nil)
	ptMarshal, err := json.Marshal(pt)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	req := httptest.NewRequest("POST", "/", bytes.NewReader(ptMarshal))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w := httptest.NewRecorder()
	repo.Create(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("expected resp status 200, got %d", resp.StatusCode)
		return
	}
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read all error")
		return
	}
	resp.Body.Close()
	if !bytes.Contains(bodyResp, ptMarshal) {
		t.Errorf("no text found")
		return
	}

	ptMarshal, err = json.Marshal(pt)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	p.EXPECT().AddPost(pt, "1", "lake").Return(errors.New(""))
	req = httptest.NewRequest("POST", "/", bytes.NewReader(ptMarshal))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w = httptest.NewRecorder()
	repo.Create(w, req)
	resp = w.Result()
	if resp.StatusCode != 500 {
		t.Errorf("expected resp status 500, got %d", resp.StatusCode)
		return
	}

	req = httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte{'0'}))
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w = httptest.NewRecorder()
	repo.Create(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}
}
