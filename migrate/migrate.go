package main

import (
	"github.com/balqees/go-crud/initializers"
	"github.com/balqees/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.User{})

}

