package server

import (
	"blog/internal/presentation/http"
	"blog/internal/repository"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupRoutes function set all necessary routes
func SetupRoutes(r *mux.Router, articleRepo repository.ArticleRepository) {
	articleHandler := http.NewArticleHandler(articleRepo)

	r.HandleFunc("/articles", articleHandler.CreateArticle).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}", articleHandler.GetArticleByID).Methods("GET")
	r.HandleFunc("/articles", articleHandler.GetAllArticles).Methods("GET")
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler).Methods("GET")
}
