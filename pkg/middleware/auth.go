package middleware

import (
	"context"
	"my/redditclone/pkg/handlers"
	"my/redditclone/pkg/session"
	"net/http"
	"strings"
)

func Auth(sm *session.SessionsManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		temp := r.Header.Get("authorization")
		if temp == "" {
			next.ServeHTTP(w, r)
			return
		}
		token := strings.Split(temp, " ")
		user, err := sm.Check(token[1])
		if err != nil {
			handlers.SendMessage(w, http.StatusInternalServerError, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), session.SessionKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
