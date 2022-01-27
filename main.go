package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaheen-728/casefox/handlers"
)
// Main function is execution of api
func main() {
	router := mux.NewRouter() //router is set up by gorilla.mux

	router.HandleFunc("/articles", handlers.GetAllArticles).Methods(http.MethodGet)//router for getting all articles by get method
	router.HandleFunc("/articles/{id}", handlers.GetArticle).Methods(http.MethodGet)//router for get an article by get method
	router.HandleFunc("/articles", handlers.AddArticle).Methods(http.MethodPost)//router for adding an article by post method
	router.HandleFunc("/articles/search", handlers.SearchArticle).Methods(http.MethodGet)//router for search an article by get method


	log.Println("API is running!")//print the log i.e. api is running
	http.ListenAndServe(":8080", router)//bind the address at localhost:8080
}
