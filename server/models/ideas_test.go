package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestGetProblemIdeas(t *testing.T) {
	// Arrange
	problemId := uint(1)
	Db, mock, _ := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	sqlRows := sqlmock.NewRows([]string{
		"id",
		"problem_id",
		"user_id",
		"is_published",
		"action_description",
		"results_description",
		"money_price",
		"time_price",
	}).
		AddRow(1, 1, 1, 1, "Test 1a", "Test 1b", 12.33, 0)
	mock.ExpectQuery("^SELECT (.+) FROM `ideas` (.+)").WillReturnRows(sqlRows)

	// Act
	ideas := GetProblemIdeas(problemId)

	// Assert
	for _, idea := range ideas {
		if idea.ProblemID != problemId {
			t.Errorf("It fetches ideas with bad Problem ID")
		}
	}
}

func TestAddProblemsName(t *testing.T) {
	// Arrange
	Db, mock, _ := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	sqlRows := sqlmock.NewRows([]string{
		"id",
		"name",
	}).
		AddRow(1, "Coronavirus").
		AddRow(2, "Pollution")
	mock.ExpectQuery("^SELECT (.+) FROM `problems` (.+)").WillReturnRows(sqlRows)
	ideas := []*Idea{
		{
			ProblemID: uint(1),
		},
		{
			ProblemID: uint(2),
		},
		{
			ProblemID: uint(4),
		},
	}

	// Act
	addProblemsName(ideas)

	// Assert
	for _, idea := range ideas {
		if (idea.ProblemID == 1 && idea.ProblemName != "Coronavirus") || (idea.ProblemID == 2 && idea.ProblemName != "Pollution") {
			t.Errorf("It fetches ideas with bad Problem ID")
		}
		if (idea.ProblemID == 4 && idea.ProblemName != "Not found") {
			t.Errorf("It does not use fallback value for ProductName")
		}
	}
}