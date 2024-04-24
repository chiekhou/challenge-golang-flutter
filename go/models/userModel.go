package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string     `gorm:"size:64"`
	LastName  string     `gorm:"size:64"`
	Username  string     `gorm:"size:64"`
	Password  string     `gorm:"size:255"`
	Email     string     `gorm:"size:100"`
	Address   string     `gorm:"size:255"`
	Feedbacks []Feedback `gorm:"feedbacks"`
}
