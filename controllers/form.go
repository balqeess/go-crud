package controllers


type UserForm struct {
	FirstName       string `form:"FirstName" binding:"required"`
	LastName        string `form:"LastName" binding:"required"`
	Email           string `form:"Email" binding:"required,email"`
	NumberOfTickets int    `form:"NumberOfTickets" binding:"required"`
}

/* func ShowForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{
		"title": "Submit User Form",
	})
}

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
	c.Redirect(http.StatusSeeOther, "/users")
}

func FormHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		ShowForm(c)
	case http.MethodPost:
		UserCreate(c)
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}

 */