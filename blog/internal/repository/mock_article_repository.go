// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/article_repository.go

package repository

import (
	model "blog/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockArticleRepository is a mock of ArticleRepository interface.
type MockArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRepositoryMockRecorder
}

// MockArticleRepositoryMockRecorder is the mock recorder for MockArticleRepository.
type MockArticleRepositoryMockRecorder struct {
	mock *MockArticleRepository
}

// NewMockArticleRepository creates a new mock instance.
func NewMockArticleRepository(ctrl *gomock.Controller) *MockArticleRepository {
	mock := &MockArticleRepository{ctrl: ctrl}
	mock.recorder = &MockArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleRepository) EXPECT() *MockArticleRepositoryMockRecorder {
	return m.recorder
}

// CreateArticle mocks base method.
func (m *MockArticleRepository) CreateArticle(article *model.Article) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", article)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockArticleRepositoryMockRecorder) CreateArticle(article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockArticleRepository)(nil).CreateArticle), article)
}

// GetAllArticles mocks base method.
func (m *MockArticleRepository) GetAllArticles() ([]model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllArticles")
	ret0, _ := ret[0].([]model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllArticles indicates an expected call of GetAllArticles.
func (mr *MockArticleRepositoryMockRecorder) GetAllArticles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArticles", reflect.TypeOf((*MockArticleRepository)(nil).GetAllArticles))
}

// GetArticleByID mocks base method.
func (m *MockArticleRepository) GetArticleByID(id int) (*model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleByID", id)
	ret0, _ := ret[0].(*model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleByID indicates an expected call of GetArticleByID.
func (mr *MockArticleRepositoryMockRecorder) GetArticleByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleByID", reflect.TypeOf((*MockArticleRepository)(nil).GetArticleByID), id)
}
