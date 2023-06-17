package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	FirstName       string `form:"FirstName" binding:"required"`
	LastName        string `form:"LastName" binding:"required"`
	Email           string `form:"Email" binding:"required,email"`
	NumberOfTickets int    `form:"NumberOfTickets" binding:"required"`
}

func ShowForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{
        "title": "Submit User Form",
    })
}

