package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"github.com/gorilla/context"
)

func TestAddIdea(t *testing.T) {
	// Arrange
	body := url.Values{}
	body2 := url.Values{}
	body2.Set("problem_id", "1")
	body2.Set("action_description", "It is finally long enough buddy but do not be happy yet")
	body2.Set("problems_description", "It is finally long enough buddy but do not be happy yet")
	body2.Set("money_price", "1.0")
	body2.Set("time_price", "1")
	req, err := http.NewRequest("POST", "/ideas", strings.NewReader(body.Encode()))
	req2, err2 := http.NewRequest("POST", "/ideas", strings.NewReader(body2.Encode()))
	if err != nil || err2 != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	rr2 := httptest.NewRecorder()
	ctx := req2.Context()
	handler := http.HandlerFunc(AddIdea)

	// Act
	handler.ServeHTTP(rr, req)
	handler.ServeHTTP(rr2, req2)

	// Assert
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Bad code for empty body, got %v wanted %v", status, http.StatusBadRequest)
	}
}