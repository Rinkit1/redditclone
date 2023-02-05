package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"my/redditclone/pkg/user"
	"net/http"
	"strconv"
	"text/template"
)

type UserRepo interface {
	Authorize(login, pass string) (*user.User, error)
	AddUser(login, pass string) error
}

type SessManager interface {
	Create(id, login string) (tokenString string, err error)
	Check(inToken string) (userInfo map[string]interface{}, err error)
}

type UserHandler struct {
	Tmpl     *template.Template
	Logger   *log.Logger
	UserRepo UserRepo
	Sessions SessManager
}
type UserJSON struct {
	Username string
	Password string
}

var templateName = "index.html"

func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("successfully entry userHandler.Index: [%s] %s\n",
		r.Method, r.URL.Path)
	err := h.Tmpl.ExecuteTemplate(w, templateName, nil)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "template error")
		return
	}
	h.Logger.Printf("successfully exit userHandler.Index\n")
}
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("successfully entry userHandler.Login: [%s] %s\n",
		r.Method, r.URL.Path)
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "can not read body")
		return
	}
	userJSON := new(UserJSON)
	err = json.Unmarshal(body, &userJSON)
	if err != nil {
		SendMessage(w, http.StatusBadRequest, "incorrect params")
		return
	}
	u, err := h.UserRepo.Authorize(userJSON.Username, userJSON.Password)
	if err != nil {
		SendMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.Sessions.Create(strconv.Itoa(int(u.ID)), u.Login)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	data, err := json.Marshal(map[string]interface{}{"token": token})
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "JSON Marshal error")
		return
	}
	_, err = w.Write(data)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "Write error")
		return
	}
	h.Logger.Printf("successfully exit userHandler.Login\n")
}
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	h.Logger.Printf("successfully entry userHandler.Register: [%s] %s\n",
		r.Method, r.URL.Path)
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "can not read body")
		return
	}
	userJSON := new(UserJSON)
	err = json.Unmarshal(body, &userJSON)
	if err != nil {
		SendMessage(w, http.StatusBadRequest, "incorrect params")
		return
	}
	err = h.UserRepo.AddUser(userJSON.Username, userJSON.Password)
	if err != nil {
		SendMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	u, err := h.UserRepo.Authorize(userJSON.Username, userJSON.Password)
	if err != nil {
		SendMessage(w, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.Sessions.Create(strconv.Itoa(int(u.ID)), u.Login)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	data, err := json.Marshal(map[string]interface{}{"token": token})
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "JSON Marshal error")
		return
	}
	_, err = w.Write(data)
	if err != nil {
		SendMessage(w, http.StatusInternalServerError, "Write error")
		return
	}
	h.Logger.Printf("successfully exit userHandler.Register\n")
}
