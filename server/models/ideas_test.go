package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"errors"
)

func TestGetProblemIdeas(t *testing.T) {
	// Arrange
	problemId := uint(1)
	Db, mock, err := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
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

func TestGetByUserId(t *testing.T) {
	// Arrange
	userId := uint(1)
	Db, mock, err := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
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
	idea := &Idea{}
	ideas := idea.GetByUserId(userId)

	// Assert
	for _, idea := range ideas {
		if idea.UserID != userId {
			t.Errorf("It fetches ideas with bad Problem ID")
		}
	}
}

//func TestMapperAddProblemsName(t *testing.T) {
//	// Arrange
//	Db, mock, err := sqlmock.New()
//	db, _ = gorm.Open("mysql", Db)
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//	sqlRows := sqlmock.NewRows([]string{
//		"id",
//		"name",
//	}).
//		AddRow(1, "Coronavirus").
//		AddRow(2, "Pollution")
//	mock.ExpectQuery("^SELECT (.+) FROM `problems` (.+) \\(id IN \\(\\?,\\?,\\?\\)(.+)").WithArgs(uint(1), uint(2), uint(4)).WillReturnRows(sqlRows)
//	ideas := []*Idea{
//		{
//			ProblemID: uint(1),
//		},
//		{
//			ProblemID: uint(2),
//		},
//		{
//			ProblemID: uint(4),
//		},
//	}
//
//	// Act
//	MapperAddProblemsName(ideas)
//
//	// Assert
//	for _, idea := range ideas {
//		if (idea.ProblemID == 1 && idea.ProblemName != "Coronavirus") || (idea.ProblemID == 2 && idea.ProblemName != "Pollution") {
//			t.Errorf("It fetches ideas with bad Problem ID")
//		}
//		if (idea.ProblemID == 4 && idea.ProblemName != "Not found") {
//			t.Errorf("It does not use fallback value for ProductName")
//		}
//	}
//}

func TestSaveIdea(t *testing.T) {
	// Arrange
	Db, mock, err := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
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
		AddRow(1, 1, 1, 1, "Test", "Test 1b", 12.33, 0)
	mock.ExpectQuery("^SELECT (.+) FROM `ideas` (.+)action_description = \\? AND problem_id = \\?(.+)").WithArgs("Test", 1).WillReturnRows(sqlRows)
	mock.ExpectQuery("^SELECT (.+) FROM `ideas` (.+)action_description = \\? AND problem_id = \\?(.+)").WithArgs("Dzong", 1).WillReturnRows(sqlRows)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `ideas`").WithArgs(1, 1, false, "Dzong", "Asds", float32(12.1), 0).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	idea := Idea{
		ProblemID: 1,
		UserID: 1,
		ActionDescription: "Test",
		ResultsDescription: "Asds",
		MoneyPrice: 12.1,
		TimePrice: 0,
	}
	idea2 := Idea{
		ProblemID: 1,
		UserID: 1,
		ActionDescription: "Dzong",
		ResultsDescription: "Asds",
		MoneyPrice: 12.1,
		TimePrice: 0,
	}

	// Act
	err1 := idea.Save()
	err2 := idea2.Save()

	// Assert
	if err1 == nil {
		t.Errorf("It allows to add duplicated as existing idea's action")
	}
	if err2 != nil {
		t.Errorf("It does not allow to add an idea")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Badly added idea. There were unfulfilled expectations: %s", err)
	}
}

func TestValidateIdea (t *testing.T) {
	// Arrange
	Db, mock, err := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlRows := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"name",
		"slug",
		"description",
		"is_published",
	})
	sqlRows2 := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"name",
		"slug",
		"description",
		"is_published",
	}).
		AddRow(1, 1, "Coronavirus", "coron", "adasdsdasdasdasa", 1)
	sqlRows3 := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"name",
		"slug",
		"description",
		"is_published",
	}).
		AddRow(1, 1, "Coronavirus", "coron", "adasdsdasdasdasa", 1)
	sqlRows4 := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"name",
		"slug",
		"description",
		"is_published",
	}).
		AddRow(1, 1, "Coronavirus", "coron", "adasdsdasdasdasa", 1)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(id = \\?(.+)").WithArgs(2).WillReturnRows(sqlRows)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(id = \\?(.+)").WithArgs(1).WillReturnRows(sqlRows2)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(id = \\?(.+)").WithArgs(1).WillReturnRows(sqlRows3)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(id = \\?(.+)").WithArgs(1).WillReturnRows(sqlRows4)
	ideas := []Idea{
		{
			ProblemID: 0,
		},
		{
			ProblemID: 2,
		},
		{
			ProblemID: 1,
			ActionDescription: "Too short",
		},
		{
			ProblemID: 1,
			ActionDescription: "It is finally long enough buddy but do not be happy yet",
			ResultsDescription: "Too short",
		},
		{
			ProblemID: 1,
			ActionDescription: "It is finally long enough buddy but do not be happy yet",
			ResultsDescription: "It is finally long enough buddy but do not be happy yet",
			MoneyPrice: -1.23,
		},
	}

	values := []error{
		errors.New("Bad Problem ID"),
		errors.New("Problem does not exist"),
		errors.New("Idea's action description must have at least 15 characters"),
		errors.New("Idea's results description must have at least 15 characters"),
		errors.New("Idea's price be bigger or equal $0"),
	}

	for i, idea := range ideas {
		err := idea.Validate()
		if err.Error() != values[i].Error() {
			t.Errorf("Bad error, expected %v got %v", values[i].Error(), err.Error())
		}
	}

}