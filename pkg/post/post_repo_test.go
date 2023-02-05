package post_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	i "my/redditclone/pkg/interface_mongoDB"
	"my/redditclone/pkg/mocks"
	"my/redditclone/pkg/post"
	"reflect"
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

func cursorIsOk(pt *post.Post) i.CursorInterface {
	var cur i.CursorInterface
	cur = &mocks.CursorInterface{}
	cur.(*mocks.CursorInterface).
		On("Next", context.Background()).
		Return(true).Once()
	cur.(*mocks.CursorInterface).
		On("Next", context.Background()).
		Return(false).Once()
	cur.(*mocks.CursorInterface).
		On("Close", context.Background()).
		Return(nil)
	cur.(*mocks.CursorInterface).
		On("Decode", &post.Post{}).
		Return(nil).Run(func(args mock.Arguments) {
		p := args.Get(0).(*post.Post)
		p.Score = pt.Score
		p.Views = pt.Views
		p.Type = pt.Type
		p.Category = pt.Category
		p.ID = pt.ID
		p.Author = pt.Author
		p.Comment = pt.Comment
		p.Vote = pt.Vote
	})
	cur.(*mocks.CursorInterface).
		On("Err").
		Return(nil)
	return cur

}

func TestGetAll(t *testing.T) {
	var col i.PostInterface
	var cur i.CursorInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	cur = &mocks.CursorInterface{}
	cur.(*mocks.CursorInterface).
		On("Next", context.Background()).
		Return(true).Once()
	cur.(*mocks.CursorInterface).
		On("Next", context.Background()).
		Return(false).Once()
	cur.(*mocks.CursorInterface).
		On("Close", context.Background()).
		Return(nil)
	cur.(*mocks.CursorInterface).
		On("Decode", &post.Post{}).
		Return(nil).Run(func(args mock.Arguments) {
		p := args.Get(0).(*post.Post)
		p.Score = pt.Score
		p.Views = pt.Views
		p.Type = pt.Type
		p.Category = pt.Category
		p.ID = pt.ID
		p.Author = pt.Author
		p.Comment = pt.Comment
		p.Vote = pt.Vote
	})
	cur.(*mocks.CursorInterface).
		On("Err").
		Return(errors.New(""))

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{}).
		Return(cur, nil).Once()
	res, err := databases.GetAll()
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{}).
		Return(cur, nil).Once()
	res, err = databases.GetAll()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	if !reflect.DeepEqual(res[0], pt) {
		t.Errorf("there were unfulfilled expectations: \n%#v\n%#v\n", res[0], pt)
	}

	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{}).
		Return(cur, errors.New("")).Once()
	_, err = databases.GetAll()
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}

func TestCategory(t *testing.T) {
	var col i.PostInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	var cur i.CursorInterface
	cur = &mocks.CursorInterface{}
	cur.(*mocks.CursorInterface).
		On("Close", context.Background()).
		Return(nil)
	cur.(*mocks.CursorInterface).
		On("Next", context.Background()).
		Return(true).Once()
	cur.(*mocks.CursorInterface).
		On("Decode", &post.Post{}).
		Return(errors.New(""))

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"category": "programming"}).
		Return(cur, nil).Once()
	res, err := databases.Category("programming")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}
	if res != nil {
		t.Errorf("Wrong result: %#v", res)
	}

	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"category": "programming"}).
		Return(cur, nil).Once()
	res, err = databases.Category("programming")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	if !reflect.DeepEqual(res[0], pt) {
		t.Errorf("there were unfulfilled expectations: \n%#v\n%#v\n", res[0], pt)
	}

	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}

func TestUser(t *testing.T) {
	var col i.PostInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	var cur i.CursorInterface
	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"author.username": "1"}).
		Return(cur, errors.New("")).Once()
	res, err := databases.User("1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}
	if res != nil {
		t.Errorf("Wrong result: %#v", res)
	}
	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}

func TestOpenPost(t *testing.T) {
	var col i.PostInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	var cur i.CursorInterface
	cur = cursorIsOk(pt)

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, errors.New("")).Once()
	res, err := databases.OpenPost("1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}
	if res != nil {
		t.Errorf("Wrong result: %#v", res)
	}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	res, err = databases.OpenPost("1")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}
	if !reflect.DeepEqual(res, pt) {
		t.Errorf("there were unfulfilled expectations: \n%#v\n%#v\n", res, pt)
	}
	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}

func TestDelete(t *testing.T) {
	var col i.PostInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	var cur i.CursorInterface
	cur = cursorIsOk(pt)
	var del i.DeleteResultInterface
	del = &mocks.DeleteResultInterface{}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("DeleteOne", context.Background(), bson.M{"id": "1"}).
		Return(del, nil).Once()
	err := databases.Delete("1", "1")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, errors.New("")).Once()
	err = databases.Delete("1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	cur = cursorIsOk(pt)

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	err = databases.Delete("1", "2")
	if err == nil {
		t.Errorf("unexpected err: %s", err)
	}
	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}

