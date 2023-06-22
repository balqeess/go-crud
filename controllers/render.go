package controllers

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"github.com/balqees/go-crud/models"
	"github.com/gin-gonic/gin"
)


func RenderFormHTML() string {
	content, err := ioutil.ReadFile("templates/form.html")
	if err != nil {
		// Handle the error if the file cannot be read
		return "Error reading HTML file"
	}
	return string(content)
}


func RenderUserListHTML(users []models.User) string {
	content, err := ioutil.ReadFile("templates/userList.html")
	if err != nil {
		// Handle the error if the file cannot be read
		return "Error reading HTML file"
	}

	tmpl := template.Must(template.New("userList.html").Parse(string(content)))

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, gin.H{
		"users": users,
	})
	if err != nil {
		// Handle the error if the template cannot be rendered
		return "Error rendering template"
	}

	return buf.String()
}
