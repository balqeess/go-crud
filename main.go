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

	gin.SetMode(gin.DebugMode)


	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/users/form", controllers.ShowUserForm)
	r.GET("/users", controllers.ShowUserList)
	r.POST("/users", controllers.UserCreate)
	r.GET("/users/:id/update", controllers.ShowUserUpdateForm)
	r.PUT("/users/:id/update", controllers.UserUpdate)
	r.POST("/users/:id/delete", controllers.UserDelete)

	r.Run()
}
