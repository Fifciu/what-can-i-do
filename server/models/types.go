package models

import 	"github.com/dgrijalva/jwt-go"

type Claims struct {
	ID uint `json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Flags uint `json:"flags"`
	jwt.StandardClaims
}

type ClaimsUser struct {
	ID uint `json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Flags uint `json:"flags"`
}

type UserCreatedEntity interface {
	GetByUserId(userId uint) []UserCreatedEntity
	PluralName() string
}

type DatabaseType interface {
	SetUserId(userId uint)
	Validate() error
	Save() error
	GetNewInstance() DatabaseType
}