package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"my/redditclone/pkg/post"
	"my/redditclone/pkg/session"
	"net/http"
)

type PostHandler struct {
	PostRepo PostsRepo
	Logger   *log.Logger
}

type PostsRepo interface {
	GetAll() ([]*post.Post, error)
	FindPostsInMongoDB(bson interface{}) (posts []*post.Post, err error)
	AddPost(postJSON *post.Post, id, login string) error
	OpenPost(id string) (*post.Post, error)
	Vote(vote int, authorID string, postID string) (*post.Post, error)
	UnVote(authorID string, postID string) (*post.Post, error)
	Delete(postID string, authorID string) (err error)
	Category(name string) ([]*post.Post, error)
	User(name string) ([]*post.Post, error)
	AddComment(postID, body, authorID, login string) (*post.Post, error)
	DeleteComment(postID, commentID, authorID string) (*post.Post, error)
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("successfully entry postHandler.List: [%s]  %s\n",
		r.Method, r.URL.Path)
	elems, err := h.PostRepo.GetAll()
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, err.Error())
		return
	}
	data, err := json.Marshal(elems)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "JSON Marshal error")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "Write error")
		return
	}
	h.Logger.Printf("successfully exit postHandler.List\n")
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("successfully entry postHandler.Create: [%s] %s\n",
		r.Method, r.URL.Path)
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "can not read body")
		return
	}
	postJSON := new(post.Post)
	err = json.Unmarshal(body, postJSON)
	if err != nil {
		SendMessage(w, http.StatusBadRequest, "incorrect params")
		return
	}
	id, login := session.SessionFromContext(r.Context())
	err = h.PostRepo.AddPost(postJSON, id, login)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, err.Error())
		return
	}
	data, err := json.Marshal(postJSON)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "JSON Marshal error")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "Write error")
		return
	}
	h.Logger.Printf("successfully exit postHandler.Create\n")
}

func (h *PostHandler) Open(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("successfully entry postHandler.Open: [%s] %s\n",
		r.Method, r.URL.Path)
	vars := mux.Vars(r)
	post, err := h.PostRepo.OpenPost(vars["id"])
	if err != nil {
		SendMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	data, err := json.Marshal(post)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "JSON Marshal error")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "Write error")
		return
	}
	h.Logger.Printf("successfully exit postHandler.Open\n")
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("successfully entry postHandler.Delete: [%s] %s\n",
		r.Method, r.URL.Path)
	vars := mux.Vars(r)
	id, _ := session.SessionFromContext(r.Context())
	err := h.PostRepo.Delete(vars["id"], id)
	if err != nil {
		SendMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	SendMessage(w, http.StatusOK, "success")
	h.Logger.Printf("successfully exit postHandler.Delete\n")
}

func (h *PostHandler) Category(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("successfully entry postHandler.Category: [%s] %s\n",
		r.Method, r.URL.Path)
	vars := mux.Vars(r)
	posts, err := h.PostRepo.Category(vars["category"])
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, err.Error())
		return
	}
	data, err := json.Marshal(posts)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "JSON Marshal error")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "Write error")
		return
	}
	h.Logger.Printf("successfully exit postHandler.Category\n")
}

func (h *PostHandler) User(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("successfully entry postHandler.User: [%s] %s\n",
		r.Method, r.URL.Path)
	vars := mux.Vars(r)
	posts, err := h.PostRepo.User(vars["user"])
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, err.Error())
		return
	}
	data, err := json.Marshal(posts)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "JSON Marshal error")
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "Write error")
		return
	}
	h.Logger.Printf("successfully exit postHandler.User\n")
}

func SendMessage(w http.ResponseWriter, status int, text string) {
	w.WriteHeader(status)
	resp, err1 := json.Marshal(map[string]interface{}{"message": text})
	if err1 != nil {
		fmt.Printf("JSON Marshal error")
		return
	}
	_, err := w.Write(resp)
	if err != nil {
		fmt.Printf("Write error")
	}
}
