package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	UserID uint   `gorm:"not null" json:"user_id"`
}