func TestVote(t *testing.T) {
	var col i.PostInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	var cur i.CursorInterface
	cur = cursorIsOk(pt)
	var up i.UpdateResultInterface
	up = &mocks.UpdateResultInterface{}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("ReplaceOne", context.Background(), bson.M{"id": "1"}, mock.Anything).
		Return(nil, errors.New("")).Once()
	_, err := databases.Vote(1, "1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("ReplaceOne", context.Background(), bson.M{"id": "1"}, mock.Anything).
		Return(up, nil).Once()
	_, err = databases.Vote(1, "1", "1")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, errors.New("")).Once()
	_, err = databases.Vote(1, "1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}
	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}

func TestUnVote(t *testing.T) {
	var col i.PostInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	var cur i.CursorInterface
	cur = cursorIsOk(pt)
	var up i.UpdateResultInterface
	up = &mocks.UpdateResultInterface{}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("ReplaceOne", context.Background(), bson.M{"id": "1"}, mock.Anything).
		Return(nil, errors.New("")).Once()
	_, err := databases.UnVote("1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	pt2 := &post.Post{
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
		}, {
			UserID: "2",
			Votes:  -1,
		}},
		Comment: []*post.Comment{{
			Author: post.Author{
				Username: "lake",
				ID:       "1",
			},
			Body: "hi",
		}},
		Score: 1,
	}

	cur = cursorIsOk(pt2)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("ReplaceOne", context.Background(), bson.M{"id": "1"}, mock.Anything).
		Return(up, nil).Once()
	_, err = databases.UnVote("1", "1")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, errors.New("")).Once()
	_, err = databases.UnVote("1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}
	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}

func TestDeleteComment(t *testing.T) {
	var col i.PostInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	var cur i.CursorInterface
	cur = cursorIsOk(pt)
	var up i.UpdateResultInterface
	up = &mocks.UpdateResultInterface{}
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("ReplaceOne", context.Background(), bson.M{"id": "1"}, mock.Anything).
		Return(nil, errors.New("")).Once()
	_, err := databases.DeleteComment("1", "1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("ReplaceOne", context.Background(), bson.M{"id": "1"}, mock.Anything).
		Return(up, nil).Once()
	_, err = databases.DeleteComment("1", "1", "1")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, errors.New("")).Once()
	_, err = databases.DeleteComment("1", "1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	_, err = databases.DeleteComment("1", "1", "2")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	pt2 := &post.Post{
		ID:       "1",
		Type:     "text",
		Category: "programming",
		Author: post.Author{
			Username: "lake",
			ID:       "1",
		},
		Score: 1,
	}
	cur = cursorIsOk(pt2)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	_, err = databases.DeleteComment("1", "1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}
	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}

func TestAdd(t *testing.T) {
	var col i.PostInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	var cur i.CursorInterface
	cur = cursorIsOk(pt)
	var in i.InsertOneResultInterface
	in = &mocks.InsertOneResultInterface{}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("InsertOne", context.Background(), mock.Anything).
		Return(in, nil).Once()
	err := databases.AddPost(pt, "1", "1")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{}).
		Return(cur, errors.New("")).Once()
	err = databases.AddPost(pt, "1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("InsertOne", context.Background(), mock.Anything).
		Return(in, errors.New("")).Once()
	err = databases.AddPost(pt, "1", "1")
	if err == nil {
		t.Errorf("unexpected err: %s", err)
	}

	pt2 := pt
	pt2.ID = "s"
	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{}).
		Return(cur, nil).Once()
	err = databases.AddPost(pt, "1", "1")
	if err == nil {
		t.Errorf("unexpected err: %s", err)
	}
	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}

func TestAddComment(t *testing.T) {
	var col i.PostInterface
	col = &mocks.PostInterface{}
	databases := post.NewMemoryRepo(col)

	var cur i.CursorInterface
	cur = cursorIsOk(pt)
	var up i.UpdateResultInterface
	up = &mocks.UpdateResultInterface{}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("ReplaceOne", context.Background(), bson.M{"id": "1"}, mock.Anything).
		Return(up, nil).Once()
	_, err := databases.AddComment("1", "lake", "1", "1")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
	}

	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, errors.New("")).Once()
	_, err = databases.AddComment("1", "lake", "1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}

	cur = cursorIsOk(pt)
	col.(*mocks.PostInterface).On("Find", context.Background(), bson.M{"id": "1"}).
		Return(cur, nil).Once()
	col.(*mocks.PostInterface).On("ReplaceOne", context.Background(), bson.M{"id": "1"}, mock.Anything).
		Return(up, errors.New("")).Once()
	_, err = databases.AddComment("1", "lake", "1", "1")
	if err == nil {
		t.Errorf("expected err: %s", err)
	}
	if !col.(*mocks.PostInterface).AssertExpectations(t) {
		t.Errorf("Assert Expectations isn't ok!")
	}
}
