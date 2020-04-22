package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestGetAllProblems (t *testing.T) {
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
	}).
		AddRow(1, 1, "Coronavirus", "coron", "adasdsdasdasdasa", 1)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(is_published = 1(.+)").WillReturnRows(sqlRows)

	// Act
	GetAllProblems()

	// Assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}

func TestGetUserProblems (t *testing.T) {
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
		"user_id",
		"name",
		"slug",
		"description",
		"is_published",
	}).
		AddRow(1, userId, "Coronavirus", "coron", "adasdsdasdasdasa", 1)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(user_id = \\?(.+)").WithArgs(userId).WillReturnRows(sqlRows)

	// Act
	GetUserProblems(userId)

	// Assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}

func TestGetProblem (t *testing.T) {
	// Arrange
	problemSlug := "coron"
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
	}).
		AddRow(1, 1, "Coronavirus", problemSlug, "adasdsdasdasdasa", 1)
	sqlRows2 := sqlmock.NewRows([]string{
		"id",
		"user_id",
		"name",
		"slug",
		"description",
		"is_published",
	}).
		AddRow(1, 1, "Coronavirus", problemSlug, "adasdsdasdasdasa", 1)
	sqlRows3 := sqlmock.NewRows([]string{
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
	sqlRows4 := sqlmock.NewRows([]string{
		"id",
		"email",
		"fullname",
		"provider",
		"flags",
	}).
		AddRow(1, "zardziol@gmail.com", "Tony", "google", 0)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(slug = \\? AND is_published = 1(.+)").WithArgs(problemSlug).WillReturnRows(sqlRows)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(slug = \\? AND is_published = 1(.+)").WithArgs(problemSlug).WillReturnRows(sqlRows2)
	mock.ExpectQuery("^SELECT (.+) FROM `ideas` WHERE \\(problem_id = \\?(.+)").WithArgs(1).WillReturnRows(sqlRows3)
	mock.ExpectQuery("^SELECT (.+) FROM `users` WHERE \\(id IN \\(\\?\\)(.+)").WithArgs(1).WillReturnRows(sqlRows4)

	// Act
	GetProblem(problemSlug, false)
	GetProblem(problemSlug, true)

	// Assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}

func TestGetProblemsByQuery (t *testing.T) {
	// Arrange
	query := "cor"
	queryPattern := "%" + query + "%"
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
	}).
		AddRow(1, 1, "Coronavirus", "coron", "adasdsdasdasdasa", 1)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(name LIKE \\? AND is_published = 1").WithArgs(queryPattern).WillReturnRows(sqlRows)

	// Act
	GetProblemsByQuery(query)

	// Assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}

func TestProblemExists (t *testing.T) {
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
		"user_id",
		"name",
		"slug",
		"description",
		"is_published",
	}).
		AddRow(problemId, 1, "Coronavirus", "coron", "adasdsdasdasdasa", 1)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(id = \\?(.+)").WithArgs(problemId).WillReturnRows(sqlRows)

	// Act
	ProblemExists(problemId)

	// Assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}

func TestProblemSave (t *testing.T) {
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
		"user_id",
		"name",
		"slug",
		"description",
		"is_published",
	}).
		AddRow(problemId, 1, "Coronavirus", "coron", "adasdsdasdasdasa", 1)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(name = \\? AND description = \\?(.+)").WithArgs("Coronavirus", "adasdsdasdasdasa").WillReturnRows(sqlRows)
	mock.ExpectQuery("^SELECT (.+) FROM `problems` WHERE \\(name = \\? AND description = \\?(.+)").WithArgs("Properone", "adasdsdasdasdasa").WillReturnRows(sqlRows)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `problems`").WithArgs(1, "Properone", "properone", "adasdsdasdasdasa", false).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	problemInstance := &Problem{}
	success1 := problemInstance.Save(1, "Coronavirus", "adasdsdasdasdasa")
	success2 := problemInstance.Save(1, "Properone", "adasdsdasdasdasa")

	if success1 {
		t.Errorf("Anti same payload spam filter did not work")
	}

	if !success2 {
		t.Errorf("Could not add new problem")
	}

	// Assert
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}