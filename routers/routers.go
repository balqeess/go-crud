package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/balqees/go-crud/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	// Group routes under /users
	users := r.Group("/users")
	{
		users.GET("/form", controllers.ShowUserForm)
		users.GET("/list", controllers.ShowUserList)
		users.POST("/", controllers.UserCreate)
		users.GET("/:id", controllers.ShowUserUpdate)
		users.POST("/:id/update", controllers.UserUpdate)
		users.POST("/:id/delete", controllers.UserDelete)
	}

	return r
}