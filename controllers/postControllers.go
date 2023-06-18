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

	// Redirect to the userList.html page
	c.Redirect(http.StatusFound, "/users")
}
func ShowUserForm(c *gin.Context) {
    c.HTML(http.StatusOK, "form.html", nil)
}

// FETCH ALL POSTS FUNCTION
func GetUsers(c *gin.Context) {
	// Get the users
	var users []models.User
	initializers.DB.Find(&users)

	// respond with them
	c.HTML(http.StatusOK, "userList.html", gin.H{
		"title": "User List",
		"users": users,
	})

}

// FETCH A SINGLE POST FUNCTION
func GetUserbyID(c *gin.Context) {
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
// UserUpdate updates the user details in the database
func UserUpdate(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")

	// Get data from the request body
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user in the database
	result := initializers.DB.Model(&models.User{}).Where("id = ?", userID).Updates(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Redirect to the userList.html page
	c.Redirect(http.StatusFound, "/users")
}

// to delete a post 
func UserDelete(c *gin.Context){
	// get id off url 
	id := c.Param("id")
	// Delete the posts
	initializers.DB.Delete(&models.User{}, id)
	// respond 
	c.Redirect(http.StatusFound, "/users")
}