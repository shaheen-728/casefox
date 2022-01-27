package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shaheen-728/casefox/data"
)
//Test Function for search an article
func SearchArticle(w http.ResponseWriter, r *http.Request) {
   // search an article by query parameter title,subtitle,content
	title := r.URL.Query().Get("title")
    subtitle := r.URL.Query().Get("subtitle")
	content:= r.URL.Query().Get("content")

	// Iterate over all the data articles
	for _, article := range data.Articles {
		if article.Title == title && article.SubTitle == subtitle && article.Content == content{
			// If all parameters are equal send article as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(article)
			break
		}else{
			fmt.Println("Article not found")
		}

	}
}
