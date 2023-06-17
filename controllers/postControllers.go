package controllers

import (
	"net/http"

	"github.com/balqees/go-crud/initializers"
	"github.com/balqees/go-crud/models"
	"github.com/gin-gonic/gin"
	
)

func UserCreate(c *gin.Context) {
	// Bind form data to the UserForm struct
	var userForm UserForm
	if err := c.ShouldBind(&userForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new user based on the form data
	user := models.User{
		FirstName:       userForm.FirstName,
		LastName:        userForm.LastName,
		Email:           userForm.Email,
		NumberOfTickets: userForm.NumberOfTickets,
	}

	

	// Create the user in the database
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	
	// retrun it
	c.JSON(200, gin.H{
		"message": user,
	})
	

	
}

// FETCH ALL POSTS FUNCTION 
func GetUsers(c *gin.Context) {
    // Get the users
    var users []models.User
    initializers.DB.Find(&users)

	// respond with them 
	c.JSON(200, gin.H{
		"message": users,
	})

    
}

// FETCH A SINGLE POST FUNCTION
func GetUserbyID(c *gin.Context){
	// get id off url 
	id := c.Param("id")
	// get the posts (find the posts)
	var user models.User
	initializers.DB.First(&user, id)
	// respond with them 
	c.JSON(200, gin.H{
		"post": user,
	})
}
// UPDATE A POST 
func UserUpdate(c *gin.Context){
	// get id off url 
	id := c.Param("id")

	// get data off req body
	var user struct {
		FirstName       string 
		LastName        string 
		Email           string 
		NumberOfTickets	int
	}
	c.Bind(&user)

	// find the post we are updating(same as getting it )
	var post models.User
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.User{
		FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, NumberOfTickets: user.NumberOfTickets,
	})

	//  Respond with it 
	c.JSON(200, gin.H{
		"post": post,
	})
}

// to delete a post 
func UserDelete(c *gin.Context){
	// get id off url 
	id := c.Param("id")
	// Delete the posts
	initializers.DB.Delete(&models.User{}, id)
	// respond 
	c.Status(200)
}
