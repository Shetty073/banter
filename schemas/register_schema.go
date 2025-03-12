package schemas

import (
	"errors"
	"time"
)

type RegisterSchema struct {
	Username     string `json:"username" binding:"required,alphanum"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=8,max=32"`
	FirstName    string `json:"first_name" binding:"required,min=2,max=50,alpha"`
	LastName     string `json:"last_name" binding:"required,min=2,max=50,alpha"`
	DateOfBirth  string `json:"date_of_birth" binding:"required"` // Keep as string
	Gender       string `json:"gender" binding:"required,oneof=male female other"`
	MobileNumber string `json:"mobile_number" binding:"required,len=10,numeric"`
}

// Function to parse DateOfBirth before saving
func (r *RegisterSchema) GetParsedDOB() (time.Time, error) {
	// Convert "24-03-1996" (DD-MM-YYYY) to Go's time.Time
	dob, err := time.Parse("02-01-2006", r.DateOfBirth)
	if err != nil {
		return time.Time{}, errors.New("invalid date format, expected DD-MM-YYYY")
	}
	return dob, nil
}
