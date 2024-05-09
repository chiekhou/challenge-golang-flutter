package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string     `gorm:"size:64; not null; unique" json:"first_name"`
	LastName  string     `gorm:"size:64; not null" json:"last_name"`
	Username  string     `gorm:"size:64; not null; unique " json:"username"`
	Password  string     `gorm:"size:255; not null" json:"password"`
	Email     string     `gorm:"size:100; not null; unique" json:"email"`
	Address   string     `gorm:"size:255; not null" json:"address"`
	RoleId    uint       `gorm:"not null;DEFAULT: 0" json:"role_id"`
	Feedbacks []Feedback `gorm:"feedbacks"`
}
