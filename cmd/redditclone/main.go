package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"my/redditclone/pkg/handlers"
	i "my/redditclone/pkg/interface_mongoDB"
	"my/redditclone/pkg/middleware"
	"my/redditclone/pkg/post"
	"my/redditclone/pkg/session"
	"my/redditclone/pkg/user"
	"net/http"
	"os"
	"text/template"
	"time"
)

func main() {
	templates := template.Must(template.ParseGlob("./redditclone/static/html/index.html"))
	logger := log.New(os.Stdout, "STD ", log.Ltime|log.Lshortfile)

	dsn := "root:love@tcp(localhost:3306)/golang?"
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Printf("Can't open mysql database. %s", err.Error())
		return
	}
	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		logger.Printf("Can't ping mysql database. %s", err.Error())
		return
	}
	sessRepo := session.NewSessionsManager(db)
	userRepo := user.NewMemoryRepo(db)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Printf("Can't connect to mongodb database. %s", err.Error())
		return
	}
	ctx1, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	defer func() {
		err = client.Disconnect(ctx1)
		if err != nil {
			logger.Printf("Can't disconnect to mongodb database. %s", err.Error())
		}
	}()
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx2, nil)
	if err != nil {
		logger.Printf("Can't ping mongodb database. %s", err.Error())
		return
	}
	postsCollection := client.Database("posts").Collection("posts")

	postRepo := post.NewMemoryRepo(i.NewUserDatabase(postsCollection))

	postHandlers := &handlers.PostHandler{
		Logger:   logger,
		PostRepo: postRepo,
	}
	userHandlers := &handlers.UserHandler{
		Tmpl:     templates,
		Logger:   logger,
		UserRepo: userRepo,
		Sessions: sessRepo,
	}
	voteHandlers := &handlers.VoteHandler{
		Logger:   logger,
		PostRepo: postRepo,
	}
	commentHandlers := &handlers.CommentHandler{
		Logger:   logger,
		PostRepo: postRepo,
	}

	r := mux.NewRouter()

	r.Handle("/static/js/{file}", http.FileServer(http.Dir("./redditclone")))
	r.Handle("/static/css/{file}", http.FileServer(http.Dir("./redditclone")))

	r.NotFoundHandler = http.HandlerFunc(userHandlers.Index)
	r.HandleFunc("/api/login", userHandlers.Login).Methods("POST")
	r.HandleFunc("/api/register", userHandlers.Register).Methods("POST")

	r.HandleFunc("/api/posts/", postHandlers.List).Methods("GET")
	r.HandleFunc("/api/posts", postHandlers.Create).Methods("POST")
	r.HandleFunc("/api/post/{id}", postHandlers.Open).Methods("GET")
	r.HandleFunc("/api/post/{id}", postHandlers.Delete).Methods("DELETE")
	r.HandleFunc("/api/posts/{category}", postHandlers.Category).Methods("GET")
	r.HandleFunc("/api/user/{user}", postHandlers.User).Methods("GET")

	r.HandleFunc("/api/post/{id}/upvote", voteHandlers.Upvote).Methods("GET")
	r.HandleFunc("/api/post/{id}/downvote", voteHandlers.Downvote).Methods("GET")
	r.HandleFunc("/api/post/{id}/unvote", voteHandlers.Unvote).Methods("GET")

	r.HandleFunc("/api/post/{id}", commentHandlers.Create).Methods("POST")
	r.HandleFunc("/api/post/{postID}/{commentID}", commentHandlers.Delete).Methods("DELETE")

	mux := middleware.Auth(sessRepo, r)
	mux = middleware.AccessLog(logger, mux)
	mux = middleware.Panic(mux)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Println("listen and serve error")
	}
}
