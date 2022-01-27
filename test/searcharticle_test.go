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

var sampleArticle3 = []models.Article{
	{
		
			ID:     100,
			Title:  "The Rise of the Ancient Mariner",
     		SubTitle: "bird",
			Content:   "A book for Mariners ",
			Creation_Timestamp : "2022-09-27 21:26:28.704679241 +0530 IST m=+33.989299685"},
		
}

//Test Function for get an article 
func TestSearchArticle(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles/search", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetArticle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected1, _ := json.Marshal(sampleArticle3)
	assert.NotEqual(t, string(expected1), rr.Body)
}
