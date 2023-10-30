package service

import (
	"blog/internal/model"
	"blog/internal/repository"
)

type ArticleService interface {
	CreateArticle(article *model.Article) (int, error)
	GetArticleByID(id int) (*model.Article, error)
	GetAllArticles() ([]model.Article, error)
}

type articleService struct {
	repo repository.ArticleRepository
}

// NewArticleService creates article service
func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo}
}

// CreateArticle adds article in db and return success if entry is successful
func (uc *articleService) CreateArticle(article *model.Article) (int, error) {
	id, err := uc.repo.CreateArticle(article)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetArticleByID returns article for given ID
func (uc *articleService) GetArticleByID(id int) (*model.Article, error) {
	article, err := uc.repo.GetArticleByID(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

// GetAllArticles returns list of articles stored in database
func (uc *articleService) GetAllArticles() ([]model.Article, error) {
	articles, err := uc.repo.GetAllArticles()
	if err != nil {
		return nil, err
	}
	return articles, nil
}
