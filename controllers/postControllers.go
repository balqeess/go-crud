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
		
	}

	

	// Create the user in the database
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	
	// Redirect to the userList.html page
	c.Redirect(http.StatusSeeOther, "/users/list")
}



func GetUsers(c *gin.Context) gin.H {
	// Get the users
	var users []models.User
	initializers.DB.Find(&users)

	return gin.H{
		"users": users,
	}
}



func GetUserByID(c *gin.Context) models.User{
	// get id off url
	id := c.Param("id")
	// get the posts (find the posts)
	var user models.User
	initializers.DB.First(&user, id)
	// respond with them
	return user
}




func UserDelete(c *gin.Context) {
	// get id off url
	id := c.Param("id")
	// Delete the posts
	initializers.DB.Delete(&models.User{}, id)
	// respond
	c.Redirect(http.StatusFound, "/users/list")
}
