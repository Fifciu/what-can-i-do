package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"github.com/gorilla/context"
	"github.com/fifciu/what-can-i-do/server/models"
)

type mockEntity struct {
	UserId uint
	Name string
}

type mockEntityWrapper struct {
	Entity mockEntity
}

func (mck *mockEntity) GetNewInstance() models.DatabaseType {
	return &mockEntity{}
}

func (mck *mockEntity) SetUserId(uid uint) {
	mck.UserId = uid
}

func (mck *mockEntity) Validate() error {
	return nil
}

func (mck *mockEntity) Save() error {
	return nil
}

func TestAddRecordFactory(t *testing.T) {
	// Arrange
	body := url.Values{}
	bodyMap := map[string]interface{}{
		"user_id": 1,
		"name": "Something",
	}
	bodyJson, _ := json.Marshal(bodyMap)
	req, err := http.NewRequest("POST", "/ideas", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err.Error())
	}
	req2, err2 := http.NewRequest("POST", "/ideas", bytes.NewReader(bodyJson))
	if err2 != nil {
		t.Fatal(err2.Error())
	}
	rr := httptest.NewRecorder()
	rr2 := httptest.NewRecorder()
	context.Set(req2, "CurrentUser", &models.Claims{ID: 2})
	handler := AddRecordFactory(&mockEntity{})

	// Act
	handler.ServeHTTP(rr, req)
	handler.ServeHTTP(rr2, req2)

	// Assert
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Bad code for empty body, got %v wanted %v", status, http.StatusBadRequest)
	}

	responseEntity := &mockEntityWrapper{}
	err = json.NewDecoder(rr2.Body).Decode(responseEntity)
	if err != nil {
		t.Errorf(err.Error())
	}

	if responseEntity.Entity.UserId != 2 {
		t.Errorf("It is possible to add an idea as different user")
	}
}
