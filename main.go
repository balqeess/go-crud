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

	routes := gin.Default()
	routes.LoadHTMLGlob("templates/*.html")
	r := routes.Group("/users")
	r.GET("/form", controllers.ShowUserForm)
	r.GET("/list", controllers.ShowUserList)
	r.POST("/", controllers.UserCreate)
	r.GET("/:id", controllers.ShowUserUpdate)
	r.POST("/:id/update", controllers.UserUpdate)
	r.POST("/:id/delete", controllers.UserDelete)
	routes.Run()

}
