// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/handlers/post.go

// Package handlers is a generated GoMock package.
package handlers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	post 	"my/redditclone/pkg/post"
)

// MockPostsRepo is a mock of PostsRepo interface.
type MockPostsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPostsRepoMockRecorder
}

// MockPostsRepoMockRecorder is the mock recorder for MockPostsRepo.
type MockPostsRepoMockRecorder struct {
	mock *MockPostsRepo
}

// NewMockPostsRepo creates a new mock instance.
func NewMockPostsRepo(ctrl *gomock.Controller) *MockPostsRepo {
	mock := &MockPostsRepo{ctrl: ctrl}
	mock.recorder = &MockPostsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostsRepo) EXPECT() *MockPostsRepoMockRecorder {
	return m.recorder
}

// AddComment mocks base method.
func (m *MockPostsRepo) AddComment(postID, body, authorID, login string) (*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", postID, body, authorID, login)
	ret0, _ := ret[0].(*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment.
func (mr *MockPostsRepoMockRecorder) AddComment(postID, body, authorID, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockPostsRepo)(nil).AddComment), postID, body, authorID, login)
}

// AddPost mocks base method.
func (m *MockPostsRepo) AddPost(postJSON *post.Post, id, login string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPost", postJSON, id, login)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPost indicates an expected call of AddPost.
func (mr *MockPostsRepoMockRecorder) AddPost(postJSON, id, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPost", reflect.TypeOf((*MockPostsRepo)(nil).AddPost), postJSON, id, login)
}

// Category mocks base method.
func (m *MockPostsRepo) Category(name string) ([]*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Category", name)
	ret0, _ := ret[0].([]*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Category indicates an expected call of Category.
func (mr *MockPostsRepoMockRecorder) Category(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Category", reflect.TypeOf((*MockPostsRepo)(nil).Category), name)
}

// Delete mocks base method.
func (m *MockPostsRepo) Delete(postID, authorID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", postID, authorID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostsRepoMockRecorder) Delete(postID, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostsRepo)(nil).Delete), postID, authorID)
}

// DeleteComment mocks base method.
func (m *MockPostsRepo) DeleteComment(postID, commentID, authorID string) (*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", postID, commentID, authorID)
	ret0, _ := ret[0].(*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockPostsRepoMockRecorder) DeleteComment(postID, commentID, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockPostsRepo)(nil).DeleteComment), postID, commentID, authorID)
}

// FindPostsInMongoDB mocks base method.
func (m *MockPostsRepo) FindPostsInMongoDB(bson interface{}) ([]*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPostsInMongoDB", bson)
	ret0, _ := ret[0].([]*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPostsInMongoDB indicates an expected call of FindPostsInMongoDB.
func (mr *MockPostsRepoMockRecorder) FindPostsInMongoDB(bson interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPostsInMongoDB", reflect.TypeOf((*MockPostsRepo)(nil).FindPostsInMongoDB), bson)
}

// GetAll mocks base method.
func (m *MockPostsRepo) GetAll() ([]*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPostsRepoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPostsRepo)(nil).GetAll))
}

// OpenPost mocks base method.
func (m *MockPostsRepo) OpenPost(id string) (*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenPost", id)
	ret0, _ := ret[0].(*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenPost indicates an expected call of OpenPost.
func (mr *MockPostsRepoMockRecorder) OpenPost(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenPost", reflect.TypeOf((*MockPostsRepo)(nil).OpenPost), id)
}

// UnVote mocks base method.
func (m *MockPostsRepo) UnVote(authorID, postID string) (*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnVote", authorID, postID)
	ret0, _ := ret[0].(*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnVote indicates an expected call of UnVote.
func (mr *MockPostsRepoMockRecorder) UnVote(authorID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnVote", reflect.TypeOf((*MockPostsRepo)(nil).UnVote), authorID, postID)
}

// User mocks base method.
func (m *MockPostsRepo) User(name string) ([]*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User", name)
	ret0, _ := ret[0].([]*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// User indicates an expected call of User.
func (mr *MockPostsRepoMockRecorder) User(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockPostsRepo)(nil).User), name)
}

// Vote mocks base method.
func (m *MockPostsRepo) Vote(vote int, authorID, postID string) (*post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vote", vote, authorID, postID)
	ret0, _ := ret[0].(*post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vote indicates an expected call of Vote.
func (mr *MockPostsRepoMockRecorder) Vote(vote, authorID, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vote", reflect.TypeOf((*MockPostsRepo)(nil).Vote), vote, authorID, postID)
}