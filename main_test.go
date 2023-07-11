package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/balqees/go-crud/controllers"
	"github.com/balqees/go-crud/initializers"
	"github.com/balqees/go-crud/models"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine{
	r := gin.Default()
	return r
}

func TestUserCreate(t *testing.T) {
	// Create a new Gin router
	r := SetUpRouter()

	// Set up the route
	r.POST("/users", controllers.UserCreate)

	// Create a new HTTP POST request with the test route and form data
	formData := url.Values{
		"FirstName": {"Ahmed"},
		"LastName":  {"AlBarwani"},
		"Email":     {"Ahmed@gmail.com"},
		"DateOfBirth": {"1990-01-01"},
	}
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response status code
	assert.Equal(t, http.StatusSeeOther, w.Code)

	// Assert the response header for the redirect location
	assert.Equal(t, "/users/list", w.Header().Get("Location"))

	// Assert that the user has been created in the database
	 var user models.User
	initializers.DB.Where("email = ?", "Ahmed@gmail.com").First(&user)
	assert.Equal(t, "Ahmed", user.FirstName)
	assert.Equal(t, "AlBarwani", user.LastName)
	assert.Equal(t, "Ahmed@gmail.com", user.Email)
	assert.Equal(t, time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC), user.DateOfBirth.UTC())

	initializers.DB.Delete(&user)

}

func TestGetUsers(t *testing.T) {
	// Create a few test users to be stored in the database
	testUsers := []models.User{
		{
			FirstName:   "Ahmed",
			LastName:    "AlBarwani",
			Email:       "Ahmed@gmail.com",
			DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			FirstName:   "Yusra",
			LastName:    "AlHarthi",
			Email:       "Yusra@gmail.com",
			DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	// Insert the test users into the database
	for i := range testUsers {
		initializers.DB.Create(&testUsers[i])
	}

	// Call the GetUsers function to retrieve all users from the database
	allUsers := controllers.GetUsers()

	// Assert that the retrieved users include the test users
	for _, user := range testUsers {
	// adding this addition step because this test case retrieves multiple users from the database
	// and compares them with the test users  created.
	// The issue lies in the comparison of the retrieved users with the test users.
	// that will include created deleted and updated at for each user which differs


		found := false
		for _, retrievedUser := range allUsers {
			if retrievedUser.FirstName == user.FirstName &&
				retrievedUser.LastName == user.LastName &&
				retrievedUser.Email == user.Email &&
				retrievedUser.DateOfBirth.Equal(user.DateOfBirth) {
				found = true
				break
			}
		}
		assert.True(t, found, "Expected user not retrieved: %+v", user)
	}

	// Clean up by deleting the test users from the database
	for i := range testUsers {
		initializers.DB.Delete(&testUsers[i])
	}
}



func TestGetUserByID(t *testing.T) {
	// Create a new user to be stored in the database
	user := models.User{
		FirstName: "Ahmed",
		LastName:  "AlBarwani",
		Email:     "Ahmed@gmail.com",
		DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	// Create the user in the database
	initializers.DB.Create(&user)

	// Test case: User found in the database
	foundUser := controllers.GetUserByID(strconv.FormatUint(uint64(user.ID), 10))
	assert.Equal(t, user.FirstName, foundUser.FirstName)
	assert.Equal(t, user.LastName, foundUser.LastName)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.Equal(t, user.DateOfBirth.UTC(), foundUser.DateOfBirth.UTC())
	
	// delete the user
	initializers.DB.Delete(&user)
}

func TestUserUpdate(t *testing.T) {
	// Create a new Gin router
	r := SetUpRouter()

	// Set up the route
	r.POST("/users/:id/update", controllers.UserUpdate)

	// Create a new user to be stored in the database
	user := models.User{
		FirstName: "Ahmed",
		LastName:  "AlBarwani",
		Email:     "Ahmed@gmail.com",
		DateOfBirth: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	// Create the user in the database
	initializers.DB.Create(&user)
	// update the user fields as needed 
	updatedUser := models.User{
	FirstName :"Nasser",
	LastName :"AlJabri",
	Email :"Naseer@gmail.com",
	DateOfBirth: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	// Save the updated user to the database
	initializers.DB.Save(&updatedUser)

	// Convert the updated user details to JSON
	jsonData, _ := json.Marshal(updatedUser)

	// Create a new HTTP PUT request with the test route and updated user details
	req, _ := http.NewRequest("POST", "/users/"+strconv.Itoa(int(user.ID))+"/update", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	// Create a response recorder to record the response
	w := httptest.NewRecorder()
	
	// Perform the request
	r.ServeHTTP(w, req)
	
	// Assert the response status code
	assert.Equal(t, http.StatusSeeOther, w.Code)
	
	// Retrieve the updated user from the database
	updatedUserFromDB := controllers.GetUserByID(strconv.Itoa(int(user.ID)))
	
	// Assert that the user details have been updated correctly
	assert.NotNil(t, updatedUserFromDB)
	assert.Equal(t, updatedUser.FirstName, updatedUserFromDB.FirstName)
	assert.Equal(t, updatedUser.LastName, updatedUserFromDB.LastName)
	assert.Equal(t, updatedUser.Email, updatedUserFromDB.Email)
	assert.Equal(t, updatedUser.DateOfBirth.UTC(), updatedUserFromDB.DateOfBirth.UTC())
	
	// Save the updated user to the database
	initializers.DB.Save(updatedUserFromDB)
	
	// Delete the user from the database
	initializers.DB.Delete(&user)

}
func TestUserDelete(t *testing.T) {
	// Create a new Gin router
	r := SetUpRouter()

	// Set up the route
	r.POST("/users/:id/delete", controllers.UserDelete)

	// Create a new HTTP DELETE request with the test route
	req, _ := http.NewRequest("POST", "/users/123/delete", nil)

	// Create a response recorder to record the response
	w := httptest.NewRecorder()
	
	// Perform the request
	r.ServeHTTP(w, req)
	
	// Assert the response status code and location header
	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/users/list", w.Header().Get("Location"))
}

func TestCalculateAge(t *testing.T) {
	dateOfBirth := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	expectedAge := 23

	age := controllers.CalculateAge(dateOfBirth)

	assert.Equal(t, expectedAge, age)
}
