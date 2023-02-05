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
	"my/redditclone/pkg/session"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &CommentHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	com, err := json.Marshal(CommentInfo{
		Comment: "text",
	})
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	p.EXPECT().AddComment("1", "text", "1", "lake").Return(pt, nil)
	ptMarshal, err := json.Marshal(pt)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	req := httptest.NewRequest("POST", "/", bytes.NewReader(com))
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
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

	com, err = json.Marshal(CommentInfo{
		Comment: "text",
	})
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	p.EXPECT().AddComment("1", "text", "1", "lake").Return(nil, errors.New(""))
	req = httptest.NewRequest("POST", "/", bytes.NewReader(com))
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
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

	req = httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte{'0'}))
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
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

func TestDeleteComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &CommentHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	p.EXPECT().DeleteComment("1", "1", "1").Return(pt, nil)
	ptMarshal, err := json.Marshal(pt)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	req := httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"postID":    "1",
		"commentID": "1",
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
	if !bytes.Contains(bodyResp, ptMarshal) {
		t.Errorf("no text found")
		return
	}

	p.EXPECT().DeleteComment("1", "1", "1").Return(nil, errors.New(""))
	req = httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"postID":    "1",
		"commentID": "1",
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
