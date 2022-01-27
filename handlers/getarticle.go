package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shaheen-728/casefox/data"
)

//Func GetArticle is used to get an article by given id
func GetArticle(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter bu mux
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"]) // convert id into integer

	// Iterate over all the data articles
	for _, article := range data.Articles {
		if article.ID == id {
			// If ids are equal send article as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(article)
			break // get out of the loop if you are get an article
		}
	}
}
