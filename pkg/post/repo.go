package post

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	i "my/redditclone/pkg/interface_mongoDB"
	"strconv"
	"time"
)

// WARNING! completly unsafe in multi-goroutine use, need add mutexes

type PostsMongoRepository struct {
	Posts i.PostInterface
}

var (
	ErrNoAuthor  = errors.New("you are not author")
	ErrNoComment = errors.New("no comment")
	ErrFind      = errors.New("error with find in databases")
	ErrCurDecode = errors.New("error with decode cursor")
	ErrInsert    = errors.New("error with insert to mongodb")
)

func NewMemoryRepo(collection i.PostInterface) *PostsMongoRepository {
	return &PostsMongoRepository{Posts: collection}
}

func (repo *PostsMongoRepository) GetAll() ([]*Post, error) {
	return repo.FindPostsInMongoDB(bson.M{})
}

func (repo *PostsMongoRepository) FindPostsInMongoDB(bson interface{}) (posts []*Post, err error) {
	ctx := context.Background()
	cur, err := repo.Posts.Find(ctx, bson)
	defer cur.Close(ctx)
	if err != nil {
		return nil, ErrFind
	}
	for cur.Next(ctx) {
		var post Post
		err = cur.Decode(&post)
		if err != nil {
			return nil, ErrCurDecode
		}
		posts = append(posts, &post)
		err = cur.Err()
		if err != nil {
			return nil, err
		}
	}
	return posts, nil
}

func (repo *PostsMongoRepository) AddPost(postJSON *Post, id, login string) error {
	posts, err := repo.FindPostsInMongoDB(bson.M{})
	if err != nil {
		return err
	}
	max := -1
	for _, val := range posts {
		var temp int
		temp, err = strconv.Atoi(val.ID)
		if err != nil {
			return err
		}
		if temp > max {
			max = temp
		}
	}
	postJSON.NewVote(1, id)
	postJSON.ID = strconv.Itoa(max + 1)
	postJSON.Created = time.Now().Format(time.RFC3339)
	postJSON.Views = 0
	postJSON.Author = Author{
		ID:       id,
		Username: login,
	}
	_, err = repo.Posts.InsertOne(context.TODO(), postJSON)
	if err != nil {
		return ErrInsert
	}
	return nil
}

func (repo *PostsMongoRepository) OpenPost(id string) (*Post, error) {
	posts, err := repo.FindPostsInMongoDB(bson.M{"id": id})
	if err != nil || len(posts) == 0 {
		return nil, err
	}
	return posts[0], err
}

func (repo *PostsMongoRepository) Category(name string) ([]*Post, error) {
	return repo.FindPostsInMongoDB(bson.M{"category": name})
}

func (repo *PostsMongoRepository) User(name string) ([]*Post, error) {
	return repo.FindPostsInMongoDB(bson.M{"author.username": name})
}

func (repo *PostsMongoRepository) Delete(postID string, authorID string) (err error) {
	posts, err := repo.FindPostsInMongoDB(bson.M{"id": postID})
	if err != nil {
		return err
	}
	if posts[0].Author.ID != authorID {
		return ErrNoAuthor
	}
	_, err = repo.Posts.DeleteOne(context.TODO(), bson.M{"id": postID})
	return err
}

func (repo *PostsMongoRepository) Vote(vote int, authorID string, postID string) (*Post, error) {

	post, err := repo.OpenPost(postID)
	if err != nil {
		return nil, err
	}
	post.NewVote(vote, authorID)

	_, err = repo.Posts.ReplaceOne(context.TODO(), bson.M{"id": postID}, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *PostsMongoRepository) UnVote(authorID string, postID string) (*Post, error) {
	post, err := repo.OpenPost(postID)
	if err != nil {
		return nil, err
	}
	post.DeleteVote(authorID)

	_, err = repo.Posts.ReplaceOne(context.TODO(), bson.M{"id": postID}, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *PostsMongoRepository) AddComment(postID, body, authorID, login string) (*Post, error) {
	post, err := repo.OpenPost(postID)
	if err != nil {
		return nil, err
	}
	post.NewComment(body, authorID, login)

	_, err = repo.Posts.ReplaceOne(context.TODO(), bson.M{"id": postID}, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *PostsMongoRepository) DeleteComment(postID, commentID, authorID string) (*Post, error) {
	post, err := repo.OpenPost(postID)
	if err != nil {
		return nil, err
	}
	err = post.DeleteComment(authorID, commentID)
	if err != nil {
		return nil, err
	}
	_, err = repo.Posts.ReplaceOne(context.TODO(), bson.M{"id": postID}, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
