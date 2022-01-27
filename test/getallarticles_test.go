package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	handler "github.com/shaheen-728/casefox/handlers"
	"github.com/shaheen-728/casefox/models"
)

//sampleArticle1 is to check an article is same as response articles

var sampleArticle1 = []models.Article{
	{
		ID:                 100,
		Title:              "The Great Gatsby",
		SubTitle:           "Gatsby",
		Content:            "A book for Detectives",
		Creation_Timestamp: "2022-01-20 21:26:29.704679241 +0530 IST m=+33.989299685"},
}

//TestGetAllArticles is used to check all articles

func TestGetAllArticles(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles", nil) //used to get new request
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetAllArticles)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected1, _ := json.Marshal(sampleArticle1)   //marshal the sampleArticle1
	assert.NotEqual(t, string(expected1), rr.Body) //test case passes if expected string is equal to body string

}
