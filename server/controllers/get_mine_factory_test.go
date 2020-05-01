package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"github.com/gorilla/context"
	"github.com/fifciu/what-can-i-do/server/models"
)

type UceMockEntity struct {
	UserId uint
	Name string
}

func (uce *UceMockEntity) PluralName() string {
	return "Somestr"
}

func (uce *UceMockEntity) GetByUserId(userId uint) []models.UserCreatedEntity {
	ume := []*UceMockEntity{
		{
			UserId: 1,
			Name:   "Sdsds",
		},
		{
			UserId: 1,
			Name:   "Sdsds",
		},
	}

	mapper := make([]models.UserCreatedEntity, len(ume))
	for i, um := range ume {
		mapper[i] = um
	}
	return mapper
}

func TestGetMineFactory(t *testing.T) {
	// Arrange
	body := url.Values{}
	req, err := http.NewRequest("POST", "/ideas", strings.NewReader(body.Encode()))
	if err != nil {
		t.Fatal(err.Error())
	}
	rr := httptest.NewRecorder()
	context.Set(req, "CurrentUser", &models.Claims{ID: 2})
	handler := GetMineFactory(&UceMockEntity{})

	// Act
	handler.ServeHTTP(rr, req)

	// Assert
	responseEntity := make(map[string]interface{})
	mockObj := &UceMockEntity{}
	pluralName := mockObj.PluralName()

	err = json.NewDecoder(rr.Body).Decode(&responseEntity)
	if err != nil {
		t.Errorf(err.Error())
	}

	switch iterative := responseEntity[pluralName].(type) {
	case []*UceMockEntity:
		for _, record := range iterative {
			if record.UserId != 2 {
				t.Errorf("It does not return proper records")
				break
			}
		}
	}
}
