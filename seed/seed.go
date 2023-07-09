package seed

import (
	"time"

	"github.com/balqees/go-crud/models"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []models.User{
		{
			FirstName:   "Balqees",
			LastName:    "AlJabri",
			Email:       "Balqees@gmail.com",
			DateOfBirth: time.Date(2001, time.August, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			FirstName:   "Nadia",
			LastName:    "AlSalmi",
			Email:       "Nadia@gmail.com",
			DateOfBirth: time.Date(1992, time.February, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			FirstName:   "Khalifa",
			LastName:    "AlJabri",
			Email:       "Khalifa@gmail.com",
			DateOfBirth: time.Date(1970, time.January, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			FirstName:   "Shihab",
			LastName:    "alWahaibi",
			Email:       "Shihab@gmail.com",
			DateOfBirth: time.Date(2000, time.February, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			FirstName:   "Hasnaa",
			LastName:    "AlRiyami",
			Email:       "Hasnaa@gmail.com",
			DateOfBirth: time.Date(2005, time.February, 29, 0, 0, 0, 0, time.UTC),
		},
		{
			FirstName:   "Yousef",
			LastName:    "AlKindi",
			Email:       "Yousef@gmail.com",
			DateOfBirth: time.Date(2000, time.August, 17, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}
