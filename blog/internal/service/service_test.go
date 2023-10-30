package service

import (
	"blog/internal/model"
	"blog/internal/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArticleService_CreateArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockArticleRepository(ctrl)

	articleService := NewArticleService(mockRepo)

	article := &model.Article{
		Title:   "Sample Article",
		Content: "This is a sample article",
	}

	mockRepo.EXPECT().CreateArticle(article).Return(1, nil)

	id, err := articleService.CreateArticle(article)

	assert.Nil(t, err)
	assert.Equal(t, 1, id)
}

func TestArticleService_GetArticleByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockArticleRepository(ctrl)

	articleService := NewArticleService(mockRepo)

	mockRepo.EXPECT().GetArticleByID(1).Return(&model.Article{}, nil)

	article, err := articleService.GetArticleByID(1)

	assert.Nil(t, err)
	assert.NotNil(t, article)
}

func TestArticleService_GetAllArticles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockArticleRepository(ctrl)

	articleService := NewArticleService(mockRepo)

	mockRepo.EXPECT().GetAllArticles().Return([]model.Article{}, nil)

	articles, err := articleService.GetAllArticles()

	assert.Nil(t, err)
	assert.NotNil(t, articles)
}
