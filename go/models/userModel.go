package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string     `gorm:"type:varchar(30);not null"`
	LastName  string     `gorm:"type:varchar(30);not null"`
	Username  string     `gorm:"username:varchar(30);not null"`
	Password  string     `gorm:"password:varchar(30);not null"`
	Email     string     `gorm:"email:varchar(30);not null"`
	Address   string     `gorm:"adress:varchar(30);not null"`
	Feedbacks []Feedback `gorm:"feedbacks"`
}
