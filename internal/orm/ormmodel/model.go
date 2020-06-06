package ormmodel

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName string
}

type Appointment struct {
	gorm.Model
	Description string
	Date string
	Start string
	End string
	User User
	UserID int
}

