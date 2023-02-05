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

func TestUpvote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &VoteHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	p.EXPECT().Vote(1, "1", "1").Return(pt, nil)
	ptMarshal, err := json.Marshal(pt)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	req := httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w := httptest.NewRecorder()
	repo.Upvote(w, req)
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

	p.EXPECT().Vote(1, "1", "1").Return(nil, errors.New(""))
	req = httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w = httptest.NewRecorder()
	repo.Upvote(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}
}

func TestDownvote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &VoteHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	pt2 := pt
	pt2.Vote = pt2.Vote[0:]
	p.EXPECT().Vote(-1, "1", "1").Return(pt2, nil)

	req := httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w := httptest.NewRecorder()
	repo.Downvote(w, req)
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
	ptMarshal, err := json.Marshal(pt2)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	resp.Body.Close()
	if !bytes.Contains(bodyResp, ptMarshal) {
		t.Errorf("no text found")
		return
	}

	p.EXPECT().Vote(-1, "1", "1").Return(nil, errors.New(""))
	req = httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w = httptest.NewRecorder()
	repo.Downvote(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}
}

func TestUnvote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := NewMockPostsRepo(ctrl)
	repo := &VoteHandler{
		PostRepo: p,
		Logger:   log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile),
	}

	pt2 := pt
	pt2.Vote = pt2.Vote[0:]
	p.EXPECT().UnVote("1", "1").Return(pt2, nil)
	ptMarshal, err := json.Marshal(pt2)
	if err != nil {
		t.Errorf("Marshal error")
		return
	}
	req := httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w := httptest.NewRecorder()
	repo.Unvote(w, req)
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

	p.EXPECT().UnVote("1", "1").Return(nil, errors.New(""))
	req = httptest.NewRequest("POST", "/", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	req = req.WithContext(context.WithValue(req.Context(), session.SessionKey, map[string]interface{}{
		"id":       interface{}("1"),
		"username": interface{}("lake"),
	}))
	w = httptest.NewRecorder()
	repo.Unvote(w, req)
	resp = w.Result()
	if resp.StatusCode != 400 {
		t.Errorf("expected resp status 400, got %d", resp.StatusCode)
		return
	}
}
