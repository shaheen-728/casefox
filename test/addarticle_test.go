package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	handler "github.com/shaheen-728/casefox/handlers"
)

//sampleArticle2 is to add an article for testing
var sampleArticle2 = []byte(`{"id":4,"title":"xyz","subtitle":"pqr","content":"xy"}`)

//Test function for Adding Article i.e. article is added or not
func TestAddArticle(t *testing.T) {
	req, err := http.NewRequest("POST", "/article", bytes.NewBuffer(sampleArticle2))
	if err != nil {
		t.Fatal(err) // if error occurs while reading the request
	}
	req.Header.Set("Content-Type", "application/json") //set content-type
	rr := httptest.NewRecorder()                       //create a newrecorder to record the response.
	handler := http.HandlerFunc(handler.AddArticle)
	handler.ServeHTTP(rr, req)                           //reading all the http requests and send to the response body
	if status := rr.Code; status != http.StatusCreated { //if request body code  is not equal to status created
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected1 := `{"id":5,"title":"abc","subtitle":"123","content":"09pip"}`
	assert.NotEqual(t, string(expected1), rr.Body) //test pass if expected string is not equal to request body

}
