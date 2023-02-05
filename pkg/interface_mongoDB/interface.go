package dop

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type postDatabase struct {
	db *mongo.Collection
}
type postCursor struct {
	cursor *mongo.Cursor
}
type postInsertOneResult struct {
	result *mongo.InsertOneResult
}
type postDeleteResult struct {
	result *mongo.DeleteResult
}
type postReplaceResult struct {
	result *mongo.UpdateResult
}

type PostInterface interface {
	Find(context.Context, interface{}, ...*options.FindOptions) (CursorInterface, error)
	InsertOne(context.Context, interface{}, ...*options.InsertOneOptions) (InsertOneResultInterface, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (DeleteResultInterface, error)
	ReplaceOne(context.Context, interface{}, interface{}, ...*options.ReplaceOptions) (UpdateResultInterface, error)
}

func NewUserDatabase(db *mongo.Collection) PostInterface {
	return &postDatabase{db: db}
}

type CursorInterface interface {
	Close(context.Context) error
	Next(context.Context) bool
	Decode(interface{}) error
	Err() error
}
type InsertOneResultInterface interface{}
type DeleteResultInterface interface{}
type UpdateResultInterface interface{}

func (mc *postCursor) Close(ctx context.Context) error {
	return mc.cursor.Close(ctx)
}
func (mc *postCursor) Next(ctx context.Context) bool {
	return mc.cursor.Next(ctx)
}
func (mc *postCursor) Decode(val interface{}) error {
	return mc.cursor.Decode(val)
}
func (mc *postCursor) Err() error {
	return mc.cursor.Err()
}

func (p *postDatabase) Find(ctx context.Context, i interface{}, findOptions ...*options.FindOptions) (CursorInterface, error) {
	result, err := p.db.Find(ctx, i, findOptions...)
	return &postCursor{cursor: result}, err
}
func (p *postDatabase) InsertOne(ctx context.Context, i interface{}, oneOptions ...*options.InsertOneOptions) (InsertOneResultInterface, error) {
	result, err := p.db.InsertOne(ctx, i, oneOptions...)
	return &postInsertOneResult{result: result}, err
}
func (p *postDatabase) DeleteOne(ctx context.Context, i interface{}, deleteOptions ...*options.DeleteOptions) (DeleteResultInterface, error) {
	result, err := p.db.DeleteOne(ctx, i, deleteOptions...)
	return &postDeleteResult{result: result}, err
}
func (p *postDatabase) ReplaceOne(ctx context.Context, i interface{}, i2 interface{}, replaceOptions ...*options.ReplaceOptions) (UpdateResultInterface, error) {
	result, err := p.db.ReplaceOne(ctx, i, i2, replaceOptions...)
	return &postReplaceResult{result: result}, err
}
