package models

import (
"fmt"
"os"

"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
"github.com/joho/godotenv"
)


var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")

	dbUri := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName)

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	//db.Debug().AutoMigrate(&User{}, &Message{})
}

func GetDB() *gorm.DB {
	return db
}