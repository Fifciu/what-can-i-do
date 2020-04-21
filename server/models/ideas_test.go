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
	mock.ExpectQuery("^SELECT (.+) FROM `ideas` WHERE \\(problem_id = \\?(.+)").WithArgs(problemId).WillReturnRows(sqlRows)

	// Act
	ideas := GetProblemIdeas(problemId)

	// Assert
	for _, idea := range ideas {
		if idea.ProblemID != problemId {
			t.Errorf("It fetches ideas with bad Problem ID")
		}
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}

func TestGetUserIdeas(t *testing.T) {
	// Arrange
	userId := uint(1)
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
	mock.ExpectQuery("^SELECT (.+) FROM `ideas` WHERE \\(user_id = \\?\\)").WithArgs(userId).WillReturnRows(sqlRows)

	// Act
	ideas := GetUserIdeas(userId, []IdeasMapper{})

	// Assert
	for _, idea := range ideas {
		if idea.UserID != userId {
			t.Errorf("It fetches ideas with bad Problem ID")
		}
	}
}

func TestMapperAddProblemsName(t *testing.T) {
	// Arrange
	Db, mock, _ := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	sqlRows := sqlmock.NewRows([]string{
		"id",
		"name",
	}).
		AddRow(1, "Coronavirus").
		AddRow(2, "Pollution")
	mock.ExpectQuery("^SELECT (.+) FROM `problems` (.+) \\(id IN \\(\\?,\\?,\\?\\)(.+)").WithArgs(uint(1), uint(2), uint(4)).WillReturnRows(sqlRows)
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
	MapperAddProblemsName(ideas)

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

func TestSaveIdea(t *testing.T) {
	// Arrange
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
	mock.ExpectQuery("^SELECT (.+) FROM `ideas` (.+)action_description = \\? AND problem_id = \\?(.+)").WithArgs("Test 1a", 1).WillReturnRows(sqlRows)
	mock.ExpectQuery("^SELECT (.+) FROM `ideas` (.+)action_description = \\? AND problem_id = \\?(.+)").WithArgs("Test 2a", 1).WillReturnRows(sqlRows)

	// Act
	idea := Idea{}
	success1 := idea.Save(1, 1, "Test 1a", "Asds", 12.1, 0)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `ideas`").WithArgs(1, 1, false, "Test 2a", "Asds", float32(12.1), 0).WillReturnResult(sqlmock.NewResult(1, 1))
	success2 := idea.Save(1, 1, "Test 2a", "Asds", 12.1, 0)

	// Assert
	if success1 {
		t.Errorf("It allows to add duplicated as existing idea's action")
	}
	if !success2 {
		t.Errorf("It does not allow to add an idea")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Badly added idea. There were unfulfilled expectations: %s", err)
	}
}