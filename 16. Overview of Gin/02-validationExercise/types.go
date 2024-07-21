package main

import "time"

type User struct {
	FirstName string `form:"first_name" binding:"required,min=2,max=16"`
	LastName string `form:"last_name" binding:"required,min=2,max=16"`
	Username string `form:"username" binding:"required,alphanum,min=8,max=32"`
	Email string `form:"email" binding:"required,email,min=8,max=32"`
	PhoneNumber string `form:"phone_number" binding:"required,number,len=11,startswith=09"`
	BirthDate time.Time `form:"birth_date" binding:"required,validbirthdate" time_format:"2006/01/02"`
	NationalID string `form:"national_id" binding:"required,number,validnationalid,len=10"`
}