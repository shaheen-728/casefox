package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaheen-728/casefox/handlers"
)
// Main function 
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/articles", handlers.GetAllArticles).Methods(http.MethodGet)
	router.HandleFunc("/articles/{id}", handlers.GetArticle).Methods(http.MethodGet)
	router.HandleFunc("/articles", handlers.AddArticle).Methods(http.MethodPost)
	router.HandleFunc("/articles/search", handlers.SearchArticle).Methods(http.MethodGet)


	log.Println("API is running!")
	http.ListenAndServe(":8080", router)
}
