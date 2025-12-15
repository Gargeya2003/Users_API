package models

import "time"

/* ---------- Domain / Response Model ---------- */

type User struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	DOB  time.Time `json:"dob"`
	Age  int       `json:"age,omitempty"`
}

/* ---------- Request DTOs ---------- */

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required"` // yyyy-mm-dd
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required"`
}
