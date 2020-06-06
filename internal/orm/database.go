package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vikusku/book-freelancer/internal/orm/ormmodel"
	"log"
)

// Replace with actual values
var dbUsername, dbPassword, dbName = "username", "password", "dbname"

func OpenDB() *gorm.DB{
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbName))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Opened DB connection")

	if !db.HasTable(&ormmodel.User{}) {
		db.CreateTable(&ormmodel.User{})
	}

	if !db.HasTable(&ormmodel.Appointment{}) {
		db.CreateTable(&ormmodel.Appointment{})
	}

	return db
}