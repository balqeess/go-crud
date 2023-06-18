package main

import (
	"github.com/balqees/go-crud/controllers"
	"github.com/balqees/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	//r.Any("/users/form", controllers.FormHandler) // Handles both GET and POST requests
	r.GET("/users/form", controllers.ShowUserForm)
	r.POST("/users", controllers.UserCreate)
	r.GET("/users", controllers.GetUsers)
	r.PUT("/users/:id/Update", controllers.UserUpdate)
	r.GET("/users/:id", controllers.GetUserbyID)
	r.DELETE("/users/:id/delete", controllers.UserDelete)
	r.Run()
}
