package models

import 	"github.com/dgrijalva/jwt-go"

type Claims struct {
	ID uint `json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type ClaimsUser struct {
	ID uint `json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
}

type UserCreatedEntity interface {
	GetByUserId(userId uint) UserCreatedEntity
	PluralName() string
}