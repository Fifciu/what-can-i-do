package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Email   string    `json:"email"`
	Fullname   string    `json:"fullname"`
	Provider string `json:"provider"`
	Flags uint `json:"flags"`
}

func (user User) TableName() string {
	return "users"
}

func GetUserByEmail(userEmail string) *User {
	user := &User{}

	GetDB().Table("users").Select("*").Where("email = ?", userEmail).First(user)

	return user
}

func GetUserById(userId uint) *User {
	user := &User{}

	GetDB().Table("users").Select("*").Where("id = ?", userId).First(user)

	return user
}

func UserExists (userId uint) bool {
	user := &User{}

	GetDB().Table("users").Select("id").Where("id = ?", userId).First(user)

	if user.ID > 0 {
		return true
	}
	return false
}

func (user *User) CreateOrGet(email string, fullname string, provider string) (*User, error) {
	existingUser := &User{}
	GetDB().Table("users").Select("*").Where("email = ?", email).First(existingUser)

	if existingUser.Email == email {
		return existingUser, nil
	}

	fmt.Println(existingUser)
	fmt.Println("aaa", email, fullname, provider)

	newUser := &User{}

	newUser.Email = email
	newUser.Fullname = fullname
	newUser.Provider = provider
	newUser.Flags = 0
	d := GetDB().Create(newUser)
	if d.Error != nil {
		return nil, errors.New("Couldn't add user to database")
	}
	return newUser, nil
}