package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shaheen-728/casefox/data"
)
//Function SearchArticle is used to  search  for an article
func SearchArticle(w http.ResponseWriter, r *http.Request) {
   // search an article by query parameter title,subtitle,content
	title := r.URL.Query().Get("title")//retrieve title by Query paramter url
    subtitle := r.URL.Query().Get("subtitle")//retrieve subtitle by Query parameter url
	content:= r.URL.Query().Get("content")//retrieve content by Query parameter url

	// Iterate over all the data articles
	for _, article := range data.Articles {
		if article.Title == title && article.SubTitle == subtitle && article.Content == content{
			// If all parameters are equal send article as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(article)
			break //out of the loop if article is found
		}else{
			fmt.Println("Article not found")//print if article is not found
		}

	}
}
