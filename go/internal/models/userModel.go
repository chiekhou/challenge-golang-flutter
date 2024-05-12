package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string       `gorm:"size:64; not null; unique" json:"first_name"`
	LastName    string       `gorm:"size:64; not null" json:"last_name"`
	Username    string       `gorm:"size:64; not null; unique " json:"username"`
	Password    string       `gorm:"size:255; not null" json:"password"`
	Email       string       `gorm:"size:100; not null; unique" json:"email"`
	Address     string       `gorm:"size:255; not null" json:"address"`
	RoleID      uint         `json:"role_id"`
	Role        Role         `gorm:"foreignKey:RoleID; not null; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	GroupVoyage GroupeVoyage `json:"group_voyage"`
	Destination Destination  `json:"destination"`
}
