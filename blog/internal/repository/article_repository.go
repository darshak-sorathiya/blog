package repository

import (
	"blog/internal/model"
	"database/sql"
)

type ArticleRepository interface {
	CreateArticle(article *model.Article) (int, error)
	GetArticleByID(id int) (*model.Article, error)
	GetAllArticles() ([]model.Article, error)
}

type MySQLArticleRepository struct {
	db *sql.DB
}

// NewArticleRepository returns db instance of Article Repository
func NewArticleRepository(db *sql.DB) *MySQLArticleRepository {
	return &MySQLArticleRepository{db}
}

// CreateArticle adds article in articles table
func (ar *MySQLArticleRepository) CreateArticle(article *model.Article) (int, error) {
	res, err := ar.db.Exec(
		"INSERT INTO articles (title, content, author) VALUES (?,?,?)",
		article.Title, article.Content, article.Author,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// GetArticleByID fetch article from articles table for give ID
func (ar *MySQLArticleRepository) GetArticleByID(id int) (*model.Article, error) {
	var article model.Article
	err := ar.db.QueryRow(
		"SELECT id, title, content, author FROM articles WHERE id = ?",
		id,
	).Scan(&article.ID, &article.Title, &article.Content, &article.Author)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &article, nil
}

// GetAllArticles fetch all articles stored in articles table
func (ar *MySQLArticleRepository) GetAllArticles() ([]model.Article, error) {
	rows, err := ar.db.Query(
		"SELECT id, title, content, author FROM articles",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []model.Article
	for rows.Next() {
		var article model.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Author)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
