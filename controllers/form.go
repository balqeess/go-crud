package controllers

type UserForm struct {
	FirstName       string `form:"FirstName" binding:"required"`
	LastName        string `form:"LastName" binding:"required"`
	Email           string `form:"Email" binding:"required,email"`
	NumberOfTickets int    `form:"NumberOfTickets" binding:"required"`
}
