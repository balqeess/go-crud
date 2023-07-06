package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName       string
	LastName        string
	Email           string
	DateOfBirth time.Time
	Age         int `gorm:"-"`

}
