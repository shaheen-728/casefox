package handlers

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

func AddArticle(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var article models.Article
	json.Unmarshal(body, &article)

	// Append to the Article data
	article.ID = rand.Intn(100)
	article.Creation_Timestamp=time.Now().String()
	data.Articles = append(data.Articles, article)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}
