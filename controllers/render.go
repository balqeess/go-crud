package controllers

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"github.com/balqees/go-crud/models"
)

// the following functions are responsible for rendering the HTML files.
// It reads the content of the files using ioutil.ReadFile.


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

	// we use template.Parse function to convert the content of the "userList.html" file 
	//from a string to a template object (tmpl) 
	//that can be used to execute the template and generate the desired output.

	tmpl := template.Must(template.New("userList.html").Parse(string(content)))

	//buffer is used to store the output of executing the template
	var buf bytes.Buffer
	//tmpl.Execute function writes the rendered template output to the buffer.
	// convert the content of buffer to a string
	err = tmpl.Execute(&buf, map[string]interface{}{
		"users": users,
	})
	if err != nil {
		// Handle the error if the template cannot be rendered
		return "Error rendering template"
	}

	return buf.String()
}


func RenderUpdateHTML() string {
	content, err := ioutil.ReadFile("templates/updateUser.html")
	if err != nil {
		// Handle the error if the file cannot be read
		return "Error reading HTML file: " + err.Error()
	}
	
	return string(content)
}

