package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shaheen-728/casefox/data"
	"github.com/shaheen-728/casefox/models"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// Add Pagination
	Limit := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(Limit)
    Offset := r.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(Offset)

    var article_after_pagination []models.Article
    for i := offset; i < limit; i++{
	   article_after_pagination=append(article_after_pagination,data.Articles[i])
	   json.NewEncoder(w).Encode(article_after_pagination)

	}
	json.NewEncoder(w).Encode(article_after_pagination)
	w.WriteHeader(http.StatusOK)

}
