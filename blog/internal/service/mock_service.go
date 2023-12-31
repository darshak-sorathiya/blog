// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go

package service

import (
	model "blog/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockArticleService is a mock of ArticleService interface.
type MockArticleService struct {
	ctrl     *gomock.Controller
	recorder *MockArticleServiceMockRecorder
}

// MockArticleServiceMockRecorder is the mock recorder for MockArticleService.
type MockArticleServiceMockRecorder struct {
	mock *MockArticleService
}

// NewMockArticleService creates a new mock instance.
func NewMockArticleService(ctrl *gomock.Controller) *MockArticleService {
	mock := &MockArticleService{ctrl: ctrl}
	mock.recorder = &MockArticleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleService) EXPECT() *MockArticleServiceMockRecorder {
	return m.recorder
}

// CreateArticle mocks base method.
func (m *MockArticleService) CreateArticle(article *model.Article) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", article)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockArticleServiceMockRecorder) CreateArticle(article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockArticleService)(nil).CreateArticle), article)
}

// GetAllArticles mocks base method.
func (m *MockArticleService) GetAllArticles() ([]model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllArticles")
	ret0, _ := ret[0].([]model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllArticles indicates an expected call of GetAllArticles.
func (mr *MockArticleServiceMockRecorder) GetAllArticles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArticles", reflect.TypeOf((*MockArticleService)(nil).GetAllArticles))
}

// GetArticleByID mocks base method.
func (m *MockArticleService) GetArticleByID(id int) (*model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleByID", id)
	ret0, _ := ret[0].(*model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleByID indicates an expected call of GetArticleByID.
func (mr *MockArticleServiceMockRecorder) GetArticleByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleByID", reflect.TypeOf((*MockArticleService)(nil).GetArticleByID), id)
}
