package controllers

import (
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)


type UserForm struct {
	FirstName string `form:"FirstName" binding:"required"`
	LastName  string `form:"LastName" binding:"required"`
	Email     string `form:"Email" binding:"required,email"`
	DateOfBirth time.Time `form:"DateOfBirth" binding:"required" time_format:"2006-01-02"`
}


type UserUpdateForm struct {
	FirstName string `form:"FirstName" binding:"required"`
	LastName  string `form:"LastName" binding:"required"`
	Email     string `form:"Email" binding:"required,email"`
	DateOfBirth time.Time `form:"DateOfBirth" binding:"required" time_format:"2006-01-02"`
}


// the following functions are responsible for handling the route to display:
// User Form

func ShowUserForm(c *gin.Context) {
	formHTML := RenderFormHTML()
	c.HTML(http.StatusOK, "base.html", gin.H{
		"content": template.HTML(formHTML), // Pass the form HTML as a string to be rendered as raw HTML
	})
}



func ShowUserList(c *gin.Context) {
	// Get the users
	users := GetUsers()
	totalUsers := len(users)

	// Calculate age and sum of ages for each user
	var totalAge int
	for i := range users {
		age := CalculateAge(users[i].DateOfBirth)
		users[i].Age = age
		totalAge += age
	}

	// Calculate the average age
	var averageAge int
	if totalUsers > 0 {
		averageAge = (totalAge) / (totalUsers)
	}

	// passes users to generate the HTML content for the user list.
	userListHTML := RenderUserListHTML(users)

	// Render the base.html template with the user list HTML content
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":   "User List",
		"content": template.HTML(userListHTML),
		"users":   users,
		"totalUsers": totalUsers,
		"averageAge": averageAge,
	})
}


// Update form
func ShowUserUpdate(c *gin.Context) {
	userID := c.Param("id")
	user := GetUserByID(userID)
	if user == nil {
		// Handle the case where the user is not found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	updateHTML := RenderUpdateHTML()

	replacements := map[string]string{
        "{{ .user.ID }}":       userID,
        "{{ .user.FirstName }}": user.FirstName,
        "{{ .user.LastName }}":  user.LastName,
        "{{ .user.Email }}":     user.Email,
		"{{ .user.DateOfBirth }}": user.DateOfBirth.Format("2006-01-02"),

    }

    renderedHTML := updateHTML
    for tmplVar, value := range replacements {
        renderedHTML = strings.ReplaceAll(renderedHTML, tmplVar, value)
    }

	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":   "User Update Form",
		"content": template.HTML(renderedHTML),
	
	})
}

