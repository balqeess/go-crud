package tests

import (
	"strconv"
	"testing"

	"github.com/balqees/go-crud/controllers"
	"github.com/balqees/go-crud/initializers"
	"github.com/balqees/go-crud/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	// Connect to the database
	initializers.ConnectToDB()
	
	// Create a test user
	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	initializers.DB.Create(user)


	// Convert user.ID from uint to string
	userID := strconv.Itoa(int(user.ID))

	// Call the GetUserByID function to retrieve the user
	
	retrievedUser := controllers.GetUserByID(userID)

	// Check if the retrieved user matches the original user
	assert.NotNil(t, retrievedUser)
	assert.Equal(t, user.ID, retrievedUser.ID)
	assert.Equal(t, user.FirstName, retrievedUser.FirstName)
	assert.Equal(t, user.LastName, retrievedUser.LastName)
	assert.Equal(t, user.Email, retrievedUser.Email)

	// Clean up: Delete the test user
	initializers.DB.Delete(user)
}


/* func TestUserCreate(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	router.POST("/users", controllers.UserCreate)

	// Create a new multipart/form-data request
	formData := url.Values{
		"FirstName": {"John"},
		"LastName":  {"Doe"},
		"Email":     {"john.doe@example.com"},
	}
	body := strings.NewReader(formData.Encode())

	// Create a new HTTP request with the form data
	req := httptest.NewRequest("POST", "/users", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	// Perform the request and get the response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code and body
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "User created successfully", w.Body.String())
}
 */