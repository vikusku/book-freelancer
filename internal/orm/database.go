package orm

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vikusku/book-freelancer/internal/orm/ormmodel"
	"github.com/vikusku/book-freelancer/internal/secretmanager"
	"log"
)

type dbCredentials struct {
	Username string
	Password string
	Engine string
	Host string
	Port string
	DbName string
}

func OpenDB() *gorm.DB{
	secretValue, err := secretmanager.GetSecretValue("db_credentials")
	if err != nil {
		panic(err)
	}

	var credentials dbCredentials
	err = json.Unmarshal([]byte(secretValue), &credentials)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", credentials.Username, credentials.Password, credentials.DbName))
	if err != nil {
		panic(err)
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