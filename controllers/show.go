package controllers

import (
	"html/template"
	"net/http"

	"github.com/balqees/go-crud/initializers"
	"github.com/balqees/go-crud/models"
	"github.com/gin-gonic/gin"
)

type UserForm struct {
	FirstName       string `form:"FirstName" binding:"required"`
	LastName        string `form:"LastName" binding:"required"`
	Email           string `form:"Email" binding:"required,email"`
}

func ShowUserForm(c *gin.Context) {
	formHTML := RenderFormHTML()
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":   "Submit User Form",
		"content": template.HTML(formHTML), // Pass the form HTML as a string to be rendered as raw HTML
	})
}

func ShowUserList(c *gin.Context) {
	var users []models.User
    initializers.DB.Find(&users)
	userListHTML := RenderUserListHTML(users)
	c.HTML(http.StatusOK, "base.html", gin.H{
		"title":   "User List",
		"content": template.HTML(userListHTML),
		"users": users,  // Pass the form HTML as a string to be rendered as raw HTML
	})
}
