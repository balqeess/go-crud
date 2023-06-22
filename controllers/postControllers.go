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


 func GetUsers() []models.User {
	// Retrieve the users from the database 
	var users []models.User
	initializers.DB.Find(&users)
	return users
}


func GetUserByID(userID string) *models.User {
	// Retrieve the user from the database based on the provided ID
	var user models.User
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		// Handle the error if the user is not found or any other database error occurs
		return nil
	}
	return &user
}


func UserUpdate(c *gin.Context) {
	//get id of url
	ID := c.Param("id")
	// get the user by the given id from the url
	user := GetUserByID(ID)
	if user == nil {
		// Handle the case where the user is not found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Parse and bind the form data to the UserUpdateForm struct
	var form UserUpdateForm
	if err := c.ShouldBind(&form); err != nil {
		// Handle the error if form data validation fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user with the new details
	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.Email = form.Email

	// Save the updated user to the database
	if err := initializers.DB.Save(user).Error; err != nil {
		// Handle the error if the user update fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Redirect to the user details page
	c.Redirect(http.StatusSeeOther, "/users/list")
}


func UserDelete(c *gin.Context) {
	// get id off url
	id := c.Param("id")
	// Delete the posts
	initializers.DB.Delete(&models.User{}, id)
	// respond
	c.Redirect(http.StatusFound, "/users/list")
}

