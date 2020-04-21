package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestGetUserByEmail (t *testing.T) {
	// Arrange
	email := "zardziol@gmail.com"
	Db, mock, err := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlRows := sqlmock.NewRows([]string{
		"id",
		"email",
		"fullname",
		"provider",
		"flags",
	}).
		AddRow(1, "zardziol@gmail.com", "Tony", "google", 0)
	mock.ExpectQuery("^SELECT (.+) FROM `users` WHERE \\(email = \\?(.+)").WithArgs(email).WillReturnRows(sqlRows)

	// Act
	user := GetUserByEmail(email)

	// Assert
	if user.Email != email || user.Fullname != "Tony" || user.Provider != "google" || user.Flags != 0 {
		t.Errorf("Not full user returned")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}

func TestGetUserById(t *testing.T) {
	// Arrange
	id := uint(1)
	Db, mock, err := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlRows := sqlmock.NewRows([]string{
		"id",
		"email",
		"fullname",
		"provider",
		"flags",
	}).
		AddRow(1, "zardziol@gmail.com", "Tony", "google", 0)
	mock.ExpectQuery("^SELECT (.+) FROM `users` WHERE \\(id = \\?(.+)").WithArgs(id).WillReturnRows(sqlRows)

	// Act
	user := GetUserById(id)

	// Assert
	if user.Email != "zardziol@gmail.com" || user.Fullname != "Tony" || user.Provider != "google" || user.Flags != 0 {
		t.Errorf("Not full user returned")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}

func TestCreateOrGet(t *testing.T) {
	// Arrange
	email := "zardziol@gmail.com"
	email2 := "new-zardziol@gmail.com"
	Db, mock, err := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlRows := sqlmock.NewRows([]string{
		"id",
		"email",
		"fullname",
		"provider",
		"flags",
	}).
		AddRow(1, "zardziol@gmail.com", "Tony", "google", 0)
	sqlRows2 := sqlmock.NewRows([]string{
		"id",
		"email",
		"fullname",
		"provider",
		"flags",
	})
	mock.ExpectQuery("^SELECT (.+) FROM `users` WHERE \\(email = \\?(.+)").WithArgs(email).WillReturnRows(sqlRows)
	mock.ExpectQuery("^SELECT (.+) FROM `users` WHERE \\(email = \\?(.+)").WithArgs(email2).WillReturnRows(sqlRows2)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").WithArgs(email2, "Abc", "google", 0).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	userType := &User{}
	user, err := userType.CreateOrGet(email, "Abc", "google")
	user2, err2 := userType.CreateOrGet(email2, "Abc", "google")

	// Assert
	if user.Email != email || user.Fullname != "Tony" || user.Provider != "google" || user.Flags != 0 {
		t.Errorf("Not full user returned")
	}

	if err2 != nil || user2.Email != email2 || user2.Fullname != "Abc" || user.Provider != "google" || user.Flags != 0 {
		t.Errorf("Bad user returned after create")
	}

	if err != nil || err2 != nil{
		t.Errorf("Error appeared")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Bad select query. There were unfulfilled expectations: %s", err)
	}
}