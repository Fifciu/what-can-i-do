package models

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestGetUserByEmail (t *testing.T) {
	// Arrange
	email := "zardziol@gmail.com"
	Db, mock, _ := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
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
	Db, mock, _ := sqlmock.New()
	db, _ = gorm.Open("mysql", Db)
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
