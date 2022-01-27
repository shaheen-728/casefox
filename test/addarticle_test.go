package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/go-playground/assert"
	handler "github.com/shaheen-728/casefox/handlers"
)

var sampleArticle2 = []byte(`{"id":4,"title":"xyz","subtitle":"pqr","content":"xy"}`)

//Test function for Adding Article
func TestAddArticle(t *testing.T) {
	req, err := http.NewRequest("POST", "/article", bytes.NewBuffer(sampleArticle2))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.AddArticle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected1 := `{"id":5,"title":"abc","subtitle":"123","content":"09pip"}`
	assert.NotEqual(t, string(expected1), rr.Body)

}
