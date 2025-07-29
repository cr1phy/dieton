package models

import (
	"time"
)

type Gender uint

const (
	Male Gender = iota + 1
	Female
	Intersex
)

func (gender Gender) String() string {
	terms := []string{"Male", "Female", "Intersex"}
	if gender < Male || gender > Intersex {
		return "unknown"
	}
	return terms[gender]
}

type User struct {
	ID           uint
	FirstName    string  `json:"fname" validate:"required"`
	LastName     string  `json:"lname"`
	Email        string  `json:"email" validate:"required,email"`
	PasswordHash string  `json:"password" validate:"required"`
	Age          uint8   `json:"age" validate:"gte=0,lte=130"`
	Gender       Gender  `json:"gender" validate:"required"`
	Height       float64 `json:"height" validate:"required"`
	Weight       float64 `json:"weight" validate:"required"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
