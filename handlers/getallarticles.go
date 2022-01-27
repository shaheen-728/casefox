package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shaheen-728/casefox/data"
	"github.com/shaheen-728/casefox/models"
)
//Func GetAllArticles is used to get all the articles 
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(data.Articles)


	// Add Pagination to get the articles by limit
	Limit := r.URL.Query().Get("limit")//get the limit by url
	limit, _ := strconv.Atoi(Limit)//convert it into integer
	if limit==0{
		limit=3 //if value of limit is not in url then default value of limit is 3
	}
    Offset := r.URL.Query().Get("offset")//get the offset by url
	offset, _ := strconv.Atoi(Offset)//convert it into integer
	if offset==0{
		offset=1//if value of offset is not in url then default value of offset is 0
	}

    var article_after_pagination []models.Article
	//iterate over the limit value as offset is given
    for val := offset; val <=limit; val++{
	   article_after_pagination=append(article_after_pagination,data.Articles[val])//article is appended in list after pagination

	}
	json.NewEncoder(w).Encode(article_after_pagination)
	w.WriteHeader(http.StatusOK) //send statusOK

}
