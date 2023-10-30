package http

import (
	"blog/internal/model"
	"blog/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"strconv"
)

// ArticleHandler handles HTTP requests related to articles
type ArticleHandler struct {
	service service.ArticleService
}

// NewArticleHandler creates a new instance of ArticleHandler
func NewArticleHandler(service service.ArticleService) *ArticleHandler {
	return &ArticleHandler{service}
}

// CreateArticle creates article and enters in db
// CreateArticle ...
//
//	@Summary		Create Article
//	@Description	Create New Article
//	@Tags			External
//	@Produce		json
//	@Param			body body model.Article true "Article"
//	@Success		201				{object}	model.SuccessResponseModel
//	@Failure		400				{object}	model.ResponseError
//	@Router			/articles [POST]
func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		log.Println("Create article failed with error" + err.Error())
		writeErrorResponse(w, invalidReqErrMsg, http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateArticle(&article)
	if err != nil {
		log.Println("Create article failed with error" + err.Error())
		writeErrorResponse(w, articleCreationFailureErrMsg, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.SuccessResponse(201, "Success", map[string]int{"id": id}))
}

// GetArticleByID returns article for given ID
// GetArticleByID ...
//
//	@Summary		Return article for given ID
//	@Description	Return article for given ID
//	@Tags			External
//	@Produce		json
//	@Param			id path int true "id"
//	@Success		200				{object}	model.SuccessResponseModel
//	@Failure		400				{object}	model.ResponseError
//	@Router			/article/{id} [GET]
func (h *ArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	articleID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("GetArticleByID failed for ID" + id + " with error" + err.Error())
		writeErrorResponse(w, articleNotFoundErrMsg, http.StatusBadRequest)
		return
	}

	article, err := h.service.GetArticleByID(articleID)

	if err != nil {
		log.Println("GetArticleByID failed for ID" + id + " with error" + err.Error())
		writeErrorResponse(w, articleNotFoundErrMsg, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	writeSuccessResponse(w, []model.Article{*article})
}

// GetAllArticles returns list of all articles
// GetAllArticles ...
//
//	@Summary		Return list of articles
//	@Description	Return list of articles
//	@Tags			External
//	@Produce		json
//	@Success		200				{object}	model.SuccessResponseModel
//	@Failure		400				{object}	model.ResponseError
//	@Router			/articles [GET]
func (h *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := h.service.GetAllArticles()
	if err != nil {
		log.Println("GetArticleByID failed with error" + err.Error())
		writeErrorResponse(w, articleRetrivalFailureErrMsg, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	writeSuccessResponse(w, articles)
}

func writeErrorResponse(w http.ResponseWriter, errMsg string, errCode int) {
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(model.Error(errMsg, errCode))
}

func writeSuccessResponse(w http.ResponseWriter, articles []model.Article) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse(200, "Success", articles))
}
