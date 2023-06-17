package main
import (
	//"net/http"

	"github.com/balqees/go-crud/controllers"
	//"github.com/balqees/go-crud/data"
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
	r.GET("/users/form", controllers.ShowForm)
	r.POST("/users", controllers.UserCreate)
	r.GET("/users", controllers.GetUsers)
	r.PUT("/users/:id", controllers.UserUpdate)
	r.GET("/users/:id", controllers.GetUserbyID)
	r.DELETE("/users/:id", controllers.UserDelete)
	r.Run()
}

