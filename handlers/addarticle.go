package handlers

// import all required packages
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/shaheen-728/casefox/data"
	"github.com/shaheen-728/casefox/models"
)

//Func AddArticle is used to add an article by Post Method
func AddArticle(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
    //If error occurs while reading the body
	if err != nil {
		log.Fatalln(err)
	}

	var article models.Article
	json.Unmarshal(body, &article)// unmarshal the request body into article struct

	// Append to the Article data
	article.ID = rand.Intn(100)// rand.Intn gives the random integer ID to article
	article.Creation_Timestamp=time.Now().String()// It gives the current timestamp 
	data.Articles = append(data.Articles, article)// add the new article to articles data

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}
