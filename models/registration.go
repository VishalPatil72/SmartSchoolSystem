package models

import (
	"time"
)

type Student struct {
	StudentID   int       `json:"student_id"`
	LastName    string    `json:"last_name"`
	FirstName   string    `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	MotherName  string    `json:"mother_name"`
	Gender      int       `json:"gender"`
	DateOfBirth string    `json:"date_of_birth"` // use string for easier JSON binding (format: YYYY-MM-DD)
	Address     string    `json:"address"`
	PinCode     string    `json:"pin_code"`
	Mobile      string    `json:"mobile"`
	Email       string    `json:"email"`
	PhotoPath   string    `json:"photo_path"`
	ReligionID  int       `json:"religion_id"`
	CasteID     int       `json:"caste_id"`
	SubCasteID  int       `json:"sub_caste_id"`
	ClassID     int       `json:"class_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
