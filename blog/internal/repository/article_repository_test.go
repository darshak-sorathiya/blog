package repository

import (
	"blog/internal/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewArticleRepository(db)

	testArticle := model.Article{
		Title:   "Test Article",
		Content: "This is a test article",
		Author:  "Test Author",
	}

	mock.ExpectExec("INSERT INTO articles (.+) VALUES (.+)").
		WithArgs(testArticle.Title, testArticle.Content, testArticle.Author).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.CreateArticle(&testArticle)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestGetArticleByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewArticleRepository(db)

	testArticle := model.Article{
		ID:      1,
		Title:   "Test Article",
		Content: "This is a test article",
		Author:  "Test Author",
	}

	mock.ExpectQuery("SELECT id, title, content, author FROM articles WHERE id = ?").
		WithArgs(testArticle.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "author"}).
			AddRow(testArticle.ID, testArticle.Title, testArticle.Content, testArticle.Author))

	article, err := repo.GetArticleByID(testArticle.ID)

	assert.NoError(t, err)
	assert.Equal(t, testArticle, *article)
}

func TestGetAllArticles(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewArticleRepository(db)

	testArticles := []model.Article{
		{
			ID:      1,
			Title:   "Article 1",
			Content: "Content 1",
			Author:  "Author 1",
		},
		{
			ID:      2,
			Title:   "Article 2",
			Content: "Content 2",
			Author:  "Author 2",
		},
	}

	mock.ExpectQuery("SELECT id, title, content, author FROM articles").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "author"}).
			AddRow(testArticles[0].ID, testArticles[0].Title, testArticles[0].Content, testArticles[0].Author).
			AddRow(testArticles[1].ID, testArticles[1].Title, testArticles[1].Content, testArticles[1].Author))

	articles, err := repo.GetAllArticles()

	assert.NoError(t, err)
	assert.Equal(t, testArticles, articles)
}
