package http

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"

	"blog/internal/model"
	"blog/internal/service"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func TestArticleHandler_CreateArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockArticleService := service.NewMockArticleService(ctrl)

	articleHandler := NewArticleHandler(mockArticleService)

	article := &model.Article{
		Title:   "Sample Article",
		Content: "This is a sample article",
	}

	articleJSON, _ := json.Marshal(article)

	request, _ := http.NewRequest("POST", "/articles", bytes.NewReader(articleJSON))

	mockArticleService.EXPECT().CreateArticle(article).Return(1, nil)

	response := httptest.NewRecorder()

	articleHandler.CreateArticle(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestArticleHandler_GetArticleByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockArticleService := service.NewMockArticleService(ctrl)

	articleHandler := NewArticleHandler(mockArticleService)

	router := mux.NewRouter()
	router.HandleFunc("/article/{id:[0-9]+}", articleHandler.GetArticleByID).Methods("GET")
	request, _ := http.NewRequest("GET", "/article/1", nil)

	mockArticleService.EXPECT().GetArticleByID(1).Return(&model.Article{}, nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestArticleHandler_GetAllArticles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockArticleService := service.NewMockArticleService(ctrl)

	articleHandler := NewArticleHandler(mockArticleService)

	router := mux.NewRouter()
	router.HandleFunc("/articles", articleHandler.GetAllArticles).Methods("GET")
	request, _ := http.NewRequest("GET", "/articles", nil)

	mockArticleService.EXPECT().GetAllArticles().Return([]model.Article{}, nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}
