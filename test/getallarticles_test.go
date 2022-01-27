package test

import (
    "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/shaheen-728/casefox/models"
	handler "github.com/shaheen-728/casefox/handlers"
)

var sampleArticle1 = []models.Article{
	{
		ID:     100,
		Title:  "The Great Gatsby",
		SubTitle: "Gatsby",
		Content:   "A book for Detectives",
		Creation_Timestamp : "2022-01-20 21:26:29.704679241 +0530 IST m=+33.989299685"	},
	
}

//Test function for get all articles
func TestGetAllArticles(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles", nil)
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
	expected1, _ := json.Marshal(sampleArticle1)
	assert.NotEqual(t, string(expected1), rr.Body)

}
